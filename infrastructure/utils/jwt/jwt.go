package jwt

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/units"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	ID   uint
	Role int
}

func GenerateToken(id uint, role int) (string, error) {
	var (
		conf      = config.Get().Auth
		expiresAt = time.Now().Add(units.Duration(conf.JwtExpireTime))
	)

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    conf.Issuer,
			Subject:   fmt.Sprintf("auth-%d", id),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
		ID:   id,
		Role: role,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	claims := new(Claims)

	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Auth.JwtSecret), nil
	})
	if err != nil {
		return claims, err
	}

	if !t.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}
