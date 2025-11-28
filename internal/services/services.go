package services

import (
	"skeleton-test/internal/sqlc"
)

type Services struct {
	Auth *AuthService
}

func New(db sqlc.DBTX) *Services {
	queries := sqlc.New(db)
	return &Services{
		Auth: &AuthService{
			queries: queries,
		},
	}
}
