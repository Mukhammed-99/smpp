// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"smpp/ent/messages"
	"smpp/ent/predicate"
	"smpp/ent/rate"
	"smpp/ent/user"
	"smpp/ent/usermonthmessage"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetBalance sets the "balance" field.
func (uu *UserUpdate) SetBalance(i int16) *UserUpdate {
	uu.mutation.ResetBalance()
	uu.mutation.SetBalance(i)
	return uu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBalance(i *int16) *UserUpdate {
	if i != nil {
		uu.SetBalance(*i)
	}
	return uu
}

// AddBalance adds i to the "balance" field.
func (uu *UserUpdate) AddBalance(i int16) *UserUpdate {
	uu.mutation.AddBalance(i)
	return uu
}

// SetCreateAt sets the "create_at" field.
func (uu *UserUpdate) SetCreateAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreateAt(t)
	return uu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreateAt(*t)
	}
	return uu
}

// SetUpdateAt sets the "update_at" field.
func (uu *UserUpdate) SetUpdateAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateAt(t)
	return uu
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdateAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdateAt(*t)
	}
	return uu
}

// AddUserMessageIDs adds the "user_messages" edge to the UserMonthMessage entity by IDs.
func (uu *UserUpdate) AddUserMessageIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddUserMessageIDs(ids...)
	return uu
}

// AddUserMessages adds the "user_messages" edges to the UserMonthMessage entity.
func (uu *UserUpdate) AddUserMessages(u ...*UserMonthMessage) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserMessageIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Messages entity by IDs.
func (uu *UserUpdate) AddMessageIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddMessageIDs(ids...)
	return uu
}

// AddMessages adds the "messages" edges to the Messages entity.
func (uu *UserUpdate) AddMessages(m ...*Messages) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddMessageIDs(ids...)
}

// SetRateIDID sets the "rate_id" edge to the Rate entity by ID.
func (uu *UserUpdate) SetRateIDID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetRateIDID(id)
	return uu
}

// SetNillableRateIDID sets the "rate_id" edge to the Rate entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableRateIDID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetRateIDID(*id)
	}
	return uu
}

// SetRateID sets the "rate_id" edge to the Rate entity.
func (uu *UserUpdate) SetRateID(r *Rate) *UserUpdate {
	return uu.SetRateIDID(r.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUserMessages clears all "user_messages" edges to the UserMonthMessage entity.
func (uu *UserUpdate) ClearUserMessages() *UserUpdate {
	uu.mutation.ClearUserMessages()
	return uu
}

// RemoveUserMessageIDs removes the "user_messages" edge to UserMonthMessage entities by IDs.
func (uu *UserUpdate) RemoveUserMessageIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveUserMessageIDs(ids...)
	return uu
}

// RemoveUserMessages removes "user_messages" edges to UserMonthMessage entities.
func (uu *UserUpdate) RemoveUserMessages(u ...*UserMonthMessage) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserMessageIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Messages entity.
func (uu *UserUpdate) ClearMessages() *UserUpdate {
	uu.mutation.ClearMessages()
	return uu
}

// RemoveMessageIDs removes the "messages" edge to Messages entities by IDs.
func (uu *UserUpdate) RemoveMessageIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveMessageIDs(ids...)
	return uu
}

// RemoveMessages removes "messages" edges to Messages entities.
func (uu *UserUpdate) RemoveMessages(m ...*Messages) *UserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveMessageIDs(ids...)
}

// ClearRateID clears the "rate_id" edge to the Rate entity.
func (uu *UserUpdate) ClearRateID() *UserUpdate {
	uu.mutation.ClearRateID()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: user.FieldBalance,
		})
	}
	if value, ok := uu.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: user.FieldBalance,
		})
	}
	if value, ok := uu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateAt,
		})
	}
	if value, ok := uu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateAt,
		})
	}
	if uu.mutation.UserMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserMessagesIDs(); len(nodes) > 0 && !uu.mutation.UserMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserMessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !uu.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RateIDCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RateIDIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetBalance sets the "balance" field.
