package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/services"
)

func AdviceCommandsRoutes(app fiber.Router) {
	r := app.Group("/advice")

	r.Get("", services.GetCommandRandomAdvice)
	r.Get("/:ID", services.GetCommandAdviceByID)
}
