package auth

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	GetAccessTokenByRefreshToken(ctx *fiber.Ctx) error
}
