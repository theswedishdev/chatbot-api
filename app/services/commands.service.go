package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/commands"
)

func GetCommandsIndex(c *fiber.Ctx) error {
	return c.SendFile("public/commands.html")
}

func GetCommandRandomAdvice(c *fiber.Ctx) error {
	adviceString, err := commands.GetAdvice(""); if err != nil {
		return err
	}

	return c.SendString(adviceString)
}

func GetCommandAdviceByID(c *fiber.Ctx) error {
	ID := c.Params("ID")

	adviceString, err := commands.GetAdvice(ID); if err != nil {
		return err
	}

	return c.SendString(adviceString)
}

func GetCommandChuckNorris(c *fiber.Ctx) error {
	joke, err := commands.GetChuckNorrisJoke(); if err != nil {
		return err
	}

	return c.SendString(joke)
}

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
