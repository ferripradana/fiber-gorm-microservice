package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/controllers"
	"github.com/gofiber/fiber/v2"
)

func MedicineRoutes(router fiber.Router, controller controllers.MedicineController) {
	routerMedicine := router.Group("/medicine")
	routerMedicine.Get("/", controller.GetAllMedicines)
}
