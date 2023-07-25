package middlewares

import "github.com/gofiber/fiber/v2"

func NotFoundErr() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}
}
