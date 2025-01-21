package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"vega-server/pkg/config"

	"errors"
	"time"
)

type CustomClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	key             []byte
	expirationHours int
}

func NewJWTService(config *config.Config) *JWTService {
	return &JWTService{
		key:             []byte(config.GetString("jwt.key")),
		expirationHours: config.GetInt("jwt.expiration_hours"),
	}
}

func (jwtService *JWTService) GenerateJWT(id uint) (string, error) {
	expirationAt := time.Now().Add(time.Duration(jwtService.expirationHours) * time.Hour)
	claims := &CustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtService.key)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}

func (jwtService *JWTService) ValidateJWT(tokenString string) (*CustomClaims, error) {
	// Parse Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("failed to validate token")
		}
		return jwtService.key, nil
	})
	if err != nil {
		return nil, errors.New("failed to validate token")
	}
	// Extract Claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("failed to validate token")
}
