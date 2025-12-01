package handlers

import (
	"fmt"
	"net/http"
	"skeleton-test/internal/services"
	"skeleton-test/internal/translation"

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
	fmt.Println("some message", payload)
	err = h.services.Auth.RegisterUser(c.Context(), payload)
	if err != nil {
		return err
	}
	c.Status(http.StatusCreated)
	return c.JSON(NewSuccessResponse(translation.Localize(c, "user.register"), nil))
}
