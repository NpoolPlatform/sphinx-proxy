// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	"github.com/google/uuid"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CoinType holds the value of the "coin_type" field.
	CoinType int32 `json:"coin_type,omitempty"`
	// --will remove
	Nonce uint64 `json:"nonce,omitempty"`
	// only for btc--will remove
	Utxo []*sphinxplugin.Unspent `json:"utxo,omitempty"`
	// --will remove
	Pre *sphinxplugin.Unspent `json:"pre,omitempty"`
	// --will remove
	TransactionType int8 `json:"transaction_type,omitempty"`
	// --will remove
	RecentBhash string `json:"recent_bhash,omitempty"`
	// --will remove
	TxData []byte `json:"tx_data,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID string `json:"transaction_id,omitempty"`
	// Cid holds the value of the "cid" field.
	Cid string `json:"cid,omitempty"`
	// ExitCode holds the value of the "exit_code" field.
	ExitCode int64 `json:"exit_code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// save nonce or sign info
	Payload []byte `json:"payload,omitempty"`
	// State holds the value of the "state" field.
	State uint8 `json:"state,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldUtxo, transaction.FieldPre, transaction.FieldTxData, transaction.FieldPayload:
			values[i] = new([]byte)
		case transaction.FieldCoinType, transaction.FieldNonce, transaction.FieldTransactionType, transaction.FieldExitCode, transaction.FieldAmount, transaction.FieldState, transaction.FieldCreatedAt, transaction.FieldUpdatedAt, transaction.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case transaction.FieldRecentBhash, transaction.FieldTransactionID, transaction.FieldCid, transaction.FieldName, transaction.FieldFrom, transaction.FieldTo:
			values[i] = new(sql.NullString)
		case transaction.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case transaction.FieldCoinType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type", values[i])
			} else if value.Valid {
				t.CoinType = int32(value.Int64)
			}
		case transaction.FieldNonce:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				t.Nonce = uint64(value.Int64)
			}
		case transaction.FieldUtxo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field utxo", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Utxo); err != nil {
					return fmt.Errorf("unmarshal field utxo: %w", err)
				}
			}
		case transaction.FieldPre:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field pre", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Pre); err != nil {
					return fmt.Errorf("unmarshal field pre: %w", err)
				}
			}
		case transaction.FieldTransactionType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_type", values[i])
			} else if value.Valid {
				t.TransactionType = int8(value.Int64)
			}
		case transaction.FieldRecentBhash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recent_bhash", values[i])
			} else if value.Valid {
				t.RecentBhash = value.String
			}
		case transaction.FieldTxData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tx_data", values[i])
			} else if value != nil {
				t.TxData = *value
			}
		case transaction.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				t.TransactionID = value.String
			}
		case transaction.FieldCid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cid", values[i])
			} else if value.Valid {
				t.Cid = value.String
			}
		case transaction.FieldExitCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exit_code", values[i])
			} else if value.Valid {
				t.ExitCode = value.Int64
			}
		case transaction.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case transaction.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				t.From = value.String
			}
		case transaction.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				t.To = value.String
			}
		case transaction.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = uint64(value.Int64)
			}
		case transaction.FieldPayload:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field payload", values[i])
			} else if value != nil {
				t.Payload = *value
			}
		case transaction.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				t.State = uint8(value.Int64)
			}
		case transaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = uint32(value.Int64)
			}
		case transaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = uint32(value.Int64)
			}
		case transaction.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("coin_type=")
	builder.WriteString(fmt.Sprintf("%v", t.CoinType))
	builder.WriteString(", ")
	builder.WriteString("nonce=")
	builder.WriteString(fmt.Sprintf("%v", t.Nonce))
	builder.WriteString(", ")
	builder.WriteString("utxo=")
	builder.WriteString(fmt.Sprintf("%v", t.Utxo))
	builder.WriteString(", ")
	builder.WriteString("pre=")
	builder.WriteString(fmt.Sprintf("%v", t.Pre))
	builder.WriteString(", ")
	builder.WriteString("transaction_type=")
	builder.WriteString(fmt.Sprintf("%v", t.TransactionType))
	builder.WriteString(", ")
	builder.WriteString("recent_bhash=")
	builder.WriteString(t.RecentBhash)
	builder.WriteString(", ")
	builder.WriteString("tx_data=")
	builder.WriteString(fmt.Sprintf("%v", t.TxData))
	builder.WriteString(", ")
	builder.WriteString("transaction_id=")
	builder.WriteString(t.TransactionID)
	builder.WriteString(", ")
	builder.WriteString("cid=")
	builder.WriteString(t.Cid)
	builder.WriteString(", ")
	builder.WriteString("exit_code=")
	builder.WriteString(fmt.Sprintf("%v", t.ExitCode))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(t.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(t.To)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("payload=")
	builder.WriteString(fmt.Sprintf("%v", t.Payload))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", t.State))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", t.DeletedAt))
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
