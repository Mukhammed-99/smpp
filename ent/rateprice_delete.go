// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"smpp/ent/predicate"
	"smpp/ent/rateprice"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RatePriceDelete is the builder for deleting a RatePrice entity.
type RatePriceDelete struct {
	config
	hooks    []Hook
	mutation *RatePriceMutation
}

// Where adds a new predicate to the RatePriceDelete builder.
func (rpd *RatePriceDelete) Where(ps ...predicate.RatePrice) *RatePriceDelete {
	rpd.mutation.predicates = append(rpd.mutation.predicates, ps...)
	return rpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rpd *RatePriceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rpd.hooks) == 0 {
		affected, err = rpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatePriceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rpd.mutation = mutation
			affected, err = rpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rpd.hooks) - 1; i >= 0; i-- {
			mut = rpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpd *RatePriceDelete) ExecX(ctx context.Context) int {
	n, err := rpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rpd *RatePriceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: rateprice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: rateprice.FieldID,
			},
		},
	}
	if ps := rpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rpd.driver, _spec)
}

// RatePriceDeleteOne is the builder for deleting a single RatePrice entity.
type RatePriceDeleteOne struct {
	rpd *RatePriceDelete
}

// Exec executes the deletion query.
func (rpdo *RatePriceDeleteOne) Exec(ctx context.Context) error {
	n, err := rpdo.rpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rateprice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rpdo *RatePriceDeleteOne) ExecX(ctx context.Context) {
	rpdo.rpd.ExecX(ctx)
}