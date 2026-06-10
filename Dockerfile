# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Baixar dependências (para melhor cache no Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Fazer o build estático, minificando o binário (flags -w -s)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main cmd/api/main.go

# Final stage (Imagem limpa e otimizada)
FROM alpine:latest

# Instalar certificados raiz (essencial para conectar no Neon DB e Firebase)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copiar apenas o binário compilado do stage builder
COPY --from=builder /app/main .

# Definir a variável PORT para o padrão do Google Cloud Run
ENV PORT=8080
EXPOSE 8080

# Rodar a aplicação
CMD ["./main"]
