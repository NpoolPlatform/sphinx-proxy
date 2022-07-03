// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/schema"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescCoinType is the schema descriptor for coin_type field.
	transactionDescCoinType := transactionFields[1].Descriptor()
	// transaction.DefaultCoinType holds the default value on creation for the coin_type field.
	transaction.DefaultCoinType = transactionDescCoinType.Default.(int32)
	// transactionDescTransactionID is the schema descriptor for transaction_id field.
	transactionDescTransactionID := transactionFields[2].Descriptor()
	// transaction.TransactionIDValidator is a validator for the "transaction_id" field. It is called by the builders before save.
	transaction.TransactionIDValidator = transactionDescTransactionID.Validators[0].(func(string) error)
	// transactionDescCid is the schema descriptor for cid field.
	transactionDescCid := transactionFields[3].Descriptor()
	// transaction.DefaultCid holds the default value on creation for the cid field.
	transaction.DefaultCid = transactionDescCid.Default.(string)
	// transactionDescExitCode is the schema descriptor for exit_code field.
	transactionDescExitCode := transactionFields[4].Descriptor()
	// transaction.DefaultExitCode holds the default value on creation for the exit_code field.
	transaction.DefaultExitCode = transactionDescExitCode.Default.(int64)
	// transactionDescFrom is the schema descriptor for from field.
	transactionDescFrom := transactionFields[5].Descriptor()
	// transaction.DefaultFrom holds the default value on creation for the from field.
	transaction.DefaultFrom = transactionDescFrom.Default.(string)
	// transaction.FromValidator is a validator for the "from" field. It is called by the builders before save.
	transaction.FromValidator = transactionDescFrom.Validators[0].(func(string) error)
	// transactionDescTo is the schema descriptor for to field.
	transactionDescTo := transactionFields[6].Descriptor()
	// transaction.DefaultTo holds the default value on creation for the to field.
	transaction.DefaultTo = transactionDescTo.Default.(string)
	// transaction.ToValidator is a validator for the "to" field. It is called by the builders before save.
	transaction.ToValidator = transactionDescTo.Validators[0].(func(string) error)
	// transactionDescAmount is the schema descriptor for amount field.
	transactionDescAmount := transactionFields[7].Descriptor()
	// transaction.DefaultAmount holds the default value on creation for the amount field.
	transaction.DefaultAmount = transactionDescAmount.Default.(uint64)
	// transaction.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	transaction.AmountValidator = transactionDescAmount.Validators[0].(func(uint64) error)
	// transactionDescPayload is the schema descriptor for payload field.
	transactionDescPayload := transactionFields[8].Descriptor()
	// transaction.DefaultPayload holds the default value on creation for the payload field.
	transaction.DefaultPayload = transactionDescPayload.Default.([]byte)
	// transactionDescState is the schema descriptor for state field.
	transactionDescState := transactionFields[9].Descriptor()
	// transaction.DefaultState holds the default value on creation for the state field.
	transaction.DefaultState = transactionDescState.Default.(uint8)
	// transactionDescCreatedAt is the schema descriptor for created_at field.
	transactionDescCreatedAt := transactionFields[10].Descriptor()
	// transaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	transaction.DefaultCreatedAt = transactionDescCreatedAt.Default.(func() uint32)
	// transactionDescUpdatedAt is the schema descriptor for updated_at field.
	transactionDescUpdatedAt := transactionFields[11].Descriptor()
	// transaction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	transaction.DefaultUpdatedAt = transactionDescUpdatedAt.Default.(func() uint32)
	// transaction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	transaction.UpdateDefaultUpdatedAt = transactionDescUpdatedAt.UpdateDefault.(func() uint32)
	// transactionDescDeletedAt is the schema descriptor for deleted_at field.
	transactionDescDeletedAt := transactionFields[12].Descriptor()
	// transaction.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	transaction.DefaultDeletedAt = transactionDescDeletedAt.Default.(func() uint32)
	// transactionDescID is the schema descriptor for id field.
	transactionDescID := transactionFields[0].Descriptor()
	// transaction.DefaultID holds the default value on creation for the id field.
	transaction.DefaultID = transactionDescID.Default.(func() uuid.UUID)
}
