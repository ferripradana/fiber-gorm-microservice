package adapter

import (
	service "fiber-gorm-microservice/application/service/auth"
	repository "fiber-gorm-microservice/infrastructure/repository/user"
	"fiber-gorm-microservice/infrastructure/rest/controllers/auth"
	"gorm.io/gorm"
)

func AuthAdapter(db *gorm.DB) auth.AuthController {
	userRepository := repository.NewUserRepositoryImpl(db)
	authService := service.NewAuthServiceImpl(userRepository)
	return auth.NewAuthControllerImpl(authService)
}
