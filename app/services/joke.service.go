package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theswedishdev/chatbot-api/app/commands"
)

func GetCommandChuckNorris(c *fiber.Ctx) error {
	joke, err := commands.GetChuckNorrisJoke(); if err != nil {
		return err
	}

	return c.SendString(joke)
}
