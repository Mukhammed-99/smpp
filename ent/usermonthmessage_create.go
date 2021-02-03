// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"smpp/ent/provide"
	"smpp/ent/user"
	"smpp/ent/usermonthmessage"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserMonthMessageCreate is the builder for creating a UserMonthMessage entity.
type UserMonthMessageCreate struct {
	config
	mutation *UserMonthMessageMutation
	hooks    []Hook
}

// SetMonth sets the "month" field.
func (ummc *UserMonthMessageCreate) SetMonth(t time.Time) *UserMonthMessageCreate {
	ummc.mutation.SetMonth(t)
	return ummc
}

// SetNillableMonth sets the "month" field if the given value is not nil.
func (ummc *UserMonthMessageCreate) SetNillableMonth(t *time.Time) *UserMonthMessageCreate {
	if t != nil {
		ummc.SetMonth(*t)
	}
	return ummc
}

// SetCreateAt sets the "create_at" field.
func (ummc *UserMonthMessageCreate) SetCreateAt(t time.Time) *UserMonthMessageCreate {
	ummc.mutation.SetCreateAt(t)
	return ummc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ummc *UserMonthMessageCreate) SetNillableCreateAt(t *time.Time) *UserMonthMessageCreate {
	if t != nil {
		ummc.SetCreateAt(*t)
	}
	return ummc
}

// SetUpdateAt sets the "update_at" field.
func (ummc *UserMonthMessageCreate) SetUpdateAt(t time.Time) *UserMonthMessageCreate {
	ummc.mutation.SetUpdateAt(t)
	return ummc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (ummc *UserMonthMessageCreate) SetNillableUpdateAt(t *time.Time) *UserMonthMessageCreate {
	if t != nil {
		ummc.SetUpdateAt(*t)
	}
	return ummc
}

// SetID sets the "id" field.
func (ummc *UserMonthMessageCreate) SetID(u uuid.UUID) *UserMonthMessageCreate {
	ummc.mutation.SetID(u)
	return ummc
}

// SetProviderIDID sets the "provider_id" edge to the Provide entity by ID.
func (ummc *UserMonthMessageCreate) SetProviderIDID(id uuid.UUID) *UserMonthMessageCreate {
	ummc.mutation.SetProviderIDID(id)
	return ummc
}

// SetNillableProviderIDID sets the "provider_id" edge to the Provide entity by ID if the given value is not nil.
func (ummc *UserMonthMessageCreate) SetNillableProviderIDID(id *uuid.UUID) *UserMonthMessageCreate {
	if id != nil {
		ummc = ummc.SetProviderIDID(*id)
	}
	return ummc
}

// SetProviderID sets the "provider_id" edge to the Provide entity.
func (ummc *UserMonthMessageCreate) SetProviderID(p *Provide) *UserMonthMessageCreate {
	return ummc.SetProviderIDID(p.ID)
}

// SetUserIDID sets the "user_id" edge to the User entity by ID.
func (ummc *UserMonthMessageCreate) SetUserIDID(id uuid.UUID) *UserMonthMessageCreate {
	ummc.mutation.SetUserIDID(id)
	return ummc
}

// SetNillableUserIDID sets the "user_id" edge to the User entity by ID if the given value is not nil.
func (ummc *UserMonthMessageCreate) SetNillableUserIDID(id *uuid.UUID) *UserMonthMessageCreate {
	if id != nil {
		ummc = ummc.SetUserIDID(*id)
	}
	return ummc
}

// SetUserID sets the "user_id" edge to the User entity.
func (ummc *UserMonthMessageCreate) SetUserID(u *User) *UserMonthMessageCreate {
	return ummc.SetUserIDID(u.ID)
}

// Mutation returns the UserMonthMessageMutation object of the builder.
func (ummc *UserMonthMessageCreate) Mutation() *UserMonthMessageMutation {
	return ummc.mutation
}

// Save creates the UserMonthMessage in the database.
func (ummc *UserMonthMessageCreate) Save(ctx context.Context) (*UserMonthMessage, error) {
	var (
		err  error
		node *UserMonthMessage
	)
	ummc.defaults()
	if len(ummc.hooks) == 0 {
		if err = ummc.check(); err != nil {
			return nil, err
		}
		node, err = ummc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMonthMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ummc.check(); err != nil {
				return nil, err
			}
			ummc.mutation = mutation
			node, err = ummc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ummc.hooks) - 1; i >= 0; i-- {
			mut = ummc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ummc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ummc *UserMonthMessageCreate) SaveX(ctx context.Context) *UserMonthMessage {
	v, err := ummc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ummc *UserMonthMessageCreate) defaults() {
	if _, ok := ummc.mutation.Month(); !ok {
		v := usermonthmessage.DefaultMonth()
		ummc.mutation.SetMonth(v)
	}
	if _, ok := ummc.mutation.CreateAt(); !ok {
		v := usermonthmessage.DefaultCreateAt()
		ummc.mutation.SetCreateAt(v)
	}
	if _, ok := ummc.mutation.UpdateAt(); !ok {
		v := usermonthmessage.DefaultUpdateAt()
		ummc.mutation.SetUpdateAt(v)
	}
	if _, ok := ummc.mutation.ID(); !ok {
		v := usermonthmessage.DefaultID()
		ummc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ummc *UserMonthMessageCreate) check() error {
	if _, ok := ummc.mutation.Month(); !ok {
		return &ValidationError{Name: "month", err: errors.New("ent: missing required field \"month\"")}
	}
	if _, ok := ummc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New("ent: missing required field \"create_at\"")}
	}
	if _, ok := ummc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New("ent: missing required field \"update_at\"")}
	}
	return nil
}

func (ummc *UserMonthMessageCreate) sqlSave(ctx context.Context) (*UserMonthMessage, error) {
	_node, _spec := ummc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ummc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (ummc *UserMonthMessageCreate) createSpec() (*UserMonthMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &UserMonthMessage{config: ummc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usermonthmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: usermonthmessage.FieldID,
			},
		}
	)
	if id, ok := ummc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ummc.mutation.Month(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usermonthmessage.FieldMonth,
		})
		_node.Month = value
	}
	if value, ok := ummc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usermonthmessage.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := ummc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usermonthmessage.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if nodes := ummc.mutation.ProviderIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermonthmessage.ProviderIDTable,
			Columns: []string{usermonthmessage.ProviderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provide.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ummc.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usermonthmessage.UserIDTable,
			Columns: []string{usermonthmessage.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
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

// UserMonthMessageCreateBulk is the builder for creating many UserMonthMessage entities in bulk.
type UserMonthMessageCreateBulk struct {
	config
	builders []*UserMonthMessageCreate
}

// Save creates the UserMonthMessage entities in the database.
func (ummcb *UserMonthMessageCreateBulk) Save(ctx context.Context) ([]*UserMonthMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ummcb.builders))
	nodes := make([]*UserMonthMessage, len(ummcb.builders))
	mutators := make([]Mutator, len(ummcb.builders))
	for i := range ummcb.builders {
		func(i int, root context.Context) {
			builder := ummcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMonthMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, ummcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ummcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ummcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ummcb *UserMonthMessageCreateBulk) SaveX(ctx context.Context) []*UserMonthMessage {
	v, err := ummcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}