package main

import (
	"fiber-gorm-microservice/infrastructure/repository/config"
	"fiber-gorm-microservice/infrastructure/rest/controllers/errors"
	"fiber-gorm-microservice/infrastructure/rest/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var err error
	db, err := config.GormOpen()
	if err != nil {
		panic("Database not connect")
	}

	app := fiber.New()
	app.Use(errors.Handler)
	routes.ApplicationV1Router(app, db)
	app.Listen(":3000")
}
