package utils

import (
	"errors"
	"finenance-app/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(userId string, role string, expiredTime time.Duration) (string, error) {

	// Create the Claims
	claims := entity.CustomClaims{
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "123",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(viper.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ValidateToken(tokenString string) (*entity.CustomClaims, error) {
	// Parse token dengan claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan pakai algoritma yang benar
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(viper.GetString("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	// Ambil claims setelah parse
	claims, ok := token.Claims.(*entity.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
