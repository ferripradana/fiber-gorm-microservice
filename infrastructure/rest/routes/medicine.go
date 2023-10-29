package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/controllers/medicine"
	"fiber-gorm-microservice/infrastructure/rest/routes/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MedicineRoutes(router fiber.Router, controller medicine.MedicineController) {
	routerMedicine := router.Group("/medicine")
	routerMedicine.Use(middlewares.AuthJWTMiddleware)
	{
		routerMedicine.Get("/", controller.GetAllMedicines)
	}

}
