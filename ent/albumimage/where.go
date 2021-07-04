// Code generated by entc, DO NOT EDIT.

package albumimage

import (
	"goreact/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
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
func IDGT(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// ImageMimeType applies equality check predicate on the "image_mime_type" field. It's identical to ImageMimeTypeEQ.
func ImageMimeType(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageMimeType), v))
	})
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImage), v))
	})
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...[]byte) predicate.AlbumImage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AlbumImage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImage), v...))
	})
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...[]byte) predicate.AlbumImage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AlbumImage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImage), v...))
	})
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImage), v))
	})
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImage), v))
	})
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImage), v))
	})
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v []byte) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImage), v))
	})
}

// ImageMimeTypeEQ applies the EQ predicate on the "image_mime_type" field.
func ImageMimeTypeEQ(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeNEQ applies the NEQ predicate on the "image_mime_type" field.
func ImageMimeTypeNEQ(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeIn applies the In predicate on the "image_mime_type" field.
func ImageMimeTypeIn(vs ...string) predicate.AlbumImage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AlbumImage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImageMimeType), v...))
	})
}

// ImageMimeTypeNotIn applies the NotIn predicate on the "image_mime_type" field.
func ImageMimeTypeNotIn(vs ...string) predicate.AlbumImage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AlbumImage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImageMimeType), v...))
	})
}

// ImageMimeTypeGT applies the GT predicate on the "image_mime_type" field.
func ImageMimeTypeGT(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeGTE applies the GTE predicate on the "image_mime_type" field.
func ImageMimeTypeGTE(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeLT applies the LT predicate on the "image_mime_type" field.
func ImageMimeTypeLT(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeLTE applies the LTE predicate on the "image_mime_type" field.
func ImageMimeTypeLTE(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeContains applies the Contains predicate on the "image_mime_type" field.
func ImageMimeTypeContains(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeHasPrefix applies the HasPrefix predicate on the "image_mime_type" field.
func ImageMimeTypeHasPrefix(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeHasSuffix applies the HasSuffix predicate on the "image_mime_type" field.
func ImageMimeTypeHasSuffix(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeEqualFold applies the EqualFold predicate on the "image_mime_type" field.
func ImageMimeTypeEqualFold(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImageMimeType), v))
	})
}

// ImageMimeTypeContainsFold applies the ContainsFold predicate on the "image_mime_type" field.
func ImageMimeTypeContainsFold(v string) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImageMimeType), v))
	})
}

// HasAlbum applies the HasEdge predicate on the "album" edge.
func HasAlbum() predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlbumTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AlbumTable, AlbumColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlbumWith applies the HasEdge predicate on the "album" edge with a given conditions (other predicates).
func HasAlbumWith(preds ...predicate.Album) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AlbumInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AlbumTable, AlbumColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AlbumImage) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AlbumImage) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
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
func Not(p predicate.AlbumImage) predicate.AlbumImage {
	return predicate.AlbumImage(func(s *sql.Selector) {
		p(s.Not())
	})
}