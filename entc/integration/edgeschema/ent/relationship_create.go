// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationship"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// RelationshipCreate is the builder for creating a Relationship entity.
type RelationshipCreate struct {
	config
	mutation *RelationshipMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWeight sets the "weight" field.
func (rc *RelationshipCreate) SetWeight(i int) *RelationshipCreate {
	rc.mutation.SetWeight(i)
	return rc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (rc *RelationshipCreate) SetNillableWeight(i *int) *RelationshipCreate {
	if i != nil {
		rc.SetWeight(*i)
	}
	return rc
}

// SetUserID sets the "user_id" field.
func (rc *RelationshipCreate) SetUserID(i int) *RelationshipCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetRelativeID sets the "relative_id" field.
func (rc *RelationshipCreate) SetRelativeID(i int) *RelationshipCreate {
	rc.mutation.SetRelativeID(i)
	return rc
}

// SetInfoID sets the "info_id" field.
func (rc *RelationshipCreate) SetInfoID(i int) *RelationshipCreate {
	rc.mutation.SetInfoID(i)
	return rc
}

// SetNillableInfoID sets the "info_id" field if the given value is not nil.
func (rc *RelationshipCreate) SetNillableInfoID(i *int) *RelationshipCreate {
	if i != nil {
		rc.SetInfoID(*i)
	}
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *RelationshipCreate) SetUser(u *User) *RelationshipCreate {
	return rc.SetUserID(u.ID)
}

// SetRelative sets the "relative" edge to the User entity.
func (rc *RelationshipCreate) SetRelative(u *User) *RelationshipCreate {
	return rc.SetRelativeID(u.ID)
}

// SetInfo sets the "info" edge to the RelationshipInfo entity.
func (rc *RelationshipCreate) SetInfo(r *RelationshipInfo) *RelationshipCreate {
	return rc.SetInfoID(r.ID)
}

// Mutation returns the RelationshipMutation object of the builder.
func (rc *RelationshipCreate) Mutation() *RelationshipMutation {
	return rc.mutation
}

// Save creates the Relationship in the database.
func (rc *RelationshipCreate) Save(ctx context.Context) (*Relationship, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Relationship, RelationshipMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RelationshipCreate) SaveX(ctx context.Context) *Relationship {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RelationshipCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RelationshipCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RelationshipCreate) defaults() error {
	if _, ok := rc.mutation.Weight(); !ok {
		v := relationship.DefaultWeight
		rc.mutation.SetWeight(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RelationshipCreate) check() error {
	if _, ok := rc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Relationship.weight"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Relationship.user_id"`)}
	}
	if _, ok := rc.mutation.RelativeID(); !ok {
		return &ValidationError{Name: "relative_id", err: errors.New(`ent: missing required field "Relationship.relative_id"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Relationship.user"`)}
	}
	if _, ok := rc.mutation.RelativeID(); !ok {
		return &ValidationError{Name: "relative", err: errors.New(`ent: missing required edge "Relationship.relative"`)}
	}
	return nil
}

func (rc *RelationshipCreate) sqlSave(ctx context.Context) (*Relationship, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (rc *RelationshipCreate) createSpec() (*Relationship, *sqlgraph.CreateSpec) {
	var (
		_node = &Relationship{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(relationship.Table, nil)
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.Weight(); ok {
		_spec.SetField(relationship.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.UserTable,
			Columns: []string{relationship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RelativeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.RelativeTable,
			Columns: []string{relationship.RelativeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelativeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.InfoTable,
			Columns: []string{relationship.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relationshipinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.InfoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Relationship.Create().
//		SetWeight(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RelationshipUpsert) {
//			SetWeight(v+v).
//		}).
//		Exec(ctx)
func (rc *RelationshipCreate) OnConflict(opts ...sql.ConflictOption) *RelationshipUpsertOne {
	rc.conflict = opts
	return &RelationshipUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Relationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RelationshipCreate) OnConflictColumns(columns ...string) *RelationshipUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RelationshipUpsertOne{
		create: rc,
	}
}

type (
	// RelationshipUpsertOne is the builder for "upsert"-ing
	//  one Relationship node.
	RelationshipUpsertOne struct {
		create *RelationshipCreate
	}

	// RelationshipUpsert is the "OnConflict" setter.
	RelationshipUpsert struct {
		*sql.UpdateSet
	}
)

// SetWeight sets the "weight" field.
func (u *RelationshipUpsert) SetWeight(v int) *RelationshipUpsert {
	u.Set(relationship.FieldWeight, v)
	return u
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *RelationshipUpsert) UpdateWeight() *RelationshipUpsert {
	u.SetExcluded(relationship.FieldWeight)
	return u
}

// AddWeight adds v to the "weight" field.
func (u *RelationshipUpsert) AddWeight(v int) *RelationshipUpsert {
	u.Add(relationship.FieldWeight, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *RelationshipUpsert) SetUserID(v int) *RelationshipUpsert {
	u.Set(relationship.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *RelationshipUpsert) UpdateUserID() *RelationshipUpsert {
	u.SetExcluded(relationship.FieldUserID)
	return u
}

// SetRelativeID sets the "relative_id" field.
func (u *RelationshipUpsert) SetRelativeID(v int) *RelationshipUpsert {
	u.Set(relationship.FieldRelativeID, v)
	return u
}

// UpdateRelativeID sets the "relative_id" field to the value that was provided on create.
func (u *RelationshipUpsert) UpdateRelativeID() *RelationshipUpsert {
	u.SetExcluded(relationship.FieldRelativeID)
	return u
}

// SetInfoID sets the "info_id" field.
func (u *RelationshipUpsert) SetInfoID(v int) *RelationshipUpsert {
	u.Set(relationship.FieldInfoID, v)
	return u
}

// UpdateInfoID sets the "info_id" field to the value that was provided on create.
func (u *RelationshipUpsert) UpdateInfoID() *RelationshipUpsert {
	u.SetExcluded(relationship.FieldInfoID)
	return u
}

// ClearInfoID clears the value of the "info_id" field.
func (u *RelationshipUpsert) ClearInfoID() *RelationshipUpsert {
	u.SetNull(relationship.FieldInfoID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Relationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RelationshipUpsertOne) UpdateNewValues() *RelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Relationship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RelationshipUpsertOne) Ignore() *RelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RelationshipUpsertOne) DoNothing() *RelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RelationshipCreate.OnConflict
// documentation for more info.
func (u *RelationshipUpsertOne) Update(set func(*RelationshipUpsert)) *RelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetWeight sets the "weight" field.
func (u *RelationshipUpsertOne) SetWeight(v int) *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetWeight(v)
	})
}

// AddWeight adds v to the "weight" field.
func (u *RelationshipUpsertOne) AddWeight(v int) *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.AddWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *RelationshipUpsertOne) UpdateWeight() *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateWeight()
	})
}

// SetUserID sets the "user_id" field.
func (u *RelationshipUpsertOne) SetUserID(v int) *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *RelationshipUpsertOne) UpdateUserID() *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateUserID()
	})
}

// SetRelativeID sets the "relative_id" field.
func (u *RelationshipUpsertOne) SetRelativeID(v int) *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetRelativeID(v)
	})
}

