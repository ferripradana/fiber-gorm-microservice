package controllers

import (
	medicineService "fiber-gorm-microservice/application/service/medicine"
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
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	medicines, err := controller.Service.GetAll(page, limit)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(medicines)
}
