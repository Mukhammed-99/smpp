// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"smpp/ent/rate"
	"smpp/ent/user"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance int16 `json:"balance,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges   UserEdges `json:"edges"`
	rate_id *uuid.UUID
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserMessages holds the value of the user_messages edge.
	UserMessages []*UserMonthMessage
	// Messages holds the value of the messages edge.
	Messages []*Messages
	// RateID holds the value of the rate_id edge.
	RateID *Rate
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserMessagesOrErr returns the UserMessages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserMessagesOrErr() ([]*UserMonthMessage, error) {
	if e.loadedTypes[0] {
		return e.UserMessages, nil
	}
	return nil, &NotLoadedError{edge: "user_messages"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessagesOrErr() ([]*Messages, error) {
	if e.loadedTypes[1] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// RateIDOrErr returns the RateID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RateIDOrErr() (*Rate, error) {
	if e.loadedTypes[2] {
		if e.RateID == nil {
			// The edge rate_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: rate.Label}
		}
		return e.RateID, nil
	}
	return nil, &NotLoadedError{edge: "rate_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldBalance:
			values[i] = &sql.NullInt64{}
		case user.FieldCreateAt, user.FieldUpdateAt:
			values[i] = &sql.NullTime{}
		case user.FieldID:
			values[i] = &uuid.UUID{}
		case user.ForeignKeys[0]: // rate_id
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				u.Balance = int16(value.Int64)
			}
		case user.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				u.CreateAt = value.Time
			}
		case user.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				u.UpdateAt = value.Time
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field rate_id", values[i])
			} else if value != nil {
				u.rate_id = value
			}
		}
	}
	return nil
}

// QueryUserMessages queries the "user_messages" edge of the User entity.
func (u *User) QueryUserMessages() *UserMonthMessageQuery {
	return (&UserClient{config: u.config}).QueryUserMessages(u)
}

// QueryMessages queries the "messages" edge of the User entity.
func (u *User) QueryMessages() *MessagesQuery {
	return (&UserClient{config: u.config}).QueryMessages(u)
}

// QueryRateID queries the "rate_id" edge of the User entity.
func (u *User) QueryRateID() *RateQuery {
	return (&UserClient{config: u.config}).QueryRateID(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", balance=")
	builder.WriteString(fmt.Sprintf("%v", u.Balance))
	builder.WriteString(", create_at=")
	builder.WriteString(u.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(u.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
