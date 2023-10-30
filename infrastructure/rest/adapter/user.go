package adapter

import (
	service "fiber-gorm-microservice/application/service/user"
	repository "fiber-gorm-microservice/infrastructure/repository/user"
	user "fiber-gorm-microservice/infrastructure/rest/controllers/user"
	"gorm.io/gorm"
)

func UserAdapter(db *gorm.DB) user.UserController {
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	return user.NewUserControllerImpl(userService)

}
