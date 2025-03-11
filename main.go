package main

import (
	"fmt"
	"log"

	"github.com/MatTwix/Go-online-resume-module/config"
	"github.com/MatTwix/Go-online-resume-module/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()

	if cfg.ENV != "production" {
		originURL := fmt.Sprintf("%s:%s", cfg.AppUrl, cfg.ReactPort)

		app.Use(cors.New(cors.Config{
			AllowOrigins:     []string{originURL},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length"},
			MaxAge:           86400,
		}))
	}

	routes.SetupRoutes(app)

	if cfg.ENV == "production" {
		app.Use("/assets", static.New("./client/dist/assets", static.Config{
			Browse: false,
			MaxAge: 3600,
		}))

		app.Get("/*", func(c fiber.Ctx) error {
			return c.SendFile("./client/dist/index.html")
		})
	}

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
