package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, controller auth.AuthController) {
	routerAuth := router.Group("/auth")
	routerAuth.Post("/login", controller.Login)
}
