package main

import (
	"fiber-gorm-microservice/infrastructure/repository/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	var err error
	_, err = config.GormOpen()
	if err != nil {
		fmt.Errorf("fatal error in database file: #{err}")
	}

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})
	app.Listen(":3000")
}
