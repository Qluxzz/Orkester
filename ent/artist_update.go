// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"orkester/ent/album"
	"orkester/ent/artist"
	"orkester/ent/predicate"
	"orkester/ent/track"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArtistUpdate is the builder for updating Artist entities.
type ArtistUpdate struct {
	config
	hooks    []Hook
	mutation *ArtistMutation
}

// Where appends a list predicates to the ArtistUpdate builder.
func (au *ArtistUpdate) Where(ps ...predicate.Artist) *ArtistUpdate {
	au.mutation.Where(ps...)
	return au
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (au *ArtistUpdate) AddAlbumIDs(ids ...int) *ArtistUpdate {
	au.mutation.AddAlbumIDs(ids...)
	return au
}

// AddAlbums adds the "albums" edges to the Album entity.
func (au *ArtistUpdate) AddAlbums(a ...*Album) *ArtistUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAlbumIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (au *ArtistUpdate) AddTrackIDs(ids ...int) *ArtistUpdate {
	au.mutation.AddTrackIDs(ids...)
	return au
}

// AddTracks adds the "tracks" edges to the Track entity.
func (au *ArtistUpdate) AddTracks(t ...*Track) *ArtistUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTrackIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (au *ArtistUpdate) Mutation() *ArtistMutation {
	return au.mutation
}

// ClearAlbums clears all "albums" edges to the Album entity.
func (au *ArtistUpdate) ClearAlbums() *ArtistUpdate {
	au.mutation.ClearAlbums()
	return au
}

// RemoveAlbumIDs removes the "albums" edge to Album entities by IDs.
func (au *ArtistUpdate) RemoveAlbumIDs(ids ...int) *ArtistUpdate {
	au.mutation.RemoveAlbumIDs(ids...)
	return au
}

// RemoveAlbums removes "albums" edges to Album entities.
func (au *ArtistUpdate) RemoveAlbums(a ...*Album) *ArtistUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAlbumIDs(ids...)
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (au *ArtistUpdate) ClearTracks() *ArtistUpdate {
	au.mutation.ClearTracks()
	return au
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (au *ArtistUpdate) RemoveTrackIDs(ids ...int) *ArtistUpdate {
	au.mutation.RemoveTrackIDs(ids...)
	return au
}

// RemoveTracks removes "tracks" edges to Track entities.
func (au *ArtistUpdate) RemoveTracks(t ...*Track) *ArtistUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTrackIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtistUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtistUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtistUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtistUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArtistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   artist.Table,
			Columns: artist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: artist.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if au.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: []string{artist.AlbumsColumn},
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
	if nodes := au.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !au.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: []string{artist.AlbumsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: []string{artist.AlbumsColumn},
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
	if au.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: track.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTracksIDs(); len(nodes) > 0 && !au.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: track.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: track.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ArtistUpdateOne is the builder for updating a single Artist entity.
type ArtistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtistMutation
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (auo *ArtistUpdateOne) AddAlbumIDs(ids ...int) *ArtistUpdateOne {
	auo.mutation.AddAlbumIDs(ids...)
	return auo
}

// AddAlbums adds the "albums" edges to the Album entity.
func (auo *ArtistUpdateOne) AddAlbums(a ...*Album) *ArtistUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAlbumIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (auo *ArtistUpdateOne) AddTrackIDs(ids ...int) *ArtistUpdateOne {
	auo.mutation.AddTrackIDs(ids...)
	return auo
}

// AddTracks adds the "tracks" edges to the Track entity.
func (auo *ArtistUpdateOne) AddTracks(t ...*Track) *ArtistUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTrackIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (auo *ArtistUpdateOne) Mutation() *ArtistMutation {
	return auo.mutation
}

// ClearAlbums clears all "albums" edges to the Album entity.
func (auo *ArtistUpdateOne) ClearAlbums() *ArtistUpdateOne {
	auo.mutation.ClearAlbums()
	return auo
}

// RemoveAlbumIDs removes the "albums" edge to Album entities by IDs.
func (auo *ArtistUpdateOne) RemoveAlbumIDs(ids ...int) *ArtistUpdateOne {
	auo.mutation.RemoveAlbumIDs(ids...)
	return auo
}

// RemoveAlbums removes "albums" edges to Album entities.
func (auo *ArtistUpdateOne) RemoveAlbums(a ...*Album) *ArtistUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAlbumIDs(ids...)
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (auo *ArtistUpdateOne) ClearTracks() *ArtistUpdateOne {
	auo.mutation.ClearTracks()
	return auo
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (auo *ArtistUpdateOne) RemoveTrackIDs(ids ...int) *ArtistUpdateOne {
	auo.mutation.RemoveTrackIDs(ids...)
	return auo
}

// RemoveTracks removes "tracks" edges to Track entities.
func (auo *ArtistUpdateOne) RemoveTracks(t ...*Track) *ArtistUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTrackIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtistUpdateOne) Select(field string, fields ...string) *ArtistUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artist entity.
func (auo *ArtistUpdateOne) Save(ctx context.Context) (*Artist, error) {
	var (
		err  error
		node *Artist
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArtistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Artist)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ArtistMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtistUpdateOne) SaveX(ctx context.Context) *Artist {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtistUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtistUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArtistUpdateOne) sqlSave(ctx context.Context) (_node *Artist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   artist.Table,
			Columns: artist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: artist.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Artist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artist.FieldID)
		for _, f := range fields {
			if !artist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: []string{artist.AlbumsColumn},
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
	if nodes := auo.mutation.RemovedAlbumsIDs(); len(nodes) > 0 && !auo.mutation.AlbumsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: []string{artist.AlbumsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: []string{artist.AlbumsColumn},
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
	if auo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: track.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTracksIDs(); len(nodes) > 0 && !auo.mutation.TracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: track.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: track.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Artist{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
