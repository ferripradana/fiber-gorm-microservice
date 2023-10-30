package user

import "github.com/gofiber/fiber/v2"

type UserController interface {
	NewUser(ctx *fiber.Ctx) error
}
