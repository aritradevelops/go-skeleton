package http

import "skeleton-test/internal/handlers"

func (s *Server) setupRoutes() {
	handlers := handlers.New(s.db, s.config, s.services)
	root := s.app
	// base routes
	root.Get("/", handlers.Hello)
	root.Get("/health", handlers.HealthCheck)

	// auth routes
	authRoutes := root.Group("/auth")
	authRoutes.Post("/register", handlers.Auth.Register)
}
