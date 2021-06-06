// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/track"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrackCreate is the builder for creating a Track entity.
type TrackCreate struct {
	config
	mutation *TrackMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (tc *TrackCreate) SetTitle(s string) *TrackCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetTrackNumber sets the "track_number" field.
func (tc *TrackCreate) SetTrackNumber(i int) *TrackCreate {
	tc.mutation.SetTrackNumber(i)
	return tc
}

// SetPath sets the "path" field.
func (tc *TrackCreate) SetPath(s string) *TrackCreate {
	tc.mutation.SetPath(s)
	return tc
}

// SetDate sets the "date" field.
func (tc *TrackCreate) SetDate(t time.Time) *TrackCreate {
	tc.mutation.SetDate(t)
	return tc
}

// SetLength sets the "length" field.
func (tc *TrackCreate) SetLength(i int) *TrackCreate {
	tc.mutation.SetLength(i)
	return tc
}

// SetMimetype sets the "mimetype" field.
func (tc *TrackCreate) SetMimetype(s string) *TrackCreate {
	tc.mutation.SetMimetype(s)
	return tc
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (tc *TrackCreate) AddArtistIDs(ids ...int) *TrackCreate {
	tc.mutation.AddArtistIDs(ids...)
	return tc
}

// AddArtists adds the "artists" edges to the Artist entity.
func (tc *TrackCreate) AddArtists(a ...*Artist) *TrackCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tc.AddArtistIDs(ids...)
}

// SetAlbumID sets the "album" edge to the Album entity by ID.
func (tc *TrackCreate) SetAlbumID(id int) *TrackCreate {
	tc.mutation.SetAlbumID(id)
	return tc
}

// SetNillableAlbumID sets the "album" edge to the Album entity by ID if the given value is not nil.
func (tc *TrackCreate) SetNillableAlbumID(id *int) *TrackCreate {
	if id != nil {
		tc = tc.SetAlbumID(*id)
	}
	return tc
}

// SetAlbum sets the "album" edge to the Album entity.
func (tc *TrackCreate) SetAlbum(a *Album) *TrackCreate {
	return tc.SetAlbumID(a.ID)
}

// Mutation returns the TrackMutation object of the builder.
func (tc *TrackCreate) Mutation() *TrackMutation {
	return tc.mutation
}

// Save creates the Track in the database.
func (tc *TrackCreate) Save(ctx context.Context) (*Track, error) {
	var (
		err  error
		node *Track
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TrackMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TrackCreate) SaveX(ctx context.Context) *Track {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tc *TrackCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if _, ok := tc.mutation.TrackNumber(); !ok {
		return &ValidationError{Name: "track_number", err: errors.New("ent: missing required field \"track_number\"")}
	}
	if _, ok := tc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New("ent: missing required field \"path\"")}
	}
	if _, ok := tc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New("ent: missing required field \"date\"")}
	}
	if _, ok := tc.mutation.Length(); !ok {
		return &ValidationError{Name: "length", err: errors.New("ent: missing required field \"length\"")}
	}
	if _, ok := tc.mutation.Mimetype(); !ok {
		return &ValidationError{Name: "mimetype", err: errors.New("ent: missing required field \"mimetype\"")}
	}
	if len(tc.mutation.ArtistsIDs()) == 0 {
		return &ValidationError{Name: "artists", err: errors.New("ent: missing required edge \"artists\"")}
	}
	return nil
}

func (tc *TrackCreate) sqlSave(ctx context.Context) (*Track, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TrackCreate) createSpec() (*Track, *sqlgraph.CreateSpec) {
	var (
		_node = &Track{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: track.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: track.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: track.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tc.mutation.TrackNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: track.FieldTrackNumber,
		})
		_node.TrackNumber = value
	}
	if value, ok := tc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: track.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := tc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: track.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := tc.mutation.Length(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: track.FieldLength,
		})
		_node.Length = value
	}
	if value, ok := tc.mutation.Mimetype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: track.FieldMimetype,
		})
		_node.Mimetype = value
	}
	if nodes := tc.mutation.ArtistsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.AlbumIDs(); len(nodes) > 0 {
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
		_node.album_tracks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TrackCreateBulk is the builder for creating many Track entities in bulk.
type TrackCreateBulk struct {
	config
	builders []*TrackCreate
}

// Save creates the Track entities in the database.
func (tcb *TrackCreateBulk) Save(ctx context.Context) ([]*Track, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Track, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TrackMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TrackCreateBulk) SaveX(ctx context.Context) []*Track {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
