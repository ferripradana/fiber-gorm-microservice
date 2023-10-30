package main

import (
	"fiber-gorm-microservice/infrastructure/repository/config"
	"fiber-gorm-microservice/infrastructure/rest/controllers/errors"
	"fiber-gorm-microservice/infrastructure/rest/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		panic(err)
	}
	serverPort := fmt.Sprintf("%s", viper.GetString("ServerPort"))
	serverAddress := fmt.Sprintf("%s", viper.GetString("Address"))

	db, err := config.GormOpen()
	if err != nil {
		panic("Database not connect")
	}

	app := fiber.New()
	app.Use(errors.Handler)
	routes.ApplicationV1Router(app, db)
	addr := fmt.Sprintf("%s:%s", serverAddress, serverPort)
	app.Listen(addr)
}
