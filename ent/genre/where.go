// Code generated by entc, DO NOT EDIT.

package genre

import (
	"goreact/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// URLName applies equality check predicate on the "url_name" field. It's identical to URLNameEQ.
func URLName(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLName), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Genre {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Genre(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Genre {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Genre(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// URLNameEQ applies the EQ predicate on the "url_name" field.
func URLNameEQ(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLName), v))
	})
}

// URLNameNEQ applies the NEQ predicate on the "url_name" field.
func URLNameNEQ(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURLName), v))
	})
}

// URLNameIn applies the In predicate on the "url_name" field.
func URLNameIn(vs ...string) predicate.Genre {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Genre(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURLName), v...))
	})
}

// URLNameNotIn applies the NotIn predicate on the "url_name" field.
func URLNameNotIn(vs ...string) predicate.Genre {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Genre(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURLName), v...))
	})
}

// URLNameGT applies the GT predicate on the "url_name" field.
func URLNameGT(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURLName), v))
	})
}

// URLNameGTE applies the GTE predicate on the "url_name" field.
func URLNameGTE(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURLName), v))
	})
}

// URLNameLT applies the LT predicate on the "url_name" field.
func URLNameLT(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURLName), v))
	})
}

// URLNameLTE applies the LTE predicate on the "url_name" field.
func URLNameLTE(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURLName), v))
	})
}

// URLNameContains applies the Contains predicate on the "url_name" field.
func URLNameContains(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURLName), v))
	})
}

// URLNameHasPrefix applies the HasPrefix predicate on the "url_name" field.
func URLNameHasPrefix(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURLName), v))
	})
}

// URLNameHasSuffix applies the HasSuffix predicate on the "url_name" field.
func URLNameHasSuffix(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURLName), v))
	})
}

// URLNameEqualFold applies the EqualFold predicate on the "url_name" field.
func URLNameEqualFold(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURLName), v))
	})
}

// URLNameContainsFold applies the ContainsFold predicate on the "url_name" field.
func URLNameContainsFold(v string) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURLName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Genre) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Genre) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
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
func Not(p predicate.Genre) predicate.Genre {
	return predicate.Genre(func(s *sql.Selector) {
		p(s.Not())
	})
}
