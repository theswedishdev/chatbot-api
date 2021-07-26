package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/services"
)

func JokeCommandsRoutes(app fiber.Router) {
	r := app.Group("/jokes")

	r.Get("", services.GetNotFound)
	r.Get("/chucknorris", services.GetCommandChuckNorris)
}
