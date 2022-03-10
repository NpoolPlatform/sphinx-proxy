// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "nonce", Type: field.TypeUint64, Default: 0},
		{Name: "utxo", Type: field.TypeJSON},
		{Name: "transaction_type", Type: field.TypeInt8, Default: 0},
		{Name: "coin_type", Type: field.TypeInt32, Default: 0},
		{Name: "transaction_id", Type: field.TypeString, Unique: true},
		{Name: "cid", Type: field.TypeString, Default: ""},
		{Name: "exit_code", Type: field.TypeInt64, Default: -1},
		{Name: "from", Type: field.TypeString, Default: ""},
		{Name: "to", Type: field.TypeString, Default: ""},
		{Name: "amount", Type: field.TypeUint64, Default: 0},
		{Name: "state", Type: field.TypeUint8, Default: 0},
		{Name: "created_at", Type: field.TypeUint32, Default: 0},
		{Name: "updated_at", Type: field.TypeUint32, Default: 0},
		{Name: "deleted_at", Type: field.TypeUint32, Default: 0},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_created_at",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[12]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TransactionsTable,
	}
)

func init() {
}
