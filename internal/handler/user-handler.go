package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/SubhaNITS150/dob-calculator/internal/repository"
	"github.com/SubhaNITS150/dob-calculator/internal/service"
)

type UserHandler struct {
	repo *repository.UserRepository
	v    *validator.Validate
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
		v:    validator.New(),
	}
}

// ================= CREATE USER =================
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req CreateUserDTO
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.v.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// DOB String -> time.Time
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"dob must be in YYYY-MM-DD format",
		)
	}

	user, err := h.repo.Create(c.Context(), req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(
		service.MapUserFromSQLC(user.ID, user.Name, user.Dob),
	)
}

// ================= UPDATE USER =================
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid user id",
		)
	}

	var req CreateUserDTO
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.v.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"dob must be in YYYY-MM-DD format",
		)
	}

	user, err := h.repo.Update(c.Context(), id, req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(
		service.MapUserFromSQLC(user.ID, user.Name, user.Dob),
	)
}

// ================= DELETE USER =================
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid user id",
		)
	}

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ================= LIST USERS =================
func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.repo.List(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	res := make([]interface{}, 0, len(users))
	for _, u := range users {
		res = append(res,
			service.MapUserFromSQLC(u.ID, u.Name, u.Dob),
		)
	}

	return c.JSON(res)
}

// ================= HEALTH= =================
func (h* UserHandler) Health(c *fiber.Ctx) error {
	return c.SendString("OK")
}