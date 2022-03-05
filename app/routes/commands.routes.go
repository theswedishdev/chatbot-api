package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/services"
)

func CommandsRoutes(app fiber.Router) {
	r := app.Group("/commands")

	r.Get("", services.GetCommandsIndex)

	AdviceCommandsRoutes(r)
	JokeCommandsRoutes(r)
	TemperatureCommandsRoutes(r)
	TimeCommandsRoutes(r)
}
