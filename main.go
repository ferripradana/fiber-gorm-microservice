package main

import (
	"fiber-gorm-microservice/infrastructure/repository/config"
	"fiber-gorm-microservice/infrastructure/rest/controllers/errors"
	"fiber-gorm-microservice/infrastructure/rest/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var err error
	db, err := config.GormOpen()
	if err != nil {
		fmt.Errorf("fatal error in database file: #{err}")
	}

	app := fiber.New()
	app.Use(errors.Handler)
	routes.ApplicationV1Router(app, db)
	app.Listen(":3000")
}
