package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecretKey         = "your_secret"
	defaultTokenDuration = 24 * time.Hour
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateJwtToken(userId uint, email string, name string,role string) (string, error)
	TokenValidation(tokenStr string) (*JWTClaims, error) 
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJwtService(secretKey string, duration time.Duration) JwtService {
	return &jwtService{
		secretKey, duration,
	}
}

func (js *jwtService) GenerateJwtToken(userId uint, email string, name string, role string) (string, error) {
	claims := JWTClaims{
		UserID: userId,
		Email:  email,
		Name:   name,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(js.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "sportsync",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (js *jwtService) TokenValidation(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&JWTClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return []byte(js.secretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unexpected signing method: %w", err)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil

}
