// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetNonce sets the "nonce" field.
func (tu *TransactionUpdate) SetNonce(u uint64) *TransactionUpdate {
	tu.mutation.ResetNonce()
	tu.mutation.SetNonce(u)
	return tu
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableNonce(u *uint64) *TransactionUpdate {
	if u != nil {
		tu.SetNonce(*u)
	}
	return tu
}

// AddNonce adds u to the "nonce" field.
func (tu *TransactionUpdate) AddNonce(u int64) *TransactionUpdate {
	tu.mutation.AddNonce(u)
	return tu
}

// SetUtxo sets the "utxo" field.
func (tu *TransactionUpdate) SetUtxo(s []*sphinxplugin.Unspent) *TransactionUpdate {
	tu.mutation.SetUtxo(s)
	return tu
}

// SetTransactionType sets the "transaction_type" field.
func (tu *TransactionUpdate) SetTransactionType(i int8) *TransactionUpdate {
	tu.mutation.ResetTransactionType()
	tu.mutation.SetTransactionType(i)
	return tu
}

// SetNillableTransactionType sets the "transaction_type" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTransactionType(i *int8) *TransactionUpdate {
	if i != nil {
		tu.SetTransactionType(*i)
	}
	return tu
}

// AddTransactionType adds i to the "transaction_type" field.
func (tu *TransactionUpdate) AddTransactionType(i int8) *TransactionUpdate {
	tu.mutation.AddTransactionType(i)
	return tu
}

// SetCoinType sets the "coin_type" field.
func (tu *TransactionUpdate) SetCoinType(i int32) *TransactionUpdate {
	tu.mutation.ResetCoinType()
	tu.mutation.SetCoinType(i)
	return tu
}

// SetNillableCoinType sets the "coin_type" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableCoinType(i *int32) *TransactionUpdate {
	if i != nil {
		tu.SetCoinType(*i)
	}
	return tu
}

// AddCoinType adds i to the "coin_type" field.
func (tu *TransactionUpdate) AddCoinType(i int32) *TransactionUpdate {
	tu.mutation.AddCoinType(i)
	return tu
}

// SetTransactionID sets the "transaction_id" field.
func (tu *TransactionUpdate) SetTransactionID(s string) *TransactionUpdate {
	tu.mutation.SetTransactionID(s)
	return tu
}

// SetCid sets the "cid" field.
func (tu *TransactionUpdate) SetCid(s string) *TransactionUpdate {
	tu.mutation.SetCid(s)
	return tu
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableCid(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetCid(*s)
	}
	return tu
}

// SetExitCode sets the "exit_code" field.
func (tu *TransactionUpdate) SetExitCode(i int64) *TransactionUpdate {
	tu.mutation.ResetExitCode()
	tu.mutation.SetExitCode(i)
	return tu
}

// SetNillableExitCode sets the "exit_code" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableExitCode(i *int64) *TransactionUpdate {
	if i != nil {
		tu.SetExitCode(*i)
	}
	return tu
}

// AddExitCode adds i to the "exit_code" field.
func (tu *TransactionUpdate) AddExitCode(i int64) *TransactionUpdate {
	tu.mutation.AddExitCode(i)
	return tu
}

// SetFrom sets the "from" field.
func (tu *TransactionUpdate) SetFrom(s string) *TransactionUpdate {
	tu.mutation.SetFrom(s)
	return tu
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableFrom(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetFrom(*s)
	}
	return tu
}

// SetTo sets the "to" field.
func (tu *TransactionUpdate) SetTo(s string) *TransactionUpdate {
	tu.mutation.SetTo(s)
	return tu
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTo(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetTo(*s)
	}
	return tu
}

