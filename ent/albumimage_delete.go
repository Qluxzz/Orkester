// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
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
	return withHooks(ctx, aid.sqlExec, aid.mutation, aid.hooks)
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
	_spec := sqlgraph.NewDeleteSpec(albumimage.Table, sqlgraph.NewFieldSpec(albumimage.FieldID, field.TypeInt))
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
	aid.mutation.done = true
	return affected, err
}

// AlbumImageDeleteOne is the builder for deleting a single AlbumImage entity.
type AlbumImageDeleteOne struct {
	aid *AlbumImageDelete
}

// Where appends a list predicates to the AlbumImageDelete builder.
func (aido *AlbumImageDeleteOne) Where(ps ...predicate.AlbumImage) *AlbumImageDeleteOne {
	aido.aid.mutation.Where(ps...)
	return aido
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
	if err := aido.Exec(ctx); err != nil {
		panic(err)
	}
}
