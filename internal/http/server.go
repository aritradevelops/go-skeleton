package http

import (
	"context"
	"fmt"
	"skeleton-test/internal/config"
	"skeleton-test/internal/db"
	"skeleton-test/internal/translation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app    *fiber.App
	db     db.Database
	config config.Config
}

func NewServer(config config.Config, db db.Database) *Server {

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(translation.New())
	server := &Server{
		config: config,
		db:     db,
		app:    app,
	}

	server.setupRoutes()

	return server
}

func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf("%s:%d", s.config.Http.Host, s.config.Http.Port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
