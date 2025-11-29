package db

import (
	"context"
	"fmt"
	"skeleton-test/internal/sqlc"
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/jackc/pgx/v5"
)

const wait = 30 * time.Second

// implements Database
type Postgres struct {
	uri  string
	conn *pgx.Conn
}

func NewPostgres(uri string) Database {
	return &Postgres{
		uri: uri,
	}
}

func (p *Postgres) Connect() error {
	connectionCtx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	conn, err := pgx.Connect(connectionCtx, p.uri)
	if err != nil {
		return fmt.Errorf("failed to connect to the database : %v", err)
	}
	p.conn = conn
	return nil
}
func (p *Postgres) Disconnect() error {
	if p.conn == nil {
		return NotInitializedErr("Postgres")
	}
	disconnectionCtx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	err := p.conn.Close(disconnectionCtx)
	if err != nil {
		return fmt.Errorf("failed to disconnect the database connection: %v", err)
	}
	return nil
}
func (p *Postgres) Health() error {
	if p.conn == nil {
		return NotInitializedErr("Postgres")
	}
	pingCtx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	err := p.conn.Ping(pingCtx)
	if err != nil {
		return fmt.Errorf("failed to ping the database: %v", err)
	}
	return nil
}

func (p *Postgres) Conn() (sqlc.DBTX, error) {
	if p.conn == nil {
		return nil, NotInitializedErr("Postgres")
	}
	return p.conn, nil
}
