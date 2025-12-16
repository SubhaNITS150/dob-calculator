package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		id := uuid.New().String()
		c.Set("X-Request-ID", id)

		err := c.Next()

		duration := time.Since(start)
		c.App().Config().ErrorHandler(c, err)

		_ = duration 
		return err
	}
}
