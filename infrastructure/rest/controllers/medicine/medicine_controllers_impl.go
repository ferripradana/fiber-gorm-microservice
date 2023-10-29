package medicine

import (
	medicineService "fiber-gorm-microservice/application/service/medicine"
	"fiber-gorm-microservice/application/utils"
	"fiber-gorm-microservice/domain/errors"
	domainMedicine "fiber-gorm-microservice/domain/medicine"
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
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(domainMedicine)
}

func (controller *MedicineControllerImpl) GetMedicineById(ctx *fiber.Ctx) error {
	medicineId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError, fiber.StatusBadRequest)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	domainMedicine, err := controller.Service.GetById(medicineId)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(domainMedicine)
}

func (controller *MedicineControllerImpl) DeleteMedicine(ctx *fiber.Ctx) error {
	medicineId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError, fiber.StatusBadRequest)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	err = controller.Service.Delete(medicineId)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Resource deleted successfully"})
}

func (controller *MedicineControllerImpl) UpdateMedicine(ctx *fiber.Ctx) error {
	medicineId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := errors.NewAppErrorImpl(err, errors.ValidationError, fiber.StatusBadRequest)
		return fiber.NewError(appError.(*errors.AppErrorImpl).Status, appError.Error())
	}

	requestBody := make(map[string]interface{})
	if err := ctx.BodyParser(&requestBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = UpdateValidation(requestBody)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	var medicine *domainMedicine.Medicine
	medicine, err = controller.Service.Update(medicineId, requestBody)
	if err != nil {
		return fiber.NewError(err.(*errors.AppErrorImpl).Status, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(medicine)
}
