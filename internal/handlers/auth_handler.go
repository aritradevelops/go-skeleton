package handlers

import (
	"net/http"
	"skeleton-test/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	services *services.Services
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var payload services.RegisterUserParams
	err := c.BodyParser(&payload)
	if err != nil {
		return err
	}
	err = h.services.Auth.RegisterUser(c.Context(), payload)
	if err != nil {
		return err
	}
	c.Status(http.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "User registered successfully!.",
		"data":    nil,
		"error":   nil,
	})
}
