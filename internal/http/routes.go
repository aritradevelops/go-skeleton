package http

import (
	"fmt"
	"skeleton-test/internal/handlers"
)

func (s *Server) setupRoutes() {
	fmt.Println("services", s.services)
	handlers := handlers.New(s.db, s.config, s.services)
	root := s.app
	// base routes
	root.Get("/", handlers.Hello)
	root.Get("/health", handlers.HealthCheck)

	// auth routes
	authRoutes := root.Group("/auth")
	authRoutes.Post("/register", handlers.Auth.Register)
	authRoutes.Post("/login", handlers.Auth.Login)
}
