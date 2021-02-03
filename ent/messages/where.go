// Code generated by entc, DO NOT EDIT.

package messages

import (
	"smpp/ent/predicate"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SequenceNumber applies equality check predicate on the "sequence_number" field. It's identical to SequenceNumberEQ.
func SequenceNumber(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequenceNumber), v))
	})
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// Dst applies equality check predicate on the "dst" field. It's identical to DstEQ.
func Dst(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDst), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// Src applies equality check predicate on the "src" field. It's identical to SrcEQ.
func Src(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSrc), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// SmscMessageID applies equality check predicate on the "smsc_message_id" field. It's identical to SmscMessageIDEQ.
func SmscMessageID(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmscMessageID), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// SequenceNumberEQ applies the EQ predicate on the "sequence_number" field.
func SequenceNumberEQ(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequenceNumber), v))
	})
}

// SequenceNumberNEQ applies the NEQ predicate on the "sequence_number" field.
func SequenceNumberNEQ(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSequenceNumber), v))
	})
}

// SequenceNumberIn applies the In predicate on the "sequence_number" field.
func SequenceNumberIn(vs ...int32) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSequenceNumber), v...))
	})
}

// SequenceNumberNotIn applies the NotIn predicate on the "sequence_number" field.
func SequenceNumberNotIn(vs ...int32) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSequenceNumber), v...))
	})
}

// SequenceNumberGT applies the GT predicate on the "sequence_number" field.
func SequenceNumberGT(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSequenceNumber), v))
	})
}

// SequenceNumberGTE applies the GTE predicate on the "sequence_number" field.
func SequenceNumberGTE(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSequenceNumber), v))
	})
}

// SequenceNumberLT applies the LT predicate on the "sequence_number" field.
func SequenceNumberLT(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSequenceNumber), v))
	})
}

// SequenceNumberLTE applies the LTE predicate on the "sequence_number" field.
func SequenceNumberLTE(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSequenceNumber), v))
	})
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExternalID), v))
	})
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExternalID), v...))
	})
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExternalID), v...))
	})
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExternalID), v))
	})
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExternalID), v))
	})
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExternalID), v))
	})
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExternalID), v))
	})
}

// ExternalIDContains applies the Contains predicate on the "external_id" field.
func ExternalIDContains(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExternalID), v))
	})
}

// ExternalIDHasPrefix applies the HasPrefix predicate on the "external_id" field.
func ExternalIDHasPrefix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExternalID), v))
	})
}

// ExternalIDHasSuffix applies the HasSuffix predicate on the "external_id" field.
func ExternalIDHasSuffix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExternalID), v))
	})
}

// ExternalIDIsNil applies the IsNil predicate on the "external_id" field.
func ExternalIDIsNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExternalID)))
	})
}

// ExternalIDNotNil applies the NotNil predicate on the "external_id" field.
func ExternalIDNotNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExternalID)))
	})
}

// ExternalIDEqualFold applies the EqualFold predicate on the "external_id" field.
func ExternalIDEqualFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExternalID), v))
	})
}

// ExternalIDContainsFold applies the ContainsFold predicate on the "external_id" field.
func ExternalIDContainsFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExternalID), v))
	})
}

// DstEQ applies the EQ predicate on the "dst" field.
func DstEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDst), v))
	})
}

// DstNEQ applies the NEQ predicate on the "dst" field.
func DstNEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDst), v))
	})
}

// DstIn applies the In predicate on the "dst" field.
func DstIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDst), v...))
	})
}

// DstNotIn applies the NotIn predicate on the "dst" field.
func DstNotIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDst), v...))
	})
}

// DstGT applies the GT predicate on the "dst" field.
func DstGT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDst), v))
	})
}

// DstGTE applies the GTE predicate on the "dst" field.
func DstGTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDst), v))
	})
}

// DstLT applies the LT predicate on the "dst" field.
func DstLT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDst), v))
	})
}

// DstLTE applies the LTE predicate on the "dst" field.
func DstLTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDst), v))
	})
}

// DstContains applies the Contains predicate on the "dst" field.
func DstContains(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDst), v))
	})
}

// DstHasPrefix applies the HasPrefix predicate on the "dst" field.
func DstHasPrefix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDst), v))
	})
}

// DstHasSuffix applies the HasSuffix predicate on the "dst" field.
func DstHasSuffix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDst), v))
	})
}

