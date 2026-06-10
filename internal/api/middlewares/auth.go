package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

// InitFirebase inicializa a conexão com o Firebase usando o arquivo JSON da conta de serviço
func InitFirebase(credentialPath string) error {
	opt := option.WithCredentialsFile(credentialPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	FirebaseApp = app
	return nil
}

// FirebaseAuthMiddleware verifica o token JWT nas requisições protegidas
func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido ou inválido"})
			c.Abort()
			return
		}

		// 🔴 MOCK APENAS PARA DESENVOLVIMENTO
		if os.Getenv("APP_ENV") == "development" && authHeader == "Bearer dev-admin-token" {
			c.Set("UID", "admin-local-123")
			c.Set("Claims", map[string]interface{}{"admin": true})
			c.Next()
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Verifica o token usando o Firebase Admin SDK
		authClient, err := FirebaseApp.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro de integração com autenticação"})
			c.Abort()
			return
		}

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		// Injeta o UID e Claims no contexto
		c.Set("UID", token.UID)
		c.Set("Claims", token.Claims)

		c.Next()
	}
}

// VerifyAdmin middleware verifica se o usuário autenticado possui privilégios de administrador
func VerifyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsData, exists := c.Get("Claims")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: claims não encontrados"})
			c.Abort()
			return
		}

		claims, ok := claimsData.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: claims em formato inválido"})
			c.Abort()
			return
		}

		isAdmin, ok := claims["admin"].(bool)
		if !ok || !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado: privilégios de administrador requeridos"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AppCheckMiddleware verifica o token do Firebase App Check para barrar requisições não autênticas
func AppCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bypass para ambiente de desenvolvimento
		if os.Getenv("APP_ENV") == "development" {
			c.Next()
			return
		}

		appCheckToken := c.GetHeader("X-Firebase-AppCheck")
		if appCheckToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "App Check token não fornecido."})
			c.Abort()
			return
		}

		appCheckClient, err := FirebaseApp.AppCheck(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inicializar verificação do App Check"})
			c.Abort()
			return
		}

		_, err = appCheckClient.VerifyToken(appCheckToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "App Check token inválido."})
			c.Abort()
			return
		}

		c.Next()
	}
}
