package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/controllers/medicine"
	"fiber-gorm-microservice/infrastructure/rest/middlewares"
	"github.com/gofiber/fiber/v2"
)

func MedicineRoutes(router fiber.Router, controller medicine.MedicineController) {
	routerMedicine := router.Group("/medicine")
	routerMedicine.Use(middlewares.AuthJWTMiddleware)
	{
		routerMedicine.Get("/", controller.GetAllMedicines)
		routerMedicine.Post("/", controller.Create)
		routerMedicine.Get("/:id", controller.GetMedicineById)
		routerMedicine.Delete("/:id", controller.DeleteMedicine)
		routerMedicine.Put("/:id", controller.UpdateMedicine)
	}

}
