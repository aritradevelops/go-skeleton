package http

import "skeleton-test/internal/handlers"

func (s *Server) setupRoutes() {
	handlers := handlers.New(s.db, s.config)

	s.app.Get("/", handlers.Hello)
	s.app.Get("/health", handlers.HealthCheck)

}
