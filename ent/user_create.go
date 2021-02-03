// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"smpp/ent/messages"
	"smpp/ent/rate"
	"smpp/ent/user"
	"smpp/ent/usermonthmessage"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetBalance sets the "balance" field.
func (uc *UserCreate) SetBalance(i int16) *UserCreate {
	uc.mutation.SetBalance(i)
	return uc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uc *UserCreate) SetNillableBalance(i *int16) *UserCreate {
	if i != nil {
		uc.SetBalance(*i)
	}
	return uc
}

// SetCreateAt sets the "create_at" field.
func (uc *UserCreate) SetCreateAt(t time.Time) *UserCreate {
	uc.mutation.SetCreateAt(t)
	return uc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreateAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreateAt(*t)
	}
	return uc
}

// SetUpdateAt sets the "update_at" field.
func (uc *UserCreate) SetUpdateAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdateAt(t)
	return uc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdateAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdateAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// AddUserMessageIDs adds the "user_messages" edge to the UserMonthMessage entity by IDs.
func (uc *UserCreate) AddUserMessageIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddUserMessageIDs(ids...)
	return uc
}

// AddUserMessages adds the "user_messages" edges to the UserMonthMessage entity.
func (uc *UserCreate) AddUserMessages(u ...*UserMonthMessage) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserMessageIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Messages entity by IDs.
func (uc *UserCreate) AddMessageIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddMessageIDs(ids...)
	return uc
}

// AddMessages adds the "messages" edges to the Messages entity.
func (uc *UserCreate) AddMessages(m ...*Messages) *UserCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddMessageIDs(ids...)
}

// SetRateIDID sets the "rate_id" edge to the Rate entity by ID.
func (uc *UserCreate) SetRateIDID(id uuid.UUID) *UserCreate {
	uc.mutation.SetRateIDID(id)
	return uc
}

// SetNillableRateIDID sets the "rate_id" edge to the Rate entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableRateIDID(id *uuid.UUID) *UserCreate {
	if id != nil {
		uc = uc.SetRateIDID(*id)
	}
	return uc
}

// SetRateID sets the "rate_id" edge to the Rate entity.
func (uc *UserCreate) SetRateID(r *Rate) *UserCreate {
	return uc.SetRateIDID(r.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Balance(); !ok {
		v := user.DefaultBalance
		uc.mutation.SetBalance(v)
	}
	if _, ok := uc.mutation.CreateAt(); !ok {
		v := user.DefaultCreateAt()
		uc.mutation.SetCreateAt(v)
	}
	if _, ok := uc.mutation.UpdateAt(); !ok {
		v := user.DefaultUpdateAt()
		uc.mutation.SetUpdateAt(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New("ent: missing required field \"balance\"")}
	}
	if _, ok := uc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New("ent: missing required field \"create_at\"")}
	}
	if _, ok := uc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New("ent: missing required field \"update_at\"")}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: user.FieldBalance,
		})
		_node.Balance = value
	}
	if value, ok := uc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := uc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if nodes := uc.mutation.UserMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserMessagesTable,
			Columns: []string{user.UserMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: usermonthmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.MessagesTable,
			Columns: []string{user.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: messages.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RateIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RateIDTable,
			Columns: []string{user.RateIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: rate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