// SetAmount sets the "amount" field.
func (tu *TransactionUpdate) SetAmount(u uint64) *TransactionUpdate {
	tu.mutation.ResetAmount()
	tu.mutation.SetAmount(u)
	return tu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableAmount(u *uint64) *TransactionUpdate {
	if u != nil {
		tu.SetAmount(*u)
	}
	return tu
}

// AddAmount adds u to the "amount" field.
func (tu *TransactionUpdate) AddAmount(u int64) *TransactionUpdate {
	tu.mutation.AddAmount(u)
	return tu
}

// SetState sets the "state" field.
func (tu *TransactionUpdate) SetState(u uint8) *TransactionUpdate {
	tu.mutation.ResetState()
	tu.mutation.SetState(u)
	return tu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableState(u *uint8) *TransactionUpdate {
	if u != nil {
		tu.SetState(*u)
	}
	return tu
}

// AddState adds u to the "state" field.
func (tu *TransactionUpdate) AddState(u int8) *TransactionUpdate {
	tu.mutation.AddState(u)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TransactionUpdate) SetCreatedAt(u uint32) *TransactionUpdate {
	tu.mutation.ResetCreatedAt()
	tu.mutation.SetCreatedAt(u)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableCreatedAt(u *uint32) *TransactionUpdate {
	if u != nil {
		tu.SetCreatedAt(*u)
	}
	return tu
}

// AddCreatedAt adds u to the "created_at" field.
func (tu *TransactionUpdate) AddCreatedAt(u int32) *TransactionUpdate {
	tu.mutation.AddCreatedAt(u)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TransactionUpdate) SetUpdatedAt(u uint32) *TransactionUpdate {
	tu.mutation.ResetUpdatedAt()
	tu.mutation.SetUpdatedAt(u)
	return tu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tu *TransactionUpdate) AddUpdatedAt(u int32) *TransactionUpdate {
	tu.mutation.AddUpdatedAt(u)
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TransactionUpdate) SetDeletedAt(u uint32) *TransactionUpdate {
	tu.mutation.ResetDeletedAt()
	tu.mutation.SetDeletedAt(u)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableDeletedAt(u *uint32) *TransactionUpdate {
	if u != nil {
		tu.SetDeletedAt(*u)
	}
	return tu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tu *TransactionUpdate) AddDeletedAt(u int32) *TransactionUpdate {
	tu.mutation.AddDeletedAt(u)
	return tu
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TransactionUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := transaction.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionUpdate) check() error {
	if v, ok := tu.mutation.TransactionID(); ok {
		if err := transaction.TransactionIDValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id", err: fmt.Errorf(`ent: validator failed for field "Transaction.transaction_id": %w`, err)}
		}
	}
	if v, ok := tu.mutation.From(); ok {
		if err := transaction.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "Transaction.from": %w`, err)}
		}
	}
	if v, ok := tu.mutation.To(); ok {
		if err := transaction.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "Transaction.to": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Amount(); ok {
		if err := transaction.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Transaction.amount": %w`, err)}
		}
	}
	return nil
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldNonce,
		})
	}
	if value, ok := tu.mutation.AddedNonce(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldNonce,
		})
	}
	if value, ok := tu.mutation.Utxo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transaction.FieldUtxo,
		})
	}
	if value, ok := tu.mutation.TransactionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: transaction.FieldTransactionType,
		})
	}
	if value, ok := tu.mutation.AddedTransactionType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: transaction.FieldTransactionType,
		})
	}
	if value, ok := tu.mutation.CoinType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: transaction.FieldCoinType,
		})
	}
	if value, ok := tu.mutation.AddedCoinType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: transaction.FieldCoinType,
		})
	}
	if value, ok := tu.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionID,
		})
	}
	if value, ok := tu.mutation.Cid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldCid,
		})
	}
	if value, ok := tu.mutation.ExitCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldExitCode,
		})
	}
	if value, ok := tu.mutation.AddedExitCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldExitCode,
		})
	}
	if value, ok := tu.mutation.From(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldFrom,
		})
	}
	if value, ok := tu.mutation.To(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTo,
		})
	}
	if value, ok := tu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmount,
		})
	}
	if value, ok := tu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmount,
		})
	}
	if value, ok := tu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: transaction.FieldState,
		})
	}
	if value, ok := tu.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: transaction.FieldState,
		})
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldDeletedAt,
		})
	}
	if value, ok := tu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetNonce sets the "nonce" field.
