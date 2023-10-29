package medicine

import (
	medicineService "fiber-gorm-microservice/application/service/medicine"
	"fiber-gorm-microservice/application/utils"
	"fiber-gorm-microservice/domain/errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type MedicineControllerImpl struct {
	Service medicineService.MedicineService
}

func NewMedicineControllerImpl(service medicineService.MedicineService) MedicineController {
	return &MedicineControllerImpl{
		Service: service,
	}
}

func (controller *MedicineControllerImpl) GetAllMedicines(ctx *fiber.Ctx) error {
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
	medicines, err := controller.Service.GetAll(page, limit)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(medicines)
}

func (controller *MedicineControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(NewMedicineRequest)
	if err := utils.NewValidation().ValidateRequest(ctx, request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newMedicine := medicineService.NewMedicine{
		Name:        request.Name,
		Description: request.Description,
		EANCode:     request.EanCode,
		Laboratory:  request.Laboratory,
	}

	domainMedicine, err := controller.Service.Create(&newMedicine)
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.RepositoryError, fiber.StatusInternalServerError)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(domainMedicine)
}
