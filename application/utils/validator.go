package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validation struct {
	Validator *validator.Validate
}

func NewValidation() *Validation {
	return &Validation{
		Validator: validator.New(),
	}
}

func (v *Validation) ValidateRequest(ctx *fiber.Ctx, request interface{}) error {
	if err := ctx.BodyParser(request); err != nil {
		return err
	}
	if err := v.Validator.Struct(request); err != nil {
		return err
	}
	return nil
}
