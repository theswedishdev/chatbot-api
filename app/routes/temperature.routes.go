package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/services"
)

func TemperatureCommandsRoutes(app fiber.Router) {
	r := app.Group("/temperature")

	// Temperature conversion from Celsius
	r.Get("/ctof/:temperature", services.GetCommandCelsiusToFahrenheit)
	r.Get("/ctok/:temperature", services.GetCommandCelsiusToKelvin)

	// Temperature conversion from Fahrenheit
	r.Get("/ftoc/:temperature", services.GetCommandFahrenheitToCelsius)
	r.Get("/ftok/:temperature", services.GetCommandFahrenheitToKelvin)

	// Temperature conversion from Kelvin
	r.Get("/ktoc/:temperature", services.GetCommandKelvinToCelsius)
	r.Get("/ktof/:temperature", services.GetCommandKelvinToFahrenheit)
}
