package routes

import (
	"github.com/MatTwix/Go-online-resume-module/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	resume := api.Group("/resume")
	resume.Get("/", handlers.GetResume)
	resume.Put("/update", handlers.UpdateResume)
}
