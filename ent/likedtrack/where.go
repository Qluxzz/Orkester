// Code generated by ent, DO NOT EDIT.

package likedtrack

import (
	"orkester/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DateAdded applies equality check predicate on the "date_added" field. It's identical to DateAddedEQ.
func DateAdded(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateAdded), v))
	})
}

// DateAddedEQ applies the EQ predicate on the "date_added" field.
func DateAddedEQ(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateAdded), v))
	})
}

// DateAddedNEQ applies the NEQ predicate on the "date_added" field.
func DateAddedNEQ(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateAdded), v))
	})
}

// DateAddedIn applies the In predicate on the "date_added" field.
func DateAddedIn(vs ...time.Time) predicate.LikedTrack {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDateAdded), v...))
	})
}

// DateAddedNotIn applies the NotIn predicate on the "date_added" field.
func DateAddedNotIn(vs ...time.Time) predicate.LikedTrack {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDateAdded), v...))
	})
}

// DateAddedGT applies the GT predicate on the "date_added" field.
func DateAddedGT(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateAdded), v))
	})
}

// DateAddedGTE applies the GTE predicate on the "date_added" field.
func DateAddedGTE(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateAdded), v))
	})
}

// DateAddedLT applies the LT predicate on the "date_added" field.
func DateAddedLT(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateAdded), v))
	})
}

// DateAddedLTE applies the LTE predicate on the "date_added" field.
func DateAddedLTE(v time.Time) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateAdded), v))
	})
}

// HasTrack applies the HasEdge predicate on the "track" edge.
func HasTrack() predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TrackTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TrackTable, TrackColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTrackWith applies the HasEdge predicate on the "track" edge with a given conditions (other predicates).
func HasTrackWith(preds ...predicate.Track) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TrackInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TrackTable, TrackColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LikedTrack) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LikedTrack) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
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
func Not(p predicate.LikedTrack) predicate.LikedTrack {
	return predicate.LikedTrack(func(s *sql.Selector) {
		p(s.Not())
	})
}
