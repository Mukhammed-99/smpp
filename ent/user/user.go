// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBalance holds the string denoting the balance field in the database.
	FieldBalance = "balance"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"

	// EdgeUserMessages holds the string denoting the user_messages edge name in mutations.
	EdgeUserMessages = "user_messages"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeRateID holds the string denoting the rate_id edge name in mutations.
	EdgeRateID = "rate_id"

	// Table holds the table name of the user in the database.
	Table = "users"
	// UserMessagesTable is the table the holds the user_messages relation/edge.
	UserMessagesTable = "user_month_messages"
	// UserMessagesInverseTable is the table name for the UserMonthMessage entity.
	// It exists in this package in order to avoid circular dependency with the "usermonthmessage" package.
	UserMessagesInverseTable = "user_month_messages"
	// UserMessagesColumn is the table column denoting the user_messages relation/edge.
	UserMessagesColumn = "user_id"
	// MessagesTable is the table the holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Messages entity.
	// It exists in this package in order to avoid circular dependency with the "messages" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "user_id"
	// RateIDTable is the table the holds the rate_id relation/edge.
	RateIDTable = "users"
	// RateIDInverseTable is the table name for the Rate entity.
	// It exists in this package in order to avoid circular dependency with the "rate" package.
	RateIDInverseTable = "rates"
	// RateIDColumn is the table column denoting the rate_id relation/edge.
	RateIDColumn = "rate_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldBalance,
	FieldCreateAt,
	FieldUpdateAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"rate_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultBalance holds the default value on creation for the "balance" field.
	DefaultBalance int16
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)