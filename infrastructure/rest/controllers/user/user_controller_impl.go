package user

import (
	userService "fiber-gorm-microservice/application/service/user"
	"fiber-gorm-microservice/application/utils"
	"fiber-gorm-microservice/domain/errors"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	Service userService.UserService
}

func NewUserControllerImpl(service userService.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (controller *UserControllerImpl) NewUser(ctx *fiber.Ctx) error {
	request := new(NewUserRequest)
	if err := utils.NewValidation().ValidateRequest(ctx, request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := controller.Service.Create(toUserServiceMapper(request))
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}
	userResponse := domainToResponseMapper(user)
	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}
