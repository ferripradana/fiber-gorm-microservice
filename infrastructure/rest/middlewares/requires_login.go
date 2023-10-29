package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func AuthJWTMiddleware(ctx *fiber.Ctx) error {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
	}

	JWTAccessSecure := viper.GetString("Secure.JWTAccessSecure")
	tokenString := ctx.Get("Authorization", "")
	signature := []byte(JWTAccessSecure)

	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "token is not provided"})
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid token"})
	}

	return ctx.Next()
}
