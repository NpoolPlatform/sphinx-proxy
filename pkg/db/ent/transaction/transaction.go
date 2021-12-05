// Code generated by entc, DO NOT EDIT.

package transaction

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
	// FieldCoinType holds the string denoting the coin_type field in the database.
	FieldCoinType = "coin_type"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldCid holds the string denoting the cid field in the database.
	FieldCid = "cid"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldNonce,
	FieldTransactionType,
	FieldCoinType,
	FieldTransactionID,
	FieldCid,
	FieldFrom,
	FieldTo,
	FieldValue,
	FieldState,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultNonce holds the default value on creation for the "nonce" field.
	DefaultNonce uint64
	// DefaultTransactionType holds the default value on creation for the "transaction_type" field.
	DefaultTransactionType int8
	// DefaultCoinType holds the default value on creation for the "coin_type" field.
	DefaultCoinType int32
	// TransactionIDValidator is a validator for the "transaction_id" field. It is called by the builders before save.
	TransactionIDValidator func(string) error
	// DefaultCid holds the default value on creation for the "cid" field.
	DefaultCid string
	// CidValidator is a validator for the "cid" field. It is called by the builders before save.
	CidValidator func(string) error
	// DefaultFrom holds the default value on creation for the "from" field.
	DefaultFrom string
	// FromValidator is a validator for the "from" field. It is called by the builders before save.
	FromValidator func(string) error
	// DefaultTo holds the default value on creation for the "to" field.
	DefaultTo string
	// ToValidator is a validator for the "to" field. It is called by the builders before save.
	ToValidator func(string) error
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue float64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateWait State = "wait"
	StateSign State = "sign"
	StateDone State = "done"
	StateFail State = "fail"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateWait, StateSign, StateDone, StateFail:
		return nil
	default:
		return fmt.Errorf("transaction: invalid enum value for state field: %q", s)
	}
}
