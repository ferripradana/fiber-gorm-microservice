package user

import (
	userService "fiber-gorm-microservice/application/service/user"
	"fiber-gorm-microservice/application/utils"
	"fiber-gorm-microservice/domain/errors"
	domainUser "fiber-gorm-microservice/domain/user"
	"github.com/gofiber/fiber/v2"
	"strconv"
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

func (controller *UserControllerImpl) GetUserById(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError, fiber.StatusBadRequest)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	user, err := controller.Service.GetById(userId)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(domainToResponseMapper(user))
}

func (controller *UserControllerImpl) GetAllUsers(ctx *fiber.Ctx) error {
	pageStr := ctx.Query("page", "1")
	limitStr := ctx.Query("limit", "10")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	users, err := controller.Service.GetaAll(page, limit)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	usersResponse := &PaginationResultUser{
		Data:       mapFromDomainToResponse(users.Data),
		Total:      users.Total,
		Limit:      users.Limit,
		Current:    users.Current,
		NextCursor: users.NextCursor,
		PrevCursor: users.PrevCursor,
		NumPages:   users.NumPages,
	}

	return ctx.Status(fiber.StatusOK).JSON(usersResponse)
}

func (controller *UserControllerImpl) DeleteUser(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError, fiber.StatusBadRequest)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	err = controller.Service.Delete(userId)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Resource deleted successfully"})
}

func (controller *UserControllerImpl) UpdateUser(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError, fiber.StatusBadRequest)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	requestBody := make(map[string]interface{})
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = UpdateValidation(requestBody)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	var user *domainUser.User
	user, err = controller.Service.Update(userId, requestBody)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(domainToResponseMapper(user))
}
