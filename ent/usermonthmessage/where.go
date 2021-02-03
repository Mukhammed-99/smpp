// Code generated by entc, DO NOT EDIT.

package usermonthmessage

import (
	"smpp/ent/predicate"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Month applies equality check predicate on the "month" field. It's identical to MonthEQ.
func Month(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMonth), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// MonthEQ applies the EQ predicate on the "month" field.
func MonthEQ(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMonth), v))
	})
}

// MonthNEQ applies the NEQ predicate on the "month" field.
func MonthNEQ(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMonth), v))
	})
}

// MonthIn applies the In predicate on the "month" field.
func MonthIn(vs ...time.Time) predicate.UserMonthMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMonth), v...))
	})
}

// MonthNotIn applies the NotIn predicate on the "month" field.
func MonthNotIn(vs ...time.Time) predicate.UserMonthMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMonth), v...))
	})
}

// MonthGT applies the GT predicate on the "month" field.
func MonthGT(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMonth), v))
	})
}

// MonthGTE applies the GTE predicate on the "month" field.
func MonthGTE(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMonth), v))
	})
}

// MonthLT applies the LT predicate on the "month" field.
func MonthLT(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMonth), v))
	})
}

// MonthLTE applies the LTE predicate on the "month" field.
func MonthLTE(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMonth), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.UserMonthMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func CreateAtNotIn(vs ...time.Time) predicate.UserMonthMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func CreateAtGT(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.UserMonthMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func UpdateAtNotIn(vs ...time.Time) predicate.UserMonthMessage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func UpdateAtGT(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// HasProviderID applies the HasEdge predicate on the "provider_id" edge.
func HasProviderID() predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderIDTable, ProviderIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderIDWith applies the HasEdge predicate on the "provider_id" edge with a given conditions (other predicates).
func HasProviderIDWith(preds ...predicate.Provide) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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

// HasUserID applies the HasEdge predicate on the "user_id" edge.
func HasUserID() predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIDWith applies the HasEdge predicate on the "user_id" edge with a given conditions (other predicates).
func HasUserIDWith(preds ...predicate.User) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserMonthMessage) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserMonthMessage) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
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
func Not(p predicate.UserMonthMessage) predicate.UserMonthMessage {
	return predicate.UserMonthMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}
