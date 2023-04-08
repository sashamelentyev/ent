// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/edgeindex/ent/city"
	"entgo.io/ent/examples/edgeindex/ent/street"
	"entgo.io/ent/schema/field"
)

// StreetCreate is the builder for creating a Street entity.
type StreetCreate struct {
	config
	mutation *StreetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *StreetCreate) SetName(s string) *StreetCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetCityID sets the "city" edge to the City entity by ID.
func (sc *StreetCreate) SetCityID(id int) *StreetCreate {
	sc.mutation.SetCityID(id)
	return sc
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (sc *StreetCreate) SetNillableCityID(id *int) *StreetCreate {
	if id != nil {
		sc = sc.SetCityID(*id)
	}
	return sc
}

// SetCity sets the "city" edge to the City entity.
func (sc *StreetCreate) SetCity(c *City) *StreetCreate {
	return sc.SetCityID(c.ID)
}

// Mutation returns the StreetMutation object of the builder.
func (sc *StreetCreate) Mutation() *StreetMutation {
	return sc.mutation
}

// Save creates the Street in the database.
func (sc *StreetCreate) Save(ctx context.Context) (*Street, error) {
	return withHooks[*Street, StreetMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreetCreate) SaveX(ctx context.Context) *Street {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StreetCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StreetCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StreetCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Street.name"`)}
	}
	return nil
}

func (sc *StreetCreate) sqlSave(ctx context.Context) (*Street, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StreetCreate) createSpec() (*Street, *sqlgraph.CreateSpec) {
	var (
		_node = &Street{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(street.Table, sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(street.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := sc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.city_streets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StreetCreateBulk is the builder for creating many Street entities in bulk.
type StreetCreateBulk struct {
	config
	builders []*StreetCreate
}

// Save creates the Street entities in the database.
func (scb *StreetCreateBulk) Save(ctx context.Context) ([]*Street, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Street, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StreetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
			for j := len(builder.hooks) - 1; j >= 0; j-- {
				mut = builder.hooks[j](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StreetCreateBulk) SaveX(ctx context.Context) []*Street {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StreetCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StreetCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
