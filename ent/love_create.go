// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ArifulProtik/BlackPen/ent/love"
	"github.com/ArifulProtik/BlackPen/ent/user"
	"github.com/google/uuid"
)

// LoveCreate is the builder for creating a Love entity.
type LoveCreate struct {
	config
	mutation *LoveMutation
	hooks    []Hook
}

// SetNoteid sets the "noteid" field.
func (lc *LoveCreate) SetNoteid(u uuid.UUID) *LoveCreate {
	lc.mutation.SetNoteid(u)
	return lc
}

// SetCreatedAt sets the "created_at" field.
func (lc *LoveCreate) SetCreatedAt(t time.Time) *LoveCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LoveCreate) SetNillableCreatedAt(t *time.Time) *LoveCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lc *LoveCreate) SetUserID(id uuid.UUID) *LoveCreate {
	lc.mutation.SetUserID(id)
	return lc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lc *LoveCreate) SetNillableUserID(id *uuid.UUID) *LoveCreate {
	if id != nil {
		lc = lc.SetUserID(*id)
	}
	return lc
}

// SetUser sets the "user" edge to the User entity.
func (lc *LoveCreate) SetUser(u *User) *LoveCreate {
	return lc.SetUserID(u.ID)
}

// Mutation returns the LoveMutation object of the builder.
func (lc *LoveCreate) Mutation() *LoveMutation {
	return lc.mutation
}

// Save creates the Love in the database.
func (lc *LoveCreate) Save(ctx context.Context) (*Love, error) {
	var (
		err  error
		node *Love
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoveMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LoveCreate) SaveX(ctx context.Context) *Love {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LoveCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LoveCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LoveCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := love.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LoveCreate) check() error {
	if _, ok := lc.mutation.Noteid(); !ok {
		return &ValidationError{Name: "noteid", err: errors.New(`ent: missing required field "Love.noteid"`)}
	}
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Love.created_at"`)}
	}
	return nil
}

func (lc *LoveCreate) sqlSave(ctx context.Context) (*Love, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lc *LoveCreate) createSpec() (*Love, *sqlgraph.CreateSpec) {
	var (
		_node = &Love{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: love.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: love.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.Noteid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: love.FieldNoteid,
		})
		_node.Noteid = value
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: love.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   love.UserTable,
			Columns: []string{love.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_loves = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LoveCreateBulk is the builder for creating many Love entities in bulk.
type LoveCreateBulk struct {
	config
	builders []*LoveCreate
}

// Save creates the Love entities in the database.
func (lcb *LoveCreateBulk) Save(ctx context.Context) ([]*Love, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Love, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoveMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LoveCreateBulk) SaveX(ctx context.Context) []*Love {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LoveCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LoveCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
