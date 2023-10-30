package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/controllers/user"
	"fiber-gorm-microservice/infrastructure/rest/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, controller user.UserController) {
	routerUser := router.Group("/user")
	routerUser.Use(middlewares.AuthJWTMiddleware)
	{
		routerUser.Post("/", controller.NewUser)
		routerUser.Get("/:id", controller.GetUserById)
		routerUser.Get("/", controller.GetAllUsers)
		routerUser.Delete("/:id", controller.DeleteUser)
		routerUser.Put("/:id", controller.UpdateUser)
	}
}
