package handlers

import (
	"skeleton-test/internal/config"
	"skeleton-test/internal/db"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	db     db.Database
	config config.Config
}

func New(db db.Database, config config.Config) *Handlers {
	return &Handlers{
		db:     db,
		config: config,
	}
}

func (h *Handlers) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello!",
		"name":    h.config.ServiceName,
		"status":  "running",
	})
}

func (h *Handlers) HealthCheck(c *fiber.Ctx) error {
	err := h.db.Health()
	status := "healthy"
	if err != nil {
		status = "unhealthy"
	}
	return c.JSON(fiber.Map{
		"status":    status,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
