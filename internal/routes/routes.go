package routes

import (
	"github.com/SubhaNITS150/dob-calculator/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.Create)
	app.Put("/users/:id", h.Update)
	app.Delete("/users/:id", h.Delete)
	app.Get("/users", h.List)
	app.Get("/health", h.Health)
}
