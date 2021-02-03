// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"smpp/ent/messages"
	"smpp/ent/provide"
	"smpp/ent/user"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Messages is the model entity for the Messages schema.
type Messages struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SequenceNumber holds the value of the "sequence_number" field.
	SequenceNumber int32 `json:"sequence_number,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID string `json:"external_id,omitempty"`
	// Dst holds the value of the "dst" field.
	Dst string `json:"dst,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Src holds the value of the "src" field.
	Src string `json:"src,omitempty"`
	// State holds the value of the "state" field.
	State int32 `json:"state,omitempty"`
	// SmscMessageID holds the value of the "smsc_message_id" field.
	SmscMessageID int32 `json:"smsc_message_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessagesQuery when eager-loading is set.
	Edges       MessagesEdges `json:"edges"`
	provider_id *uuid.UUID
	user_id     *uuid.UUID
}

// MessagesEdges holds the relations/edges for other nodes in the graph.
type MessagesEdges struct {
	// UserID holds the value of the user_id edge.
	UserID *User
	// ProviderID holds the value of the provider_id edge.
	ProviderID *Provide
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserIDOrErr returns the UserID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessagesEdges) UserIDOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserID == nil {
			// The edge user_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserID, nil
	}
	return nil, &NotLoadedError{edge: "user_id"}
}

// ProviderIDOrErr returns the ProviderID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessagesEdges) ProviderIDOrErr() (*Provide, error) {
	if e.loadedTypes[1] {
		if e.ProviderID == nil {
			// The edge provider_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: provide.Label}
		}
		return e.ProviderID, nil
	}
	return nil, &NotLoadedError{edge: "provider_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Messages) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case messages.FieldSequenceNumber, messages.FieldState, messages.FieldSmscMessageID:
			values[i] = &sql.NullInt64{}
		case messages.FieldExternalID, messages.FieldDst, messages.FieldMessage, messages.FieldSrc:
			values[i] = &sql.NullString{}
		case messages.FieldCreateAt, messages.FieldUpdateAt:
			values[i] = &sql.NullTime{}
		case messages.FieldID:
			values[i] = &uuid.UUID{}
		case messages.ForeignKeys[0]: // provider_id
			values[i] = &uuid.UUID{}
		case messages.ForeignKeys[1]: // user_id
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Messages", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Messages fields.
func (m *Messages) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messages.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case messages.FieldSequenceNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence_number", values[i])
			} else if value.Valid {
				m.SequenceNumber = int32(value.Int64)
			}
		case messages.FieldExternalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				m.ExternalID = value.String
			}
		case messages.FieldDst:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dst", values[i])
			} else if value.Valid {
				m.Dst = value.String
			}
		case messages.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				m.Message = value.String
			}
		case messages.FieldSrc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field src", values[i])
			} else if value.Valid {
				m.Src = value.String
			}
		case messages.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				m.State = int32(value.Int64)
			}
		case messages.FieldSmscMessageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field smsc_message_id", values[i])
			} else if value.Valid {
				m.SmscMessageID = int32(value.Int64)
			}
		case messages.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				m.CreateAt = value.Time
			}
		case messages.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				m.UpdateAt = value.Time
			}
		case messages.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field provider_id", values[i])
			} else if value != nil {
				m.provider_id = value
			}
		case messages.ForeignKeys[1]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				m.user_id = value
			}
		}
	}
	return nil
}

// QueryUserID queries the "user_id" edge of the Messages entity.
func (m *Messages) QueryUserID() *UserQuery {
	return (&MessagesClient{config: m.config}).QueryUserID(m)
}

// QueryProviderID queries the "provider_id" edge of the Messages entity.
func (m *Messages) QueryProviderID() *ProvideQuery {
	return (&MessagesClient{config: m.config}).QueryProviderID(m)
}

// Update returns a builder for updating this Messages.
// Note that you need to call Messages.Unwrap() before calling this method if this Messages
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Messages) Update() *MessagesUpdateOne {
	return (&MessagesClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the Messages entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Messages) Unwrap() *Messages {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Messages is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Messages) String() string {
	var builder strings.Builder
	builder.WriteString("Messages(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", sequence_number=")
	builder.WriteString(fmt.Sprintf("%v", m.SequenceNumber))
	builder.WriteString(", external_id=")
	builder.WriteString(m.ExternalID)
	builder.WriteString(", dst=")
	builder.WriteString(m.Dst)
	builder.WriteString(", message=")
	builder.WriteString(m.Message)
	builder.WriteString(", src=")
	builder.WriteString(m.Src)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", m.State))
	builder.WriteString(", smsc_message_id=")
	builder.WriteString(fmt.Sprintf("%v", m.SmscMessageID))
	builder.WriteString(", create_at=")
	builder.WriteString(m.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(m.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MessagesSlice is a parsable slice of Messages.
type MessagesSlice []*Messages

func (m MessagesSlice) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
