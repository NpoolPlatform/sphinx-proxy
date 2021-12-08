// Code generated by entc, DO NOT EDIT.

package transaction

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Nonce applies equality check predicate on the "nonce" field. It's identical to NonceEQ.
func Nonce(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// TransactionType applies equality check predicate on the "transaction_type" field. It's identical to TransactionTypeEQ.
func TransactionType(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionType), v))
	})
}

// CoinType applies equality check predicate on the "coin_type" field. It's identical to CoinTypeEQ.
func CoinType(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// TransactionID applies equality check predicate on the "transaction_id" field. It's identical to TransactionIDEQ.
func TransactionID(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// Cid applies equality check predicate on the "cid" field. It's identical to CidEQ.
func Cid(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCid), v))
	})
}

// ExitCode applies equality check predicate on the "exit_code" field. It's identical to ExitCodeEQ.
func ExitCode(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExitCode), v))
	})
}

// From applies equality check predicate on the "from" field. It's identical to FromEQ.
func From(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// To applies equality check predicate on the "to" field. It's identical to ToEQ.
func To(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// NonceEQ applies the EQ predicate on the "nonce" field.
func NonceEQ(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNonce), v))
	})
}

// NonceNEQ applies the NEQ predicate on the "nonce" field.
func NonceNEQ(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNonce), v))
	})
}

// NonceIn applies the In predicate on the "nonce" field.
func NonceIn(vs ...uint64) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNonce), v...))
	})
}

// NonceNotIn applies the NotIn predicate on the "nonce" field.
func NonceNotIn(vs ...uint64) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNonce), v...))
	})
}

// NonceGT applies the GT predicate on the "nonce" field.
func NonceGT(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNonce), v))
	})
}

// NonceGTE applies the GTE predicate on the "nonce" field.
func NonceGTE(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNonce), v))
	})
}

// NonceLT applies the LT predicate on the "nonce" field.
func NonceLT(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNonce), v))
	})
}

// NonceLTE applies the LTE predicate on the "nonce" field.
func NonceLTE(v uint64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNonce), v))
	})
}

// TransactionTypeEQ applies the EQ predicate on the "transaction_type" field.
func TransactionTypeEQ(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeNEQ applies the NEQ predicate on the "transaction_type" field.
func TransactionTypeNEQ(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeIn applies the In predicate on the "transaction_type" field.
func TransactionTypeIn(vs ...int8) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransactionType), v...))
	})
}

// TransactionTypeNotIn applies the NotIn predicate on the "transaction_type" field.
func TransactionTypeNotIn(vs ...int8) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransactionType), v...))
	})
}

// TransactionTypeGT applies the GT predicate on the "transaction_type" field.
func TransactionTypeGT(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeGTE applies the GTE predicate on the "transaction_type" field.
func TransactionTypeGTE(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeLT applies the LT predicate on the "transaction_type" field.
func TransactionTypeLT(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeLTE applies the LTE predicate on the "transaction_type" field.
func TransactionTypeLTE(v int8) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionType), v))
	})
}

// CoinTypeEQ applies the EQ predicate on the "coin_type" field.
func CoinTypeEQ(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeNEQ applies the NEQ predicate on the "coin_type" field.
func CoinTypeNEQ(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinType), v))
	})
}

// CoinTypeIn applies the In predicate on the "coin_type" field.
func CoinTypeIn(vs ...int32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinType), v...))
	})
}

// CoinTypeNotIn applies the NotIn predicate on the "coin_type" field.
func CoinTypeNotIn(vs ...int32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinType), v...))
	})
}

// CoinTypeGT applies the GT predicate on the "coin_type" field.
func CoinTypeGT(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinType), v))
	})
}

// CoinTypeGTE applies the GTE predicate on the "coin_type" field.
func CoinTypeGTE(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinType), v))
	})
}

// CoinTypeLT applies the LT predicate on the "coin_type" field.
func CoinTypeLT(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinType), v))
	})
}

// CoinTypeLTE applies the LTE predicate on the "coin_type" field.
func CoinTypeLTE(v int32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinType), v))
	})
}

// TransactionIDEQ applies the EQ predicate on the "transaction_id" field.
func TransactionIDEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDNEQ applies the NEQ predicate on the "transaction_id" field.
func TransactionIDNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionID), v))
	})
}

// TransactionIDIn applies the In predicate on the "transaction_id" field.
func TransactionIDIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDNotIn applies the NotIn predicate on the "transaction_id" field.
func TransactionIDNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransactionID), v...))
	})
}

// TransactionIDGT applies the GT predicate on the "transaction_id" field.
func TransactionIDGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDGTE applies the GTE predicate on the "transaction_id" field.
func TransactionIDGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLT applies the LT predicate on the "transaction_id" field.
func TransactionIDLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionID), v))
	})
}

// TransactionIDLTE applies the LTE predicate on the "transaction_id" field.
func TransactionIDLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContains applies the Contains predicate on the "transaction_id" field.
func TransactionIDContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasPrefix applies the HasPrefix predicate on the "transaction_id" field.
func TransactionIDHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDHasSuffix applies the HasSuffix predicate on the "transaction_id" field.
func TransactionIDHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionID), v))
	})
}

