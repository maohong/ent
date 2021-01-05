// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/migrate/entv1/customtype"
	"github.com/facebook/ent/schema/field"
)

// CustomTypeCreate is the builder for creating a CustomType entity.
type CustomTypeCreate struct {
	config
	mutation *CustomTypeMutation
	hooks    []Hook
}

// SetCustom sets the "custom" field.
func (ctc *CustomTypeCreate) SetCustom(s string) *CustomTypeCreate {
	ctc.mutation.SetCustom(s)
	return ctc
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (ctc *CustomTypeCreate) SetNillableCustom(s *string) *CustomTypeCreate {
	if s != nil {
		ctc.SetCustom(*s)
	}
	return ctc
}

// Mutation returns the CustomTypeMutation object of the builder.
func (ctc *CustomTypeCreate) Mutation() *CustomTypeMutation {
	return ctc.mutation
}

// Save creates the CustomType in the database.
func (ctc *CustomTypeCreate) Save(ctx context.Context) (*CustomType, error) {
	var (
		err  error
		node *CustomType
	)
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			node, err = ctc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			mut = ctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CustomTypeCreate) SaveX(ctx context.Context) *CustomType {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CustomTypeCreate) check() error {
	return nil
}

func (ctc *CustomTypeCreate) sqlSave(ctx context.Context) (*CustomType, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ctc *CustomTypeCreate) createSpec() (*CustomType, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomType{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customtype.FieldID,
			},
		}
	)
	if value, ok := ctc.mutation.Custom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customtype.FieldCustom,
		})
		_node.Custom = value
	}
	return _node, _spec
}

// CustomTypeCreateBulk is the builder for creating many CustomType entities in bulk.
type CustomTypeCreateBulk struct {
	config
	builders []*CustomTypeCreate
}

// Save creates the CustomType entities in the database.
func (ctcb *CustomTypeCreateBulk) Save(ctx context.Context) ([]*CustomType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CustomType, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CustomTypeCreateBulk) SaveX(ctx context.Context) []*CustomType {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}