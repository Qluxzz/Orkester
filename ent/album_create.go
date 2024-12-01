// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"orkester/ent/album"
	"orkester/ent/albumimage"
	"orkester/ent/artist"
	"orkester/ent/track"
	"orkester/indexFiles"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AlbumCreate is the builder for creating a Album entity.
type AlbumCreate struct {
	config
	mutation *AlbumMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AlbumCreate) SetName(s string) *AlbumCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetURLName sets the "url_name" field.
func (ac *AlbumCreate) SetURLName(s string) *AlbumCreate {
	ac.mutation.SetURLName(s)
	return ac
}

// SetReleased sets the "released" field.
func (ac *AlbumCreate) SetReleased(ifd *indexFiles.ReleaseDate) *AlbumCreate {
	ac.mutation.SetReleased(ifd)
	return ac
}

// SetArtistID sets the "artist" edge to the Artist entity by ID.
func (ac *AlbumCreate) SetArtistID(id int) *AlbumCreate {
	ac.mutation.SetArtistID(id)
	return ac
}

// SetArtist sets the "artist" edge to the Artist entity.
func (ac *AlbumCreate) SetArtist(a *Artist) *AlbumCreate {
	return ac.SetArtistID(a.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (ac *AlbumCreate) AddTrackIDs(ids ...int) *AlbumCreate {
	ac.mutation.AddTrackIDs(ids...)
	return ac
}

// AddTracks adds the "tracks" edges to the Track entity.
func (ac *AlbumCreate) AddTracks(t ...*Track) *AlbumCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTrackIDs(ids...)
}

// SetCoverID sets the "cover" edge to the AlbumImage entity by ID.
func (ac *AlbumCreate) SetCoverID(id int) *AlbumCreate {
	ac.mutation.SetCoverID(id)
	return ac
}

// SetNillableCoverID sets the "cover" edge to the AlbumImage entity by ID if the given value is not nil.
func (ac *AlbumCreate) SetNillableCoverID(id *int) *AlbumCreate {
	if id != nil {
		ac = ac.SetCoverID(*id)
	}
	return ac
}

// SetCover sets the "cover" edge to the AlbumImage entity.
func (ac *AlbumCreate) SetCover(a *AlbumImage) *AlbumCreate {
	return ac.SetCoverID(a.ID)
}

// Mutation returns the AlbumMutation object of the builder.
func (ac *AlbumCreate) Mutation() *AlbumMutation {
	return ac.mutation
}

// Save creates the Album in the database.
func (ac *AlbumCreate) Save(ctx context.Context) (*Album, error) {
	var (
		err  error
		node *Album
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Album)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AlbumMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AlbumCreate) SaveX(ctx context.Context) *Album {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AlbumCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AlbumCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AlbumCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Album.name"`)}
	}
	if _, ok := ac.mutation.URLName(); !ok {
		return &ValidationError{Name: "url_name", err: errors.New(`ent: missing required field "Album.url_name"`)}
	}
	if _, ok := ac.mutation.Released(); !ok {
		return &ValidationError{Name: "released", err: errors.New(`ent: missing required field "Album.released"`)}
	}
	if _, ok := ac.mutation.ArtistID(); !ok {
		return &ValidationError{Name: "artist", err: errors.New(`ent: missing required edge "Album.artist"`)}
	}
	return nil
}

func (ac *AlbumCreate) sqlSave(ctx context.Context) (*Album, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AlbumCreate) createSpec() (*Album, *sqlgraph.CreateSpec) {
	var (
		_node = &Album{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: album.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: album.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: album.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.URLName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: album.FieldURLName,
		})
		_node.URLName = value
	}
	if value, ok := ac.mutation.Released(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: album.FieldReleased,
		})
		_node.Released = value
	}
	if nodes := ac.mutation.ArtistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   album.ArtistTable,
			Columns: []string{album.ArtistColumn},
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
		_node.artist_albums = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   album.TracksTable,
			Columns: []string{album.TracksColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.CoverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   album.CoverTable,
			Columns: []string{album.CoverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: albumimage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.album_cover = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AlbumCreateBulk is the builder for creating many Album entities in bulk.
type AlbumCreateBulk struct {
	config
	builders []*AlbumCreate
}

// Save creates the Album entities in the database.
func (acb *AlbumCreateBulk) Save(ctx context.Context) ([]*Album, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Album, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlbumMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AlbumCreateBulk) SaveX(ctx context.Context) []*Album {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AlbumCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AlbumCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
