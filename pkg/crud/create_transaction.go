package crud

import "github.com/NpoolPlatform/sphinx-proxy/pkg/db"

func CreateTransaction() error {
	db.Client().Transaction.CreateBulk()

	return nil
}
