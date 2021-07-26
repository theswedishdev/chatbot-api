package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/theswedishdev/chatbot-api/app/routes"
	"github.com/theswedishdev/chatbot-api/app/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
		ProxyHeader:  "X-Forwarded-For",
	})

	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon.png",
	}))
	app.Use(recover.New())
	app.Use(logger.New())

	routes.IndexRoutes(app)
	routes.CommandsRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
