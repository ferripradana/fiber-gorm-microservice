package auth

import (
	authService "fiber-gorm-microservice/application/service/auth"
	"fiber-gorm-microservice/application/utils"
	"fiber-gorm-microservice/domain/errors"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService authService.AuthService
}

func NewAuthControllerImpl(service authService.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: service,
	}
}

func (controller *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	request := new(LoginRequest)
	if err := utils.NewValidation().ValidateRequest(ctx, request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user := authService.LoginUser{
		Email:    request.Email,
		Password: request.Password,
	}
	authDataUser, err := controller.AuthService.Login(user)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.NotAuthorized, fiber.StatusUnauthorized)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(authDataUser)
}

func (controller *AuthControllerImpl) GetAccessTokenByRefreshToken(ctx *fiber.Ctx) error {
	request := new(AccessTokenRequest)
	if err := utils.NewValidation().ValidateRequest(ctx, request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	authDataUser, err := controller.AuthService.AccessTokenByRefreshToken(request.RefreshToken)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.NotAuthorized, fiber.StatusUnauthorized)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(authDataUser)
}