// DstIsNil applies the IsNil predicate on the "dst" field.
func DstIsNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDst)))
	})
}

// DstNotNil applies the NotNil predicate on the "dst" field.
func DstNotNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDst)))
	})
}

// DstEqualFold applies the EqualFold predicate on the "dst" field.
func DstEqualFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDst), v))
	})
}

// DstContainsFold applies the ContainsFold predicate on the "dst" field.
func DstContainsFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDst), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// SrcEQ applies the EQ predicate on the "src" field.
func SrcEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSrc), v))
	})
}

// SrcNEQ applies the NEQ predicate on the "src" field.
func SrcNEQ(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSrc), v))
	})
}

// SrcIn applies the In predicate on the "src" field.
func SrcIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSrc), v...))
	})
}

// SrcNotIn applies the NotIn predicate on the "src" field.
func SrcNotIn(vs ...string) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSrc), v...))
	})
}

// SrcGT applies the GT predicate on the "src" field.
func SrcGT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSrc), v))
	})
}

// SrcGTE applies the GTE predicate on the "src" field.
func SrcGTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSrc), v))
	})
}

// SrcLT applies the LT predicate on the "src" field.
func SrcLT(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSrc), v))
	})
}

// SrcLTE applies the LTE predicate on the "src" field.
func SrcLTE(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSrc), v))
	})
}

// SrcContains applies the Contains predicate on the "src" field.
func SrcContains(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSrc), v))
	})
}

// SrcHasPrefix applies the HasPrefix predicate on the "src" field.
func SrcHasPrefix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSrc), v))
	})
}

// SrcHasSuffix applies the HasSuffix predicate on the "src" field.
func SrcHasSuffix(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSrc), v))
	})
}

// SrcEqualFold applies the EqualFold predicate on the "src" field.
func SrcEqualFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSrc), v))
	})
}

// SrcContainsFold applies the ContainsFold predicate on the "src" field.
func SrcContainsFold(v string) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSrc), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int32) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int32) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// SmscMessageIDEQ applies the EQ predicate on the "smsc_message_id" field.
func SmscMessageIDEQ(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSmscMessageID), v))
	})
}

// SmscMessageIDNEQ applies the NEQ predicate on the "smsc_message_id" field.
func SmscMessageIDNEQ(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSmscMessageID), v))
	})
}

// SmscMessageIDIn applies the In predicate on the "smsc_message_id" field.
func SmscMessageIDIn(vs ...int32) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSmscMessageID), v...))
	})
}

// SmscMessageIDNotIn applies the NotIn predicate on the "smsc_message_id" field.
func SmscMessageIDNotIn(vs ...int32) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSmscMessageID), v...))
	})
}

// SmscMessageIDGT applies the GT predicate on the "smsc_message_id" field.
func SmscMessageIDGT(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSmscMessageID), v))
	})
}

// SmscMessageIDGTE applies the GTE predicate on the "smsc_message_id" field.
func SmscMessageIDGTE(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSmscMessageID), v))
	})
}

// SmscMessageIDLT applies the LT predicate on the "smsc_message_id" field.
func SmscMessageIDLT(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSmscMessageID), v))
	})
}

// SmscMessageIDLTE applies the LTE predicate on the "smsc_message_id" field.
func SmscMessageIDLTE(v int32) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSmscMessageID), v))
	})
}

// SmscMessageIDIsNil applies the IsNil predicate on the "smsc_message_id" field.
func SmscMessageIDIsNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSmscMessageID)))
	})
}

// SmscMessageIDNotNil applies the NotNil predicate on the "smsc_message_id" field.
func SmscMessageIDNotNil() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSmscMessageID)))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.Messages {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Messages(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// HasUserID applies the HasEdge predicate on the "user_id" edge.
func HasUserID() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIDWith applies the HasEdge predicate on the "user_id" edge with a given conditions (other predicates).
func HasUserIDWith(preds ...predicate.User) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProviderID applies the HasEdge predicate on the "provider_id" edge.
func HasProviderID() predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderIDTable, ProviderIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderIDWith applies the HasEdge predicate on the "provider_id" edge with a given conditions (other predicates).
func HasProviderIDWith(preds ...predicate.Provide) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderIDTable, ProviderIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Messages) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Messages) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Messages) predicate.Messages {
	return predicate.Messages(func(s *sql.Selector) {
		p(s.Not())
	})
}
