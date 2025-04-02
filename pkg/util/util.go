package util

import (
	"log"
	"time"

	"github.com/abilsabili50/middleware-with-go-fiber/pkg/config"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func GenerateToken(config *config.App, userId string) (string, errs.MessageErr) {
	// create claims
	claims := jwt.MapClaims{
		"id":    userId,
		"admin": true,
		"exp":   time.Now().Add(config.JWTAccTokenDuration).Unix(),
	}

	// create token with choosen signing method
	token := jwt.NewWithClaims(jwt.GetSigningMethod(config.JWTSigningMethod), claims)

	// generate token with secret key
	tokenString, err := token.SignedString([]byte(config.JWTAccTokenSecretKey))
	if err != nil {
		log.Printf("[ERROR] Server error while generate token - %v", err)
		return "", errs.NewInternalServerError("failed to generate token")
	}

	return tokenString, nil
}
