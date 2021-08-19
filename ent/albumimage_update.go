// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"goreact/ent/albumimage"
	"goreact/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AlbumImageUpdate is the builder for updating AlbumImage entities.
type AlbumImageUpdate struct {
	config
	hooks    []Hook
	mutation *AlbumImageMutation
}

// Where adds a new predicate for the AlbumImageUpdate builder.
func (aiu *AlbumImageUpdate) Where(ps ...predicate.AlbumImage) *AlbumImageUpdate {
	aiu.mutation.predicates = append(aiu.mutation.predicates, ps...)
	return aiu
}

// Mutation returns the AlbumImageMutation object of the builder.
func (aiu *AlbumImageUpdate) Mutation() *AlbumImageMutation {
	return aiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *AlbumImageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aiu.hooks) == 0 {
		affected, err = aiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aiu.mutation = mutation
			affected, err = aiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aiu.hooks) - 1; i >= 0; i-- {
			mut = aiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *AlbumImageUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *AlbumImageUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *AlbumImageUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aiu *AlbumImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   albumimage.Table,
			Columns: albumimage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: albumimage.FieldID,
			},
		},
	}
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{albumimage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AlbumImageUpdateOne is the builder for updating a single AlbumImage entity.
type AlbumImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AlbumImageMutation
}

// Mutation returns the AlbumImageMutation object of the builder.
func (aiuo *AlbumImageUpdateOne) Mutation() *AlbumImageMutation {
	return aiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *AlbumImageUpdateOne) Select(field string, fields ...string) *AlbumImageUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated AlbumImage entity.
func (aiuo *AlbumImageUpdateOne) Save(ctx context.Context) (*AlbumImage, error) {
	var (
		err  error
		node *AlbumImage
	)
	if len(aiuo.hooks) == 0 {
		node, err = aiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aiuo.mutation = mutation
			node, err = aiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aiuo.hooks) - 1; i >= 0; i-- {
			mut = aiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *AlbumImageUpdateOne) SaveX(ctx context.Context) *AlbumImage {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *AlbumImageUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *AlbumImageUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aiuo *AlbumImageUpdateOne) sqlSave(ctx context.Context) (_node *AlbumImage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   albumimage.Table,
			Columns: albumimage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: albumimage.FieldID,
			},
		},
	}
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AlbumImage.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, albumimage.FieldID)
		for _, f := range fields {
			if !albumimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != albumimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &AlbumImage{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{albumimage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
