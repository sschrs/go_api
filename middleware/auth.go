package middleware

import "github.com/gofiber/fiber/v2"

// dummy function for authentication middleware layer
func Auth(c *fiber.Ctx) error {
	return c.Next()
}
