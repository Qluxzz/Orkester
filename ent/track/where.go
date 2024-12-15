// Code generated by ent, DO NOT EDIT.

package track

import (
	"orkester/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldTitle, v))
}

// TrackNumber applies equality check predicate on the "track_number" field. It's identical to TrackNumberEQ.
func TrackNumber(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldTrackNumber, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldPath, v))
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldLength, v))
}

// Mimetype applies equality check predicate on the "mimetype" field. It's identical to MimetypeEQ.
func Mimetype(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldMimetype, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldTitle, v))
}

// TrackNumberEQ applies the EQ predicate on the "track_number" field.
func TrackNumberEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldTrackNumber, v))
}

// TrackNumberNEQ applies the NEQ predicate on the "track_number" field.
func TrackNumberNEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldTrackNumber, v))
}

// TrackNumberIn applies the In predicate on the "track_number" field.
func TrackNumberIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldTrackNumber, vs...))
}

// TrackNumberNotIn applies the NotIn predicate on the "track_number" field.
func TrackNumberNotIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldTrackNumber, vs...))
}

// TrackNumberGT applies the GT predicate on the "track_number" field.
func TrackNumberGT(v int) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldTrackNumber, v))
}

// TrackNumberGTE applies the GTE predicate on the "track_number" field.
func TrackNumberGTE(v int) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldTrackNumber, v))
}

// TrackNumberLT applies the LT predicate on the "track_number" field.
func TrackNumberLT(v int) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldTrackNumber, v))
}

// TrackNumberLTE applies the LTE predicate on the "track_number" field.
func TrackNumberLTE(v int) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldTrackNumber, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldPath, v))
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldLength, v))
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v int) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldLength, v))
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldLength, vs...))
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...int) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldLength, vs...))
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v int) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldLength, v))
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v int) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldLength, v))
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v int) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldLength, v))
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v int) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldLength, v))
}

// MimetypeEQ applies the EQ predicate on the "mimetype" field.
func MimetypeEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldMimetype, v))
}

// MimetypeNEQ applies the NEQ predicate on the "mimetype" field.
func MimetypeNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldMimetype, v))
}

// MimetypeIn applies the In predicate on the "mimetype" field.
func MimetypeIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldMimetype, vs...))
}

// MimetypeNotIn applies the NotIn predicate on the "mimetype" field.
func MimetypeNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldMimetype, vs...))
}

// MimetypeGT applies the GT predicate on the "mimetype" field.
func MimetypeGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldMimetype, v))
}

// MimetypeGTE applies the GTE predicate on the "mimetype" field.
func MimetypeGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldMimetype, v))
}

// MimetypeLT applies the LT predicate on the "mimetype" field.
func MimetypeLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldMimetype, v))
}

// MimetypeLTE applies the LTE predicate on the "mimetype" field.
func MimetypeLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldMimetype, v))
}

// MimetypeContains applies the Contains predicate on the "mimetype" field.
func MimetypeContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldMimetype, v))
}

// MimetypeHasPrefix applies the HasPrefix predicate on the "mimetype" field.
func MimetypeHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldMimetype, v))
}

// MimetypeHasSuffix applies the HasSuffix predicate on the "mimetype" field.
func MimetypeHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldMimetype, v))
}

// MimetypeEqualFold applies the EqualFold predicate on the "mimetype" field.
func MimetypeEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldMimetype, v))
}

// MimetypeContainsFold applies the ContainsFold predicate on the "mimetype" field.
func MimetypeContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldMimetype, v))
}

// HasArtists applies the HasEdge predicate on the "artists" edge.
func HasArtists() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArtistsTable, ArtistsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtistsWith applies the HasEdge predicate on the "artists" edge with a given conditions (other predicates).
func HasArtistsWith(preds ...predicate.Artist) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newArtistsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAlbum applies the HasEdge predicate on the "album" edge.
func HasAlbum() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AlbumTable, AlbumColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlbumWith applies the HasEdge predicate on the "album" edge with a given conditions (other predicates).
func HasAlbumWith(preds ...predicate.Album) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newAlbumStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLiked applies the HasEdge predicate on the "liked" edge.
func HasLiked() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, LikedTable, LikedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikedWith applies the HasEdge predicate on the "liked" edge with a given conditions (other predicates).
func HasLikedWith(preds ...predicate.LikedTrack) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newLikedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImage applies the HasEdge predicate on the "image" edge.
func HasImage() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ImageTable, ImageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImageWith applies the HasEdge predicate on the "image" edge with a given conditions (other predicates).
func HasImageWith(preds ...predicate.Image) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newImageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Track) predicate.Track {
	return predicate.Track(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Track) predicate.Track {
	return predicate.Track(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Track) predicate.Track {
	return predicate.Track(sql.NotPredicates(p))
}
