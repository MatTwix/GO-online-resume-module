package main

import (
	"log"

	"github.com/MatTwix/Go-online-resume-module/config"
	"github.com/MatTwix/Go-online-resume-module/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{cfg.AppUrl + ":" + cfg.ReactPort},
		AllowMethods: []string{"GET", "PUT"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	routes.SetupRoutes(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
