package database

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/balcieren/fiber-boilerplate/internal/config"
	"github.com/balcieren/fiber-boilerplate/pkg/ent"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewPostgreSQL(config *config.Config) (*ent.Client, error) {
	dbEnt, err := sql.Open("pgx", config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, dbEnt)
	client := ent.NewClient(ent.Driver(drv))
	err = client.Schema.Create(context.Background(), schema.WithAtlas(true))

	return client, err
}
