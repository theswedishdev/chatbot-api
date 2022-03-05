package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/commands"
)

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
