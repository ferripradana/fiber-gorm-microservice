package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/controllers/medicine"
	"github.com/gofiber/fiber/v2"
)

func MedicineRoutes(router fiber.Router, controller medicine.MedicineController) {
	routerMedicine := router.Group("/medicine")
	routerMedicine.Get("/", controller.GetAllMedicines)
}
