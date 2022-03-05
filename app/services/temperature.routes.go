package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/commands"
)

func GetCommandCelsiusToFahrenheit(c *fiber.Ctx) error {
	temperatureFloat, err := commands.GetTemperatureFloat(c)
	if err != nil {
		return err
	}

	return c.SendString(commands.CToF(temperatureFloat))
}

func GetCommandCelsiusToKelvin(c *fiber.Ctx) error {
	temperatureFloat, err := commands.GetTemperatureFloat(c)
	if err != nil {
		return err
	}

	return c.SendString(commands.CToK(temperatureFloat))
}

func GetCommandFahrenheitToCelsius(c *fiber.Ctx) error {
	temperatureFloat, err := commands.GetTemperatureFloat(c)
	if err != nil {
		return err
	}

	return c.SendString(commands.FToC(temperatureFloat))
}

func GetCommandFahrenheitToKelvin(c *fiber.Ctx) error {
	temperatureFloat, err := commands.GetTemperatureFloat(c)
	if err != nil {
		return err
	}

	return c.SendString(commands.FToK(temperatureFloat))
}

func GetCommandKelvinToCelsius(c *fiber.Ctx) error {
	temperatureFloat, err := commands.GetTemperatureFloat(c)
	if err != nil {
		return err
	}

	return c.SendString(commands.KToC(temperatureFloat))
}

func GetCommandKelvinToFahrenheit(c *fiber.Ctx) error {
	temperatureFloat, err := commands.GetTemperatureFloat(c); if err != nil {
		return err
	}

	return c.SendString(commands.KToF(temperatureFloat))
}
