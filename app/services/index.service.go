package services

import (
	"github.com/gofiber/fiber/v2"
)

func GetIndex(c *fiber.Ctx) error {
	return c.SendFile("public/index.html")
}

func GetStyleCss(c *fiber.Ctx) error {
	return c.SendFile("public/style.css")
}

func GetNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.ErrNotFound.Code).SendString(fiber.ErrNotFound.Message)
}
