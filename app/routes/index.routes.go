package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/services"
)

func IndexRoutes(app fiber.Router) {
	r := app.Group("/")

	r.Get("", services.GetIndex)
	r.Get("style.css", services.GetStyleCss)
}
