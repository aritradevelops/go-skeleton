package services

import (
	"context"
	"skeleton-test/internal/sqlc"
	"skeleton-test/internal/validation"
)

type AuthService struct {
	queries *sqlc.Queries
}

// TODO: add validation logic
type RegisterUserParams struct {
	Email string `json:"email,omitempty" validate:"required,email"`
	Name  string `json:"name,omitempty" validate:"required,min=3,alphaspace"`
}

func (s *AuthService) RegisterUser(ctx context.Context, payload RegisterUserParams) error {

	validationErrors := validation.Validate(payload)
	if validationErrors != nil {
		return validationErrors
	}

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
