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
		{Name: "transaction_type", Type: field.TypeInt8, Default: 0},
		{Name: "coin_type", Type: field.TypeInt32, Default: 0},
		{Name: "transaction_id", Type: field.TypeString, Default: ""},
		{Name: "cid", Type: field.TypeString, Default: ""},
		{Name: "from", Type: field.TypeString, Default: ""},
		{Name: "to", Type: field.TypeString, Default: ""},
		{Name: "value", Type: field.TypeFloat64, Default: 0},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"wait", "sign", "done", "fail"}},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TransactionsTable,
	}
)

func init() {
}
