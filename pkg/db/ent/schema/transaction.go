package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/id"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(id.ID).
			Unique(),
		field.Uint64("nonce").
			Default(0),
		field.Int8("transaction_type").
			Default(0),
		field.Int8("coin_type").
			Default(0),
		field.String("transaction_id_insite").
			NotEmpty().
			Default(""),
		field.String("from").
			NotEmpty().
			Default(""),
		field.String("to").
			NotEmpty().
			Default(""),
		field.Float("value").
			Default(0),
		field.Enum("state").
			Values("wait", "done", "fail"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}
