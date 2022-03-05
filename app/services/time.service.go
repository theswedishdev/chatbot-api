package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/commands"
)

func GetCommandCurrentTime12hr(c *fiber.Ctx) error {
	currentTime, err := commands.GetCurrentTimeByTimezone12hr(c.Params("timezone")); if err != nil {
		return c.SendString("Invalid timezone")
	}

	return c.SendString(currentTime)
}

func GetCommandCurrentTime24hr(c *fiber.Ctx) error {
	currentTime, err := commands.GetCurrentTimeByTimezone24hr(c.Params("timezone")); if err != nil {
		return c.SendString("Invalid timezone")
	}

	return c.SendString(currentTime)
}
