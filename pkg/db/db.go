package db

import (
	"context"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
)

func Client() (*ent.Client, error) {
	return client()
}

func client() (*ent.Client, error) {
	conn, err := mysql.GetConn()
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init() error {
	c, err := client()
	if err != nil {
		return err
	}
	return c.
		Schema.Create(context.Background())
}
