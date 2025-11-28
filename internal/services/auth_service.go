package services

import (
	"context"
	"skeleton-test/internal/sqlc"
)

type AuthService struct {
	queries *sqlc.Queries
}

// TODO: add validation logic
type RegisterUserParams struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func (s *AuthService) RegisterUser(ctx context.Context, payload RegisterUserParams) error {
	// TODO: validate
	_, err := s.queries.RegisterUser(ctx, sqlc.RegisterUserParams{
		Email: payload.Email,
		Name:  payload.Name,
	})
	if err != nil {
		return err
	}
	// TODO: send email
	return nil
}
