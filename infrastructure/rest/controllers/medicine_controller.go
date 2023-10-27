package controllers

import "github.com/gofiber/fiber/v2"

type MedicineController interface {
	GetAllMedicines(ctx *fiber.Ctx) error
}
