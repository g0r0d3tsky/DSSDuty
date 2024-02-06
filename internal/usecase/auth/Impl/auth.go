package Impl

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const (
	salt       = "some-salt"
	signingKey = "some-signing-key"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId uuid.UUID `json:"user_id"`
}

func GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		userID,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.UUID{}, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
