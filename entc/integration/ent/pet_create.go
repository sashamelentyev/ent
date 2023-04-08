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
	"entgo.io/ent/entc/integration/ent/pet"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAge sets the "age" field.
func (pc *PetCreate) SetAge(f float64) *PetCreate {
	pc.mutation.SetAge(f)
	return pc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (pc *PetCreate) SetNillableAge(f *float64) *PetCreate {
	if f != nil {
		pc.SetAge(*f)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetUUID sets the "uuid" field.
func (pc *PetCreate) SetUUID(u uuid.UUID) *PetCreate {
	pc.mutation.SetUUID(u)
	return pc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (pc *PetCreate) SetNillableUUID(u *uuid.UUID) *PetCreate {
	if u != nil {
		pc.SetUUID(*u)
	}
	return pc
}

// SetNickname sets the "nickname" field.
func (pc *PetCreate) SetNickname(s string) *PetCreate {
	pc.mutation.SetNickname(s)
	return pc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (pc *PetCreate) SetNillableNickname(s *string) *PetCreate {
	if s != nil {
		pc.SetNickname(*s)
	}
	return pc
}

// SetTrained sets the "trained" field.
func (pc *PetCreate) SetTrained(b bool) *PetCreate {
	pc.mutation.SetTrained(b)
	return pc
}

// SetNillableTrained sets the "trained" field if the given value is not nil.
func (pc *PetCreate) SetNillableTrained(b *bool) *PetCreate {
	if b != nil {
		pc.SetTrained(*b)
	}
	return pc
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (pc *PetCreate) SetTeamID(id int) *PetCreate {
	pc.mutation.SetTeamID(id)
	return pc
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (pc *PetCreate) SetNillableTeamID(id *int) *PetCreate {
	if id != nil {
		pc = pc.SetTeamID(*id)
	}
	return pc
}

// SetTeam sets the "team" edge to the User entity.
func (pc *PetCreate) SetTeam(u *User) *PetCreate {
	return pc.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PetCreate) SetOwnerID(id int) *PetCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(id *int) *PetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pc *PetCreate) Mutation() *PetMutation {
	return pc.mutation
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	pc.defaults()
	return withHooks[*Pet, PetMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PetCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PetCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PetCreate) defaults() {
	if _, ok := pc.mutation.Age(); !ok {
		v := pet.DefaultAge
		pc.mutation.SetAge(v)
	}
	if _, ok := pc.mutation.Trained(); !ok {
		v := pet.DefaultTrained
		pc.mutation.SetTrained(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PetCreate) check() error {
	if _, ok := pc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Pet.age"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	if _, ok := pc.mutation.Trained(); !ok {
		return &ValidationError{Name: "trained", err: errors.New(`ent: missing required field "Pet.trained"`)}
	}
	return nil
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pet.Table, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Age(); ok {
		_spec.SetField(pet.FieldAge, field.TypeFloat64, value)
		_node.Age = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.UUID(); ok {
		_spec.SetField(pet.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := pc.mutation.Nickname(); ok {
		_spec.SetField(pet.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := pc.mutation.Trained(); ok {
		_spec.SetField(pet.FieldTrained, field.TypeBool, value)
		_node.Trained = value
	}
	if nodes := pc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.TeamTable,
			Columns: []string{pet.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_team = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_pets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pet.Create().
//		SetAge(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PetUpsert) {
//			SetAge(v+v).
//		}).
//		Exec(ctx)
func (pc *PetCreate) OnConflict(opts ...sql.ConflictOption) *PetUpsertOne {
	pc.conflict = opts
	return &PetUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PetCreate) OnConflictColumns(columns ...string) *PetUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PetUpsertOne{
		create: pc,
	}
}

type (
	// PetUpsertOne is the builder for "upsert"-ing
	//  one Pet node.
	PetUpsertOne struct {
		create *PetCreate
	}

	// PetUpsert is the "OnConflict" setter.
	PetUpsert struct {
		*sql.UpdateSet
	}
)

// SetAge sets the "age" field.
func (u *PetUpsert) SetAge(v float64) *PetUpsert {
	u.Set(pet.FieldAge, v)
	return u
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *PetUpsert) UpdateAge() *PetUpsert {
	u.SetExcluded(pet.FieldAge)
	return u
}

// AddAge adds v to the "age" field.
func (u *PetUpsert) AddAge(v float64) *PetUpsert {
	u.Add(pet.FieldAge, v)
	return u
}

// SetName sets the "name" field.
func (u *PetUpsert) SetName(v string) *PetUpsert {
	u.Set(pet.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PetUpsert) UpdateName() *PetUpsert {
	u.SetExcluded(pet.FieldName)
	return u
}

// SetUUID sets the "uuid" field.
func (u *PetUpsert) SetUUID(v uuid.UUID) *PetUpsert {
	u.Set(pet.FieldUUID, v)
	return u
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *PetUpsert) UpdateUUID() *PetUpsert {
	u.SetExcluded(pet.FieldUUID)
	return u
}

// ClearUUID clears the value of the "uuid" field.
func (u *PetUpsert) ClearUUID() *PetUpsert {
	u.SetNull(pet.FieldUUID)
	return u
}

// SetNickname sets the "nickname" field.
func (u *PetUpsert) SetNickname(v string) *PetUpsert {
	u.Set(pet.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *PetUpsert) UpdateNickname() *PetUpsert {
	u.SetExcluded(pet.FieldNickname)
	return u
}

// ClearNickname clears the value of the "nickname" field.
func (u *PetUpsert) ClearNickname() *PetUpsert {
	u.SetNull(pet.FieldNickname)
	return u
}

// SetTrained sets the "trained" field.
func (u *PetUpsert) SetTrained(v bool) *PetUpsert {
	u.Set(pet.FieldTrained, v)
	return u
}

// UpdateTrained sets the "trained" field to the value that was provided on create.
func (u *PetUpsert) UpdateTrained() *PetUpsert {
	u.SetExcluded(pet.FieldTrained)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PetUpsertOne) UpdateNewValues() *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pet.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PetUpsertOne) Ignore() *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PetUpsertOne) DoNothing() *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PetCreate.OnConflict
// documentation for more info.
func (u *PetUpsertOne) Update(set func(*PetUpsert)) *PetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PetUpsert{UpdateSet: update})
	}))
	return u
}

// SetAge sets the "age" field.
func (u *PetUpsertOne) SetAge(v float64) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *PetUpsertOne) AddAge(v float64) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateAge() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateAge()
	})
}

// SetName sets the "name" field.
func (u *PetUpsertOne) SetName(v string) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateName() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateName()
	})
}

// SetUUID sets the "uuid" field.
func (u *PetUpsertOne) SetUUID(v uuid.UUID) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateUUID() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateUUID()
	})
}

