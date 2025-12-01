package services

import (
	"context"
	"fmt"
	"skeleton-test/internal/sqlc"
	"skeleton-test/internal/validation"

	"golang.org/x/crypto/bcrypt"
)

const passwordHashingCost = 10

type AuthService struct {
	queries *sqlc.Queries
}

// TODO: add validation logic
type RegisterUserParams struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Name     string `json:"name,omitempty" validate:"required,min=3,alphaspace"`
	Password string `json:"password" validate:"required,min=8,alphanumeric"`
}

func (s *AuthService) RegisterUser(ctx context.Context, payload RegisterUserParams) error {
	fmt.Println("here")
	validationErrors := validation.Validate(payload)
	if validationErrors != nil {
		fmt.Println("here 2")
		return validationErrors
	}

	// insert user
	user, err := s.queries.RegisterUser(ctx, sqlc.RegisterUserParams{
		Email: payload.Email,
		Name:  payload.Name,
	})
	if err != nil {
		return err
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), passwordHashingCost)
	if err != nil {
		return err
	}

	// insert password
	p, err := s.queries.InsertPassword(ctx, sqlc.InsertPasswordParams{
		HashedPassword: string(hashedPassword),
		CreatedBy:      user.ID,
	})
	if err != nil {
		fmt.Printf("failed to insert password: %v", err)
		return err
	}
	fmt.Println(p)

	// TODO: send email
	return nil
}
