package routes

import (
	"github.com/Rayato159/iam-fastest-api-alive/src/controllers"
	"github.com/Rayato159/iam-fastest-api-alive/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/hello", controllers.Hello)

	app.Use(middlewares.NotFoundErr())
}