// UpdateRelativeID sets the "relative_id" field to the value that was provided on create.
func (u *RelationshipUpsertOne) UpdateRelativeID() *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateRelativeID()
	})
}

// SetInfoID sets the "info_id" field.
func (u *RelationshipUpsertOne) SetInfoID(v int) *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetInfoID(v)
	})
}

// UpdateInfoID sets the "info_id" field to the value that was provided on create.
func (u *RelationshipUpsertOne) UpdateInfoID() *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateInfoID()
	})
}

// ClearInfoID clears the value of the "info_id" field.
func (u *RelationshipUpsertOne) ClearInfoID() *RelationshipUpsertOne {
	return u.Update(func(s *RelationshipUpsert) {
		s.ClearInfoID()
	})
}

// Exec executes the query.
func (u *RelationshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RelationshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RelationshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// RelationshipCreateBulk is the builder for creating many Relationship entities in bulk.
type RelationshipCreateBulk struct {
	config
	builders []*RelationshipCreate
	conflict []sql.ConflictOption
}

// Save creates the Relationship entities in the database.
func (rcb *RelationshipCreateBulk) Save(ctx context.Context) ([]*Relationship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Relationship, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RelationshipMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RelationshipCreateBulk) SaveX(ctx context.Context) []*Relationship {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RelationshipCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RelationshipCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Relationship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RelationshipUpsert) {
//			SetWeight(v+v).
//		}).
//		Exec(ctx)
func (rcb *RelationshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *RelationshipUpsertBulk {
	rcb.conflict = opts
	return &RelationshipUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Relationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RelationshipCreateBulk) OnConflictColumns(columns ...string) *RelationshipUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RelationshipUpsertBulk{
		create: rcb,
	}
}

// RelationshipUpsertBulk is the builder for "upsert"-ing
// a bulk of Relationship nodes.
type RelationshipUpsertBulk struct {
	create *RelationshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Relationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RelationshipUpsertBulk) UpdateNewValues() *RelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Relationship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RelationshipUpsertBulk) Ignore() *RelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RelationshipUpsertBulk) DoNothing() *RelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RelationshipCreateBulk.OnConflict
// documentation for more info.
func (u *RelationshipUpsertBulk) Update(set func(*RelationshipUpsert)) *RelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetWeight sets the "weight" field.
func (u *RelationshipUpsertBulk) SetWeight(v int) *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetWeight(v)
	})
}

// AddWeight adds v to the "weight" field.
func (u *RelationshipUpsertBulk) AddWeight(v int) *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.AddWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *RelationshipUpsertBulk) UpdateWeight() *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateWeight()
	})
}

// SetUserID sets the "user_id" field.
func (u *RelationshipUpsertBulk) SetUserID(v int) *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *RelationshipUpsertBulk) UpdateUserID() *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateUserID()
	})
}

// SetRelativeID sets the "relative_id" field.
func (u *RelationshipUpsertBulk) SetRelativeID(v int) *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetRelativeID(v)
	})
}

// UpdateRelativeID sets the "relative_id" field to the value that was provided on create.
func (u *RelationshipUpsertBulk) UpdateRelativeID() *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateRelativeID()
	})
}

// SetInfoID sets the "info_id" field.
func (u *RelationshipUpsertBulk) SetInfoID(v int) *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.SetInfoID(v)
	})
}

// UpdateInfoID sets the "info_id" field to the value that was provided on create.
func (u *RelationshipUpsertBulk) UpdateInfoID() *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.UpdateInfoID()
	})
}

// ClearInfoID clears the value of the "info_id" field.
func (u *RelationshipUpsertBulk) ClearInfoID() *RelationshipUpsertBulk {
	return u.Update(func(s *RelationshipUpsert) {
		s.ClearInfoID()
	})
}

// Exec executes the query.
func (u *RelationshipUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RelationshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RelationshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RelationshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
