package db

import (
	"fmt"
	"skeleton-test/internal/sqlc"
)

type Database interface {
	Connect() error
	Disconnect() error
	Health() error
	Conn() (sqlc.DBTX, error)
}

func NotInitializedErr(what string) error {
	return fmt.Errorf("%s is not initialized, have you forgot to call Connect() ?", what)
}
