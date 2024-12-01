// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"orkester/ent/albumimage"
	"orkester/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AlbumImageDelete is the builder for deleting a AlbumImage entity.
type AlbumImageDelete struct {
	config
	hooks    []Hook
	mutation *AlbumImageMutation
}

// Where appends a list predicates to the AlbumImageDelete builder.
func (aid *AlbumImageDelete) Where(ps ...predicate.AlbumImage) *AlbumImageDelete {
	aid.mutation.Where(ps...)
	return aid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aid *AlbumImageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aid.hooks) == 0 {
		affected, err = aid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aid.mutation = mutation
			affected, err = aid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aid.hooks) - 1; i >= 0; i-- {
			if aid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aid *AlbumImageDelete) ExecX(ctx context.Context) int {
	n, err := aid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aid *AlbumImageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: albumimage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: albumimage.FieldID,
			},
		},
	}
	if ps := aid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AlbumImageDeleteOne is the builder for deleting a single AlbumImage entity.
type AlbumImageDeleteOne struct {
	aid *AlbumImageDelete
}

// Exec executes the deletion query.
func (aido *AlbumImageDeleteOne) Exec(ctx context.Context) error {
	n, err := aido.aid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{albumimage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aido *AlbumImageDeleteOne) ExecX(ctx context.Context) {
	aido.aid.ExecX(ctx)
}
