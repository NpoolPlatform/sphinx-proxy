// Code generated by entc, DO NOT EDIT.

package transaction

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldUtxo holds the string denoting the utxo field in the database.
	FieldUtxo = "utxo"
	// FieldTransactionType holds the string denoting the transaction_type field in the database.
	FieldTransactionType = "transaction_type"
	// FieldCoinType holds the string denoting the coin_type field in the database.
	FieldCoinType = "coin_type"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldCid holds the string denoting the cid field in the database.
	FieldCid = "cid"
	// FieldExitCode holds the string denoting the exit_code field in the database.
	FieldExitCode = "exit_code"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
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
	FieldUtxo,
	FieldTransactionType,
	FieldCoinType,
	FieldTransactionID,
	FieldCid,
	FieldExitCode,
	FieldFrom,
	FieldTo,
	FieldAmount,
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
	// DefaultExitCode holds the default value on creation for the "exit_code" field.
	DefaultExitCode int64
	// DefaultFrom holds the default value on creation for the "from" field.
	DefaultFrom string
	// FromValidator is a validator for the "from" field. It is called by the builders before save.
	FromValidator func(string) error
	// DefaultTo holds the default value on creation for the "to" field.
	DefaultTo string
	// ToValidator is a validator for the "to" field. It is called by the builders before save.
	ToValidator func(string) error
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount uint64
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(uint64) error
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState uint8
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