func (tuo *TransactionUpdateOne) SetNonce(u uint64) *TransactionUpdateOne {
	tuo.mutation.ResetNonce()
	tuo.mutation.SetNonce(u)
	return tuo
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableNonce(u *uint64) *TransactionUpdateOne {
	if u != nil {
		tuo.SetNonce(*u)
	}
	return tuo
}

// AddNonce adds u to the "nonce" field.
func (tuo *TransactionUpdateOne) AddNonce(u int64) *TransactionUpdateOne {
	tuo.mutation.AddNonce(u)
	return tuo
}

// SetUtxo sets the "utxo" field.
func (tuo *TransactionUpdateOne) SetUtxo(s []*sphinxplugin.Unspent) *TransactionUpdateOne {
	tuo.mutation.SetUtxo(s)
	return tuo
}

// SetTransactionType sets the "transaction_type" field.
func (tuo *TransactionUpdateOne) SetTransactionType(i int8) *TransactionUpdateOne {
	tuo.mutation.ResetTransactionType()
	tuo.mutation.SetTransactionType(i)
	return tuo
}

// SetNillableTransactionType sets the "transaction_type" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTransactionType(i *int8) *TransactionUpdateOne {
	if i != nil {
		tuo.SetTransactionType(*i)
	}
	return tuo
}

// AddTransactionType adds i to the "transaction_type" field.
func (tuo *TransactionUpdateOne) AddTransactionType(i int8) *TransactionUpdateOne {
	tuo.mutation.AddTransactionType(i)
	return tuo
}

// SetCoinType sets the "coin_type" field.
func (tuo *TransactionUpdateOne) SetCoinType(i int32) *TransactionUpdateOne {
	tuo.mutation.ResetCoinType()
	tuo.mutation.SetCoinType(i)
	return tuo
}

// SetNillableCoinType sets the "coin_type" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableCoinType(i *int32) *TransactionUpdateOne {
	if i != nil {
		tuo.SetCoinType(*i)
	}
	return tuo
}

// AddCoinType adds i to the "coin_type" field.
func (tuo *TransactionUpdateOne) AddCoinType(i int32) *TransactionUpdateOne {
	tuo.mutation.AddCoinType(i)
	return tuo
}

// SetTransactionID sets the "transaction_id" field.
func (tuo *TransactionUpdateOne) SetTransactionID(s string) *TransactionUpdateOne {
	tuo.mutation.SetTransactionID(s)
	return tuo
}

// SetCid sets the "cid" field.
func (tuo *TransactionUpdateOne) SetCid(s string) *TransactionUpdateOne {
	tuo.mutation.SetCid(s)
	return tuo
}

// SetNillableCid sets the "cid" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableCid(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetCid(*s)
	}
	return tuo
}

// SetExitCode sets the "exit_code" field.
func (tuo *TransactionUpdateOne) SetExitCode(i int64) *TransactionUpdateOne {
	tuo.mutation.ResetExitCode()
	tuo.mutation.SetExitCode(i)
	return tuo
}

// SetNillableExitCode sets the "exit_code" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableExitCode(i *int64) *TransactionUpdateOne {
	if i != nil {
		tuo.SetExitCode(*i)
	}
	return tuo
}

// AddExitCode adds i to the "exit_code" field.
func (tuo *TransactionUpdateOne) AddExitCode(i int64) *TransactionUpdateOne {
	tuo.mutation.AddExitCode(i)
	return tuo
}

// SetFrom sets the "from" field.
func (tuo *TransactionUpdateOne) SetFrom(s string) *TransactionUpdateOne {
	tuo.mutation.SetFrom(s)
	return tuo
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableFrom(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetFrom(*s)
	}
	return tuo
}

// SetTo sets the "to" field.
func (tuo *TransactionUpdateOne) SetTo(s string) *TransactionUpdateOne {
	tuo.mutation.SetTo(s)
	return tuo
}

// SetNillableTo sets the "to" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTo(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetTo(*s)
	}
	return tuo
}

