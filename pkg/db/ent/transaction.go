// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce uint64 `json:"nonce,omitempty"`
	// TransactionType holds the value of the "transaction_type" field.
	TransactionType int8 `json:"transaction_type,omitempty"`
	// CoinType holds the value of the "coin_type" field.
	CoinType int8 `json:"coin_type,omitempty"`
	// TransactionIDInsite holds the value of the "transaction_id_insite" field.
	TransactionIDInsite string `json:"transaction_id_insite,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// State holds the value of the "state" field.
	State transaction.State `json:"state,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldValue:
			values[i] = new(sql.NullFloat64)
		case transaction.FieldNonce, transaction.FieldTransactionType, transaction.FieldCoinType, transaction.FieldCreateAt, transaction.FieldUpdateAt, transaction.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case transaction.FieldID, transaction.FieldTransactionIDInsite, transaction.FieldFrom, transaction.FieldTo, transaction.FieldState:
			values[i] = new(sql.NullString)
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
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case transaction.FieldNonce:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				t.Nonce = uint64(value.Int64)
			}
		case transaction.FieldTransactionType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_type", values[i])
			} else if value.Valid {
				t.TransactionType = int8(value.Int64)
			}
		case transaction.FieldCoinType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type", values[i])
			} else if value.Valid {
				t.CoinType = int8(value.Int64)
			}
		case transaction.FieldTransactionIDInsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id_insite", values[i])
			} else if value.Valid {
				t.TransactionIDInsite = value.String
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
		case transaction.FieldValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				t.Value = value.Float64
			}
		case transaction.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				t.State = transaction.State(value.String)
			}
		case transaction.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				t.CreateAt = uint32(value.Int64)
			}
		case transaction.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				t.UpdateAt = uint32(value.Int64)
			}
		case transaction.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				t.DeleteAt = uint32(value.Int64)
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
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", nonce=")
	builder.WriteString(fmt.Sprintf("%v", t.Nonce))
	builder.WriteString(", transaction_type=")
	builder.WriteString(fmt.Sprintf("%v", t.TransactionType))
	builder.WriteString(", coin_type=")
	builder.WriteString(fmt.Sprintf("%v", t.CoinType))
	builder.WriteString(", transaction_id_insite=")
	builder.WriteString(t.TransactionIDInsite)
	builder.WriteString(", from=")
	builder.WriteString(t.From)
	builder.WriteString(", to=")
	builder.WriteString(t.To)
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", t.Value))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", t.State))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", t.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", t.DeleteAt))
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
