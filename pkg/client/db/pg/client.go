package pg

import (
	"context"

	"github.com/drizzleent/ton-bot/pkg/client/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgclient struct {
	masterDbc db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect db: %v", err)
	}

	return &pgclient{
		masterDbc: &pg{
			dbc: dbc,
		},
	}, nil
}

func (c *pgclient) DB() db.DB {
	return c.masterDbc
}

func (c *pgclient) Close() error {
	if c.masterDbc != nil {
		c.masterDbc.Close()
	}
	return nil
}
