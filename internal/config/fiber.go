package config

import (
	"github.com/gofiber/fiber/v3"
)

func InitFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName:      "finenance-app",
			ErrorHandler: NewErrorHandler(),
		})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		return ctx.Status(code).JSON(fiber.Map{
			"code":    code,
			"message": message,
		})
	}
}
