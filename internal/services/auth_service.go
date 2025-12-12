package services

import (
	"context"
	"fmt"
	"skeleton-test/internal/jwtx"
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
	Password string `json:"password" validate:"required,min=8,alphanum"`
}

type LoginUserParams struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (s *AuthService) RegisterUser(ctx context.Context, payload RegisterUserParams) error {
	fmt.Println("here")
	validationErrors := validation.Validate(payload)
	fmt.Println("validation", validationErrors)
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

func (s *AuthService) LoginUser(ctx context.Context, payload LoginUserParams) (string, error) {
	// Validate Payload
	validationErrors := validation.Validate(payload)
	if validationErrors != nil {
		return "", validationErrors
	}

	// Check User Exist
	user, err := s.queries.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		//TODO: pass special error
		return "", err
	}

	pass, err := s.queries.GetPasswordForUser(ctx, user.ID)

	if err != nil {
		//TODO: pass special error
		return "", err
	}

	// Validate Password
	err = bcrypt.CompareHashAndPassword([]byte(pass.HashedPassword), []byte(payload.Password))
	if err != nil {
		//TODO: pass special error
		return "", err
	}
	// Sign JWT
	token, err := jwtx.Sign(jwtx.Payload{
		UserID: string(user.ID.Bytes[:]),
		Email:  user.Email,
	}, "random string")
	return token, err
}