// ClearUUID clears the value of the "uuid" field.
func (u *PetUpsertOne) ClearUUID() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.ClearUUID()
	})
}

// SetNickname sets the "nickname" field.
func (u *PetUpsertOne) SetNickname(v string) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateNickname() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *PetUpsertOne) ClearNickname() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.ClearNickname()
	})
}

// SetTrained sets the "trained" field.
func (u *PetUpsertOne) SetTrained(v bool) *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.SetTrained(v)
	})
}

// UpdateTrained sets the "trained" field to the value that was provided on create.
func (u *PetUpsertOne) UpdateTrained() *PetUpsertOne {
	return u.Update(func(s *PetUpsert) {
		s.UpdateTrained()
	})
}

// Exec executes the query.
func (u *PetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	builders []*PetCreate
	conflict []sql.ConflictOption
}

// Save creates the Pet entities in the database.
func (pcb *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pet, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PetCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Pet.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PetUpsert) {
//			SetAge(v+v).
//		}).
//		Exec(ctx)
func (pcb *PetCreateBulk) OnConflict(opts ...sql.ConflictOption) *PetUpsertBulk {
	pcb.conflict = opts
	return &PetUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PetCreateBulk) OnConflictColumns(columns ...string) *PetUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PetUpsertBulk{
		create: pcb,
	}
}

// PetUpsertBulk is the builder for "upsert"-ing
// a bulk of Pet nodes.
type PetUpsertBulk struct {
	create *PetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PetUpsertBulk) UpdateNewValues() *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Pet.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PetUpsertBulk) Ignore() *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PetUpsertBulk) DoNothing() *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PetCreateBulk.OnConflict
// documentation for more info.
func (u *PetUpsertBulk) Update(set func(*PetUpsert)) *PetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PetUpsert{UpdateSet: update})
	}))
	return u
}

// SetAge sets the "age" field.
func (u *PetUpsertBulk) SetAge(v float64) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetAge(v)
	})
}

// AddAge adds v to the "age" field.
func (u *PetUpsertBulk) AddAge(v float64) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.AddAge(v)
	})
}

// UpdateAge sets the "age" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateAge() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateAge()
	})
}

// SetName sets the "name" field.
func (u *PetUpsertBulk) SetName(v string) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateName() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateName()
	})
}

// SetUUID sets the "uuid" field.
func (u *PetUpsertBulk) SetUUID(v uuid.UUID) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateUUID() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateUUID()
	})
}

// ClearUUID clears the value of the "uuid" field.
func (u *PetUpsertBulk) ClearUUID() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.ClearUUID()
	})
}

// SetNickname sets the "nickname" field.
func (u *PetUpsertBulk) SetNickname(v string) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateNickname() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *PetUpsertBulk) ClearNickname() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.ClearNickname()
	})
}

// SetTrained sets the "trained" field.
func (u *PetUpsertBulk) SetTrained(v bool) *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.SetTrained(v)
	})
}

// UpdateTrained sets the "trained" field to the value that was provided on create.
func (u *PetUpsertBulk) UpdateTrained() *PetUpsertBulk {
	return u.Update(func(s *PetUpsert) {
		s.UpdateTrained()
	})
}

// Exec executes the query.
func (u *PetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
