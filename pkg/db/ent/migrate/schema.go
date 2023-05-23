// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "nonce", Type: field.TypeUint64, Nullable: true, Default: 0},
		{Name: "utxo", Type: field.TypeJSON, Nullable: true},
		{Name: "pre", Type: field.TypeJSON, Nullable: true},
		{Name: "transaction_type", Type: field.TypeInt8, Nullable: true, Default: 0},
		{Name: "recent_bhash", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "tx_data", Type: field.TypeBytes, Nullable: true},
		{Name: "transaction_id", Type: field.TypeString, Unique: true},
		{Name: "cid", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "exit_code", Type: field.TypeInt64, Nullable: true, Default: -1},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "from", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "to", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "memo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "amount", Type: field.TypeUint64, Nullable: true, Default: 0},
		{Name: "payload", Type: field.TypeBytes, Nullable: true, Size: 4294967295},
		{Name: "state", Type: field.TypeUint8, Nullable: true, Default: 0},
		{Name: "created_at", Type: field.TypeUint32, Nullable: true},
		{Name: "updated_at", Type: field.TypeUint32, Nullable: true},
		{Name: "deleted_at", Type: field.TypeUint32, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_state_coin_type_created_at",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[17], TransactionsColumns[1], TransactionsColumns[18]},
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
