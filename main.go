package main

import (
	"github.com/Rayato159/iam-fastest-api-alive/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	app.Listen(":3000")
}
