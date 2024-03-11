package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GenerateToken(id string, username string, password string) (string, error)
	VerifyToken(tokenString string) (interface{}, error)
}

type authService struct {
	secretKey []byte
}

func NewAuthService(secretKey []byte) *authService {
	return &authService{
		secretKey: secretKey,
	}
}

func (a *authService) GenerateToken(id string, username string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    username,
		Subject:   "jwt-auth",
		ID:        id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *authService) VerifyToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims.Issuer, nil
}
