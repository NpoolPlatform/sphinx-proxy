package schema

import (
	"math"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

func (Transaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("coin_type").
			Optional().
			Default(0),
		field.Uint64("nonce").
			Optional().
			Default(0).
			Comment("--will remove"),
		field.JSON("utxo", []*sphinxplugin.Unspent{}).
			Optional().
			Default([]*sphinxplugin.Unspent{}).
			Comment("only for btc--will remove"),
		field.JSON("pre", &sphinxplugin.Unspent{}).
			Optional().
			Default(&sphinxplugin.Unspent{}).
			Comment("--will remove"),
		field.Int8("transaction_type").
			Optional().
			Default(0).
			Comment("--will remove"),
		field.String("recent_bhash").
			Optional().
			Default("").
			Comment("--will remove"),
		field.Bytes("tx_data").
			Optional().
			Default([]byte{}).
			Comment("--will remove"),
		field.String("transaction_id").
			Unique(),
		field.String("cid").
			Optional().
			Default(""),
		field.Int64("exit_code").
			Optional().
			Default(-1),
		field.String("name").
			Optional().
			Default(""),
		field.String("from").
			Optional().
			Default(""),
		field.String("to").
			Optional().
			Default(""),
		field.String("memo").
			Optional().
			Default(""),
		field.Uint64("amount").
			Optional().
			Default(0),
		field.Bytes("payload").
			Optional().
			MaxLen(math.MaxUint32).
			Default([]byte{}).
			Comment("save nonce or sign info"),
		field.Uint8("state").
			Optional().
			Default(0),
		field.Uint32("created_at").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("updated_at").
			Optional().
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("deleted_at").
			Optional().
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

func (Transaction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(
			transaction.FieldState,
			transaction.FieldCoinType,
			transaction.FieldCreatedAt,
		),
	}
}
