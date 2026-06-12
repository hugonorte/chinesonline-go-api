package rules

import "github.com/hugonorte/chinesonline-go-api/internal/models"

// GameRule define as regras matemáticas de um GameType específico
type GameRule struct {
	PointsPerCorrectAnswer int // Quantidade de pontos ganhos por acerto
	PointsToLevelUp        int // Quantidade de pontos acumulados necessários para subir de nível
}

// DefaultRules centraliza e isola as configurações de pontuação.
// Para alterar as regras do jogo, basta editar este mapa.
var DefaultRules = map[models.GameType]GameRule{
	models.GameTypeTranslation: {
		PointsPerCorrectAnswer: 10,
		PointsToLevelUp:        100,
	},
	models.GameTypePinyinWithoutTone: {
		PointsPerCorrectAnswer: 5,
		PointsToLevelUp:        50,
	},
	models.GameTypePinyinWithNumericTone: {
		PointsPerCorrectAnswer: 8,
		PointsToLevelUp:        80,
	},
	models.GameTypePinyinWithSimbolTone: {
		PointsPerCorrectAnswer: 8,
		PointsToLevelUp:        80,
	},
	models.GameTypeTranslationTimed: {
		PointsPerCorrectAnswer: 15,
		PointsToLevelUp:        150,
	},
	models.GameTypePinyinWithoutToneTimed: {
		PointsPerCorrectAnswer: 10,
		PointsToLevelUp:        100,
	},
	models.GameTypePinyinWithNumericToneTimed: {
		PointsPerCorrectAnswer: 12,
		PointsToLevelUp:        120,
	},
	models.GameTypePinyinWithSimbolToneTimed: {
		PointsPerCorrectAnswer: 12,
		PointsToLevelUp:        120,
	},
}

// GetRule retorna a regra para um dado tipo de jogo, com um valor padrão de fallback
func GetRule(gameType models.GameType) GameRule {
	if rule, exists := DefaultRules[gameType]; exists {
		return rule
	}
	// Fallback seguro caso um novo GameType seja criado e esquecido aqui
	return GameRule{
		PointsPerCorrectAnswer: 5,
		PointsToLevelUp:        100,
	}
}
