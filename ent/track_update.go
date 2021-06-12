// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/predicate"
	"goreact/ent/track"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrackUpdate is the builder for updating Track entities.
type TrackUpdate struct {
	config
	hooks    []Hook
	mutation *TrackMutation
}

// Where adds a new predicate for the TrackUpdate builder.
func (tu *TrackUpdate) Where(ps ...predicate.Track) *TrackUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (tu *TrackUpdate) AddArtistIDs(ids ...int) *TrackUpdate {
	tu.mutation.AddArtistIDs(ids...)
	return tu
}

// AddArtists adds the "artists" edges to the Artist entity.
func (tu *TrackUpdate) AddArtists(a ...*Artist) *TrackUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.AddArtistIDs(ids...)
}

// SetAlbumID sets the "album" edge to the Album entity by ID.
func (tu *TrackUpdate) SetAlbumID(id int) *TrackUpdate {
	tu.mutation.SetAlbumID(id)
	return tu
}

// SetNillableAlbumID sets the "album" edge to the Album entity by ID if the given value is not nil.
func (tu *TrackUpdate) SetNillableAlbumID(id *int) *TrackUpdate {
	if id != nil {
		tu = tu.SetAlbumID(*id)
	}
	return tu
}

// SetAlbum sets the "album" edge to the Album entity.
func (tu *TrackUpdate) SetAlbum(a *Album) *TrackUpdate {
	return tu.SetAlbumID(a.ID)
}

// Mutation returns the TrackMutation object of the builder.
func (tu *TrackUpdate) Mutation() *TrackMutation {
	return tu.mutation
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (tu *TrackUpdate) ClearArtists() *TrackUpdate {
	tu.mutation.ClearArtists()
	return tu
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (tu *TrackUpdate) RemoveArtistIDs(ids ...int) *TrackUpdate {
	tu.mutation.RemoveArtistIDs(ids...)
	return tu
}

// RemoveArtists removes "artists" edges to Artist entities.
func (tu *TrackUpdate) RemoveArtists(a ...*Artist) *TrackUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.RemoveArtistIDs(ids...)
}

// ClearAlbum clears the "album" edge to the Album entity.
func (tu *TrackUpdate) ClearAlbum() *TrackUpdate {
	tu.mutation.ClearAlbum()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TrackUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TrackUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TrackUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TrackUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TrackUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   track.Table,
			Columns: track.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: track.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !tu.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AlbumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{track.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TrackUpdateOne is the builder for updating a single Track entity.
type TrackUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrackMutation
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (tuo *TrackUpdateOne) AddArtistIDs(ids ...int) *TrackUpdateOne {
	tuo.mutation.AddArtistIDs(ids...)
	return tuo
}

// AddArtists adds the "artists" edges to the Artist entity.
func (tuo *TrackUpdateOne) AddArtists(a ...*Artist) *TrackUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.AddArtistIDs(ids...)
}

// SetAlbumID sets the "album" edge to the Album entity by ID.
func (tuo *TrackUpdateOne) SetAlbumID(id int) *TrackUpdateOne {
	tuo.mutation.SetAlbumID(id)
	return tuo
}

// SetNillableAlbumID sets the "album" edge to the Album entity by ID if the given value is not nil.
func (tuo *TrackUpdateOne) SetNillableAlbumID(id *int) *TrackUpdateOne {
	if id != nil {
		tuo = tuo.SetAlbumID(*id)
	}
	return tuo
}

// SetAlbum sets the "album" edge to the Album entity.
func (tuo *TrackUpdateOne) SetAlbum(a *Album) *TrackUpdateOne {
	return tuo.SetAlbumID(a.ID)
}

// Mutation returns the TrackMutation object of the builder.
func (tuo *TrackUpdateOne) Mutation() *TrackMutation {
	return tuo.mutation
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (tuo *TrackUpdateOne) ClearArtists() *TrackUpdateOne {
	tuo.mutation.ClearArtists()
	return tuo
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (tuo *TrackUpdateOne) RemoveArtistIDs(ids ...int) *TrackUpdateOne {
	tuo.mutation.RemoveArtistIDs(ids...)
	return tuo
}

// RemoveArtists removes "artists" edges to Artist entities.
func (tuo *TrackUpdateOne) RemoveArtists(a ...*Artist) *TrackUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.RemoveArtistIDs(ids...)
}

// ClearAlbum clears the "album" edge to the Album entity.
func (tuo *TrackUpdateOne) ClearAlbum() *TrackUpdateOne {
	tuo.mutation.ClearAlbum()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TrackUpdateOne) Select(field string, fields ...string) *TrackUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Track entity.
func (tuo *TrackUpdateOne) Save(ctx context.Context) (*Track, error) {
	var (
		err  error
		node *Track
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TrackUpdateOne) SaveX(ctx context.Context) *Track {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TrackUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TrackUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TrackUpdateOne) sqlSave(ctx context.Context) (_node *Track, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   track.Table,
			Columns: track.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: track.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Track.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, track.FieldID)
		for _, f := range fields {
			if !track.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != track.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !tuo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   track.ArtistsTable,
			Columns: track.ArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.AlbumCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AlbumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   track.AlbumTable,
			Columns: []string{track.AlbumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Track{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{track.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}