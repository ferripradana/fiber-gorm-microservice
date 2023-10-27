package medicine

import (
	medicineService "fiber-gorm-microservice/application/service/medicine"
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

func (controller MedicineControllerImpl) GetAllMedicines(ctx *fiber.Ctx) error {
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
