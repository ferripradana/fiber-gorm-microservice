package user

import "github.com/gofiber/fiber/v2"

type UserController interface {
	NewUser(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}