// TransactionIDEqualFold applies the EqualFold predicate on the "transaction_id" field.
func TransactionIDEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionID), v))
	})
}

// TransactionIDContainsFold applies the ContainsFold predicate on the "transaction_id" field.
func TransactionIDContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionID), v))
	})
}

// CidEQ applies the EQ predicate on the "cid" field.
func CidEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCid), v))
	})
}

// CidNEQ applies the NEQ predicate on the "cid" field.
func CidNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCid), v))
	})
}

// CidIn applies the In predicate on the "cid" field.
func CidIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCid), v...))
	})
}

// CidNotIn applies the NotIn predicate on the "cid" field.
func CidNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCid), v...))
	})
}

// CidGT applies the GT predicate on the "cid" field.
func CidGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCid), v))
	})
}

// CidGTE applies the GTE predicate on the "cid" field.
func CidGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCid), v))
	})
}

// CidLT applies the LT predicate on the "cid" field.
func CidLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCid), v))
	})
}

// CidLTE applies the LTE predicate on the "cid" field.
func CidLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCid), v))
	})
}

// CidContains applies the Contains predicate on the "cid" field.
func CidContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCid), v))
	})
}

// CidHasPrefix applies the HasPrefix predicate on the "cid" field.
func CidHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCid), v))
	})
}

// CidHasSuffix applies the HasSuffix predicate on the "cid" field.
func CidHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCid), v))
	})
}

// CidEqualFold applies the EqualFold predicate on the "cid" field.
func CidEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCid), v))
	})
}

// CidContainsFold applies the ContainsFold predicate on the "cid" field.
func CidContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCid), v))
	})
}

// ExitCodeEQ applies the EQ predicate on the "exit_code" field.
func ExitCodeEQ(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExitCode), v))
	})
}

// ExitCodeNEQ applies the NEQ predicate on the "exit_code" field.
func ExitCodeNEQ(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExitCode), v))
	})
}

// ExitCodeIn applies the In predicate on the "exit_code" field.
func ExitCodeIn(vs ...int64) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExitCode), v...))
	})
}

// ExitCodeNotIn applies the NotIn predicate on the "exit_code" field.
func ExitCodeNotIn(vs ...int64) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExitCode), v...))
	})
}

// ExitCodeGT applies the GT predicate on the "exit_code" field.
func ExitCodeGT(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExitCode), v))
	})
}

// ExitCodeGTE applies the GTE predicate on the "exit_code" field.
func ExitCodeGTE(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExitCode), v))
	})
}

// ExitCodeLT applies the LT predicate on the "exit_code" field.
func ExitCodeLT(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExitCode), v))
	})
}

// ExitCodeLTE applies the LTE predicate on the "exit_code" field.
func ExitCodeLTE(v int64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExitCode), v))
	})
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFrom), v))
	})
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFrom), v...))
	})
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFrom), v...))
	})
}

// FromGT applies the GT predicate on the "from" field.
func FromGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFrom), v))
	})
}

// FromGTE applies the GTE predicate on the "from" field.
func FromGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFrom), v))
	})
}

// FromLT applies the LT predicate on the "from" field.
func FromLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFrom), v))
	})
}

// FromLTE applies the LTE predicate on the "from" field.
func FromLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFrom), v))
	})
}

// FromContains applies the Contains predicate on the "from" field.
func FromContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFrom), v))
	})
}

// FromHasPrefix applies the HasPrefix predicate on the "from" field.
func FromHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFrom), v))
	})
}

// FromHasSuffix applies the HasSuffix predicate on the "from" field.
func FromHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFrom), v))
	})
}

// FromEqualFold applies the EqualFold predicate on the "from" field.
func FromEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFrom), v))
	})
}

// FromContainsFold applies the ContainsFold predicate on the "from" field.
func FromContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFrom), v))
	})
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTo), v))
	})
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTo), v...))
	})
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...string) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTo), v...))
	})
}

// ToGT applies the GT predicate on the "to" field.
func ToGT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTo), v))
	})
}

// ToGTE applies the GTE predicate on the "to" field.
func ToGTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTo), v))
	})
}

// ToLT applies the LT predicate on the "to" field.
func ToLT(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTo), v))
	})
}

// ToLTE applies the LTE predicate on the "to" field.
func ToLTE(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTo), v))
	})
}

// ToContains applies the Contains predicate on the "to" field.
func ToContains(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTo), v))
	})
}

// ToHasPrefix applies the HasPrefix predicate on the "to" field.
func ToHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTo), v))
	})
}

// ToHasSuffix applies the HasSuffix predicate on the "to" field.
func ToHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTo), v))
	})
}

// ToEqualFold applies the EqualFold predicate on the "to" field.
func ToEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTo), v))
	})
}

// ToContainsFold applies the ContainsFold predicate on the "to" field.
func ToContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTo), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...float64) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...float64) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v float64) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Transaction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transaction(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		p(s.Not())
	})
}
