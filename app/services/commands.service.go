package services

import (
	"github.com/gofiber/fiber/v2"
)

func GetCommandsIndex(c *fiber.Ctx) error {
	return c.SendFile("public/commands.html")
}
