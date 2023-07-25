package controllers

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "hello, Fiber!",
	})
}
