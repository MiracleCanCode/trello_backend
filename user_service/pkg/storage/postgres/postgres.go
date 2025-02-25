package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	dsn string
}

func New(dsn string) *Postgres {
	return &Postgres{
		dsn: dsn,
	}
}

func (p *Postgres) Conn(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, p.dsn)
	if err != nil {
		return nil, fmt.Errorf("Conn: failed create connect to postgres db: %w", err)
	}

	return conn, nil
}
