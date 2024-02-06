package auth

import "github.com/google/uuid"

type AuthUseCase interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
}
