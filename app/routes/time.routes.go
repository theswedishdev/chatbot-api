package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/services"
)

func TimeCommandsRoutes(app fiber.Router) {
	r := app.Group("/time")

	r.Get("/current/12/:timezone", services.GetCommandCurrentTime12hr)
	r.Get("/current/24/:timezone", services.GetCommandCurrentTime24hr)
}
