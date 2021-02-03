// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"smpp/ent/provide"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Provide is the model entity for the Provide schema.
type Provide struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IPAdres holds the value of the "ip_adres" field.
	IPAdres string `json:"ip_adres,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvideQuery when eager-loading is set.
	Edges ProvideEdges `json:"edges"`
}

// ProvideEdges holds the relations/edges for other nodes in the graph.
type ProvideEdges struct {
	// ProviderID holds the value of the provider_id edge.
	ProviderID []*UserMonthMessage
	// Messages holds the value of the messages edge.
	Messages []*Messages
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProviderIDOrErr returns the ProviderID value or an error if the edge
// was not loaded in eager-loading.
func (e ProvideEdges) ProviderIDOrErr() ([]*UserMonthMessage, error) {
	if e.loadedTypes[0] {
		return e.ProviderID, nil
	}
	return nil, &NotLoadedError{edge: "provider_id"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e ProvideEdges) MessagesOrErr() ([]*Messages, error) {
	if e.loadedTypes[1] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provide) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provide.FieldName, provide.FieldIPAdres:
			values[i] = &sql.NullString{}
		case provide.FieldCreateAt, provide.FieldUpdateAt:
			values[i] = &sql.NullTime{}
		case provide.FieldID:
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Provide", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provide fields.
func (pr *Provide) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provide.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case provide.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case provide.FieldIPAdres:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_adres", values[i])
			} else if value.Valid {
				pr.IPAdres = value.String
			}
		case provide.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				pr.CreateAt = value.Time
			}
		case provide.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				pr.UpdateAt = value.Time
			}
		}
	}
	return nil
}

// QueryProviderID queries the "provider_id" edge of the Provide entity.
func (pr *Provide) QueryProviderID() *UserMonthMessageQuery {
	return (&ProvideClient{config: pr.config}).QueryProviderID(pr)
}

// QueryMessages queries the "messages" edge of the Provide entity.
func (pr *Provide) QueryMessages() *MessagesQuery {
	return (&ProvideClient{config: pr.config}).QueryMessages(pr)
}

// Update returns a builder for updating this Provide.
// Note that you need to call Provide.Unwrap() before calling this method if this Provide
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provide) Update() *ProvideUpdateOne {
	return (&ProvideClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Provide entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provide) Unwrap() *Provide {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Provide is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provide) String() string {
	var builder strings.Builder
	builder.WriteString("Provide(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ip_adres=")
	builder.WriteString(pr.IPAdres)
	builder.WriteString(", create_at=")
	builder.WriteString(pr.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(pr.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Provides is a parsable slice of Provide.
type Provides []*Provide

func (pr Provides) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}