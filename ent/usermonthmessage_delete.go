// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"smpp/ent/predicate"
	"smpp/ent/usermonthmessage"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// UserMonthMessageDelete is the builder for deleting a UserMonthMessage entity.
type UserMonthMessageDelete struct {
	config
	hooks    []Hook
	mutation *UserMonthMessageMutation
}

// Where adds a new predicate to the UserMonthMessageDelete builder.
func (ummd *UserMonthMessageDelete) Where(ps ...predicate.UserMonthMessage) *UserMonthMessageDelete {
	ummd.mutation.predicates = append(ummd.mutation.predicates, ps...)
	return ummd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ummd *UserMonthMessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ummd.hooks) == 0 {
		affected, err = ummd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMonthMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ummd.mutation = mutation
			affected, err = ummd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ummd.hooks) - 1; i >= 0; i-- {
			mut = ummd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ummd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ummd *UserMonthMessageDelete) ExecX(ctx context.Context) int {
	n, err := ummd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ummd *UserMonthMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usermonthmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: usermonthmessage.FieldID,
			},
		},
	}
	if ps := ummd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ummd.driver, _spec)
}

// UserMonthMessageDeleteOne is the builder for deleting a single UserMonthMessage entity.
type UserMonthMessageDeleteOne struct {
	ummd *UserMonthMessageDelete
}

// Exec executes the deletion query.
func (ummdo *UserMonthMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := ummdo.ummd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usermonthmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ummdo *UserMonthMessageDeleteOne) ExecX(ctx context.Context) {
	ummdo.ummd.ExecX(ctx)
}
