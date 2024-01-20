package middleware

import "github.com/gofiber/fiber/v2"

func SetContentType(contentType string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Type(contentType)
		return c.Next()
	}
}