func (uuo *UserUpdateOne) SetBalance(i int16) *UserUpdateOne {
	uuo.mutation.ResetBalance()
	uuo.mutation.SetBalance(i)
	return uuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBalance(i *int16) *UserUpdateOne {
	if i != nil {
		uuo.SetBalance(*i)
	}
	return uuo
}

// AddBalance adds i to the "balance" field.
func (uuo *UserUpdateOne) AddBalance(i int16) *UserUpdateOne {
	uuo.mutation.AddBalance(i)
	return uuo
}

// SetCreateAt sets the "create_at" field.
func (uuo *UserUpdateOne) SetCreateAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreateAt(t)
	return uuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreateAt(*t)
	}
	return uuo
}

// SetUpdateAt sets the "update_at" field.
func (uuo *UserUpdateOne) SetUpdateAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateAt(t)
	return uuo
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdateAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdateAt(*t)
	}
	return uuo
}

// AddUserMessageIDs adds the "user_messages" edge to the UserMonthMessage entity by IDs.
func (uuo *UserUpdateOne) AddUserMessageIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddUserMessageIDs(ids...)
	return uuo
}

// AddUserMessages adds the "user_messages" edges to the UserMonthMessage entity.
func (uuo *UserUpdateOne) AddUserMessages(u ...*UserMonthMessage) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserMessageIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Messages entity by IDs.
func (uuo *UserUpdateOne) AddMessageIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddMessageIDs(ids...)
	return uuo
}

// AddMessages adds the "messages" edges to the Messages entity.
func (uuo *UserUpdateOne) AddMessages(m ...*Messages) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddMessageIDs(ids...)
}

// SetRateIDID sets the "rate_id" edge to the Rate entity by ID.
func (uuo *UserUpdateOne) SetRateIDID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetRateIDID(id)
	return uuo
}

// SetNillableRateIDID sets the "rate_id" edge to the Rate entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRateIDID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetRateIDID(*id)
	}
	return uuo
}

// SetRateID sets the "rate_id" edge to the Rate entity.
func (uuo *UserUpdateOne) SetRateID(r *Rate) *UserUpdateOne {
	return uuo.SetRateIDID(r.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUserMessages clears all "user_messages" edges to the UserMonthMessage entity.
func (uuo *UserUpdateOne) ClearUserMessages() *UserUpdateOne {
	uuo.mutation.ClearUserMessages()
	return uuo
}

// RemoveUserMessageIDs removes the "user_messages" edge to UserMonthMessage entities by IDs.
func (uuo *UserUpdateOne) RemoveUserMessageIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveUserMessageIDs(ids...)
	return uuo
}

// RemoveUserMessages removes "user_messages" edges to UserMonthMessage entities.
func (uuo *UserUpdateOne) RemoveUserMessages(u ...*UserMonthMessage) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserMessageIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Messages entity.
func (uuo *UserUpdateOne) ClearMessages() *UserUpdateOne {
	uuo.mutation.ClearMessages()
	return uuo
}

// RemoveMessageIDs removes the "messages" edge to Messages entities by IDs.
func (uuo *UserUpdateOne) RemoveMessageIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveMessageIDs(ids...)
	return uuo
}

// RemoveMessages removes "messages" edges to Messages entities.
func (uuo *UserUpdateOne) RemoveMessages(m ...*Messages) *UserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveMessageIDs(ids...)
}

// ClearRateID clears the "rate_id" edge to the Rate entity.
func (uuo *UserUpdateOne) ClearRateID() *UserUpdateOne {
	uuo.mutation.ClearRateID()
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: user.FieldBalance,
		})
	}
	if value, ok := uuo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: user.FieldBalance,
		})
	}
	if value, ok := uuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreateAt,
		})
	}
	if value, ok := uuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdateAt,
		})
	}
	if uuo.mutation.UserMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserMessagesIDs(); len(nodes) > 0 && !uuo.mutation.UserMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserMessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !uuo.mutation.MessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RateIDCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RateIDIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
