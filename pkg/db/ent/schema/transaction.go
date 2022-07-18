package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int32("coin_type").
			Default(0),
		field.String("transaction_id").
			Unique().
			NotEmpty(),
		field.String("cid").
			Default(""),
		field.Int64("exit_code").
			Default(-1),
		field.String("from").
			NotEmpty().
			Default(""),
		field.String("to").
			NotEmpty().
			Default(""),
		field.Uint64("amount").
			Positive().
			Default(0),
		field.Bytes("payload").
			Default([]byte{}).
			Comment("save nonce or sign info"),
		field.Uint8("state").Default(0),
		field.Uint32("created_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
