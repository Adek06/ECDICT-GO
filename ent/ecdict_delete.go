// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ECDICT-GO/ent/ecdict"
	"ECDICT-GO/ent/predicate"
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// EcdictDelete is the builder for deleting a Ecdict entity.
type EcdictDelete struct {
	config
	hooks    []Hook
	mutation *EcdictMutation
}

// Where adds a new predicate to the delete builder.
func (ed *EcdictDelete) Where(ps ...predicate.Ecdict) *EcdictDelete {
	ed.mutation.predicates = append(ed.mutation.predicates, ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EcdictDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ed.hooks) == 0 {
		affected, err = ed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EcdictMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ed.mutation = mutation
			affected, err = ed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ed.hooks) - 1; i >= 0; i-- {
			mut = ed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EcdictDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EcdictDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: ecdict.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ecdict.FieldID,
			},
		},
	}
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
}

// EcdictDeleteOne is the builder for deleting a single Ecdict entity.
type EcdictDeleteOne struct {
	ed *EcdictDelete
}

// Exec executes the deletion query.
func (edo *EcdictDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ecdict.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EcdictDeleteOne) ExecX(ctx context.Context) {
	edo.ed.ExecX(ctx)
}