// SetAmount sets the "amount" field.
func (tuo *TransactionUpdateOne) SetAmount(u uint64) *TransactionUpdateOne {
	tuo.mutation.ResetAmount()
	tuo.mutation.SetAmount(u)
	return tuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableAmount(u *uint64) *TransactionUpdateOne {
	if u != nil {
		tuo.SetAmount(*u)
	}
	return tuo
}

// AddAmount adds u to the "amount" field.
func (tuo *TransactionUpdateOne) AddAmount(u int64) *TransactionUpdateOne {
	tuo.mutation.AddAmount(u)
	return tuo
}

// SetState sets the "state" field.
func (tuo *TransactionUpdateOne) SetState(u uint8) *TransactionUpdateOne {
	tuo.mutation.ResetState()
	tuo.mutation.SetState(u)
	return tuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableState(u *uint8) *TransactionUpdateOne {
	if u != nil {
		tuo.SetState(*u)
	}
	return tuo
}

// AddState adds u to the "state" field.
func (tuo *TransactionUpdateOne) AddState(u int8) *TransactionUpdateOne {
	tuo.mutation.AddState(u)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TransactionUpdateOne) SetCreatedAt(u uint32) *TransactionUpdateOne {
	tuo.mutation.ResetCreatedAt()
	tuo.mutation.SetCreatedAt(u)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableCreatedAt(u *uint32) *TransactionUpdateOne {
	if u != nil {
		tuo.SetCreatedAt(*u)
	}
	return tuo
}

// AddCreatedAt adds u to the "created_at" field.
func (tuo *TransactionUpdateOne) AddCreatedAt(u int32) *TransactionUpdateOne {
	tuo.mutation.AddCreatedAt(u)
	return tuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TransactionUpdateOne) SetUpdatedAt(u uint32) *TransactionUpdateOne {
	tuo.mutation.ResetUpdatedAt()
	tuo.mutation.SetUpdatedAt(u)
	return tuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tuo *TransactionUpdateOne) AddUpdatedAt(u int32) *TransactionUpdateOne {
	tuo.mutation.AddUpdatedAt(u)
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TransactionUpdateOne) SetDeletedAt(u uint32) *TransactionUpdateOne {
	tuo.mutation.ResetDeletedAt()
	tuo.mutation.SetDeletedAt(u)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableDeletedAt(u *uint32) *TransactionUpdateOne {
	if u != nil {
		tuo.SetDeletedAt(*u)
	}
	return tuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tuo *TransactionUpdateOne) AddDeletedAt(u int32) *TransactionUpdateOne {
	tuo.mutation.AddDeletedAt(u)
	return tuo
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TransactionUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := transaction.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionUpdateOne) check() error {
	if v, ok := tuo.mutation.TransactionID(); ok {
		if err := transaction.TransactionIDValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id", err: fmt.Errorf(`ent: validator failed for field "Transaction.transaction_id": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.From(); ok {
		if err := transaction.FromValidator(v); err != nil {
			return &ValidationError{Name: "from", err: fmt.Errorf(`ent: validator failed for field "Transaction.from": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.To(); ok {
		if err := transaction.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "Transaction.to": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Amount(); ok {
		if err := transaction.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Transaction.amount": %w`, err)}
		}
	}
	return nil
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldNonce,
		})
	}
	if value, ok := tuo.mutation.AddedNonce(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldNonce,
		})
	}
	if value, ok := tuo.mutation.Utxo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transaction.FieldUtxo,
		})
	}
	if value, ok := tuo.mutation.TransactionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: transaction.FieldTransactionType,
		})
	}
	if value, ok := tuo.mutation.AddedTransactionType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: transaction.FieldTransactionType,
		})
	}
	if value, ok := tuo.mutation.CoinType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: transaction.FieldCoinType,
		})
	}
	if value, ok := tuo.mutation.AddedCoinType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: transaction.FieldCoinType,
		})
	}
	if value, ok := tuo.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionID,
		})
	}
	if value, ok := tuo.mutation.Cid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldCid,
		})
	}
	if value, ok := tuo.mutation.ExitCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldExitCode,
		})
	}
	if value, ok := tuo.mutation.AddedExitCode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: transaction.FieldExitCode,
		})
	}
	if value, ok := tuo.mutation.From(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldFrom,
		})
	}
	if value, ok := tuo.mutation.To(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTo,
		})
	}
	if value, ok := tuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmount,
		})
	}
	if value, ok := tuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: transaction.FieldAmount,
		})
	}
	if value, ok := tuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: transaction.FieldState,
		})
	}
	if value, ok := tuo.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: transaction.FieldState,
		})
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldCreatedAt,
		})
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldDeletedAt,
		})
	}
	if value, ok := tuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transaction.FieldDeletedAt,
		})
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
