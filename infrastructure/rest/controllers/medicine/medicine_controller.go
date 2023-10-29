package medicine

import "github.com/gofiber/fiber/v2"

type MedicineController interface {
	GetAllMedicines(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	GetMedicineById(ctx *fiber.Ctx) error
	DeleteMedicine(ctx *fiber.Ctx) error
	UpdateMedicine(ctx *fiber.Ctx) error
}
