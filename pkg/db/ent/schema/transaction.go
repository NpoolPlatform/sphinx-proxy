package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/plugin/eth"
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
		field.Uint64("nonce").
			Default(0),
		field.JSON("utxo", []*sphinxplugin.Unspent{}).
			Default([]*sphinxplugin.Unspent{}).
			Comment("only for btc"),
		field.JSON("pre", &eth.PreSignInfo{}).
			Default(&eth.PreSignInfo{}).
			Comment("only for eth"),
		field.Int8("transaction_type").
			Default(0),
		field.Int32("coin_type").
			Default(0),
		field.String("transaction_id").
			Unique().
			NotEmpty(),
		field.String("recent_bhash").
			Default("").
			Comment("only for sol,abbreviation for 'recent_bloack_hash'"),
		field.Bytes("tx_data").
			Default([]byte{}).
			Comment("tx_data of complex data structure for tx"),
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
		// index.Fields("from", "nonce").
		// 	Unique(),
		index.Fields("created_at"),
	}
}
