// Code generated by entc, DO NOT EDIT.

package ent

import (
	"clock-in/app/worktime/service/internal/data/ent/worktime"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorktimeCreate is the builder for creating a Worktime entity.
type WorktimeCreate struct {
	config
	mutation *WorktimeMutation
	hooks    []Hook
}

// SetUser sets the "user" field.
func (wc *WorktimeCreate) SetUser(i int64) *WorktimeCreate {
	wc.mutation.SetUser(i)
	return wc
}

// SetDay sets the "day" field.
func (wc *WorktimeCreate) SetDay(i int64) *WorktimeCreate {
	wc.mutation.SetDay(i)
	return wc
}

// SetMinute sets the "minute" field.
func (wc *WorktimeCreate) SetMinute(i int64) *WorktimeCreate {
	wc.mutation.SetMinute(i)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorktimeCreate) SetCreatedAt(t time.Time) *WorktimeCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorktimeCreate) SetUpdatedAt(t time.Time) *WorktimeCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorktimeCreate) SetNillableUpdatedAt(t *time.Time) *WorktimeCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// Mutation returns the WorktimeMutation object of the builder.
func (wc *WorktimeCreate) Mutation() *WorktimeMutation {
	return wc.mutation
}

// Save creates the Worktime in the database.
func (wc *WorktimeCreate) Save(ctx context.Context) (*Worktime, error) {
	var (
		err  error
		node *Worktime
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorktimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorktimeCreate) SaveX(ctx context.Context) *Worktime {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorktimeCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorktimeCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorktimeCreate) defaults() {
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := worktime.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorktimeCreate) check() error {
	if _, ok := wc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "user"`)}
	}
	if _, ok := wc.mutation.Day(); !ok {
		return &ValidationError{Name: "day", err: errors.New(`ent: missing required field "day"`)}
	}
	if _, ok := wc.mutation.Minute(); !ok {
		return &ValidationError{Name: "minute", err: errors.New(`ent: missing required field "minute"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	return nil
}

func (wc *WorktimeCreate) sqlSave(ctx context.Context) (*Worktime, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wc *WorktimeCreate) createSpec() (*Worktime, *sqlgraph.CreateSpec) {
	var (
		_node = &Worktime{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: worktime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: worktime.FieldID,
			},
		}
	)
	if value, ok := wc.mutation.User(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: worktime.FieldUser,
		})
		_node.User = value
	}
	if value, ok := wc.mutation.Day(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: worktime.FieldDay,
		})
		_node.Day = value
	}
	if value, ok := wc.mutation.Minute(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: worktime.FieldMinute,
		})
		_node.Minute = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: worktime.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: worktime.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// WorktimeCreateBulk is the builder for creating many Worktime entities in bulk.
type WorktimeCreateBulk struct {
	config
	builders []*WorktimeCreate
}

// Save creates the Worktime entities in the database.
func (wcb *WorktimeCreateBulk) Save(ctx context.Context) ([]*Worktime, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Worktime, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorktimeMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorktimeCreateBulk) SaveX(ctx context.Context) []*Worktime {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorktimeCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorktimeCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
