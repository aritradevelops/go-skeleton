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

// type RegisterUserParams struct {
//     Email    string `json:"email,omitempty" validate:"required,email"`
//     Name     string `json:"name,omitempty" validate:"required,min=3,alphaspace"`
//     Password string `json:"password" validate:"required,min=8,alphanumeric"`
// }

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var payload services.RegisterUserParams
	err := c.BodyParser(&payload)
	if err != nil {
		return err
	}
	fmt.Println("some message", payload)
	fmt.Println(h)
	err = h.services.Auth.RegisterUser(c.Context(), payload)
	if err != nil {
		return err
	}
	c.Status(http.StatusCreated)
	return c.JSON(NewSuccessResponse(translation.Localize(c, "user.register"), nil))
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var payload services.LoginUserParams
	err := c.BodyParser(&payload)
	if err != nil {
		return err
	}

	token, err := h.services.Auth.LoginUser(c.Context(), payload)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{Name: "access_token", Value: token})
	return c.JSON(NewSuccessResponse(translation.Localize(c, "user.login"), fiber.Map{
		"access_token": token,
	}))
}
