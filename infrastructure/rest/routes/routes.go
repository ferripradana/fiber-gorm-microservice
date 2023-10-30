package routes

import (
	"fiber-gorm-microservice/infrastructure/rest/adapter"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ApplicationV1Router(router *fiber.App, db *gorm.DB) {
	routerV1 := router.Group("/v1")
	MedicineRoutes(routerV1, adapter.MedicineAdapter(db))
	AuthRoutes(routerV1, adapter.AuthAdapter(db))
	UserRoutes(routerV1, adapter.UserAdapter(db))
}
