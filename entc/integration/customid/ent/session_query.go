// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/device"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/entc/integration/customid/ent/session"
	"entgo.io/ent/schema/field"
)

// SessionQuery is the builder for querying Session entities.
type SessionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Session
	withDevice *DeviceQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SessionQuery builder.
func (sq *SessionQuery) Where(ps ...predicate.Session) *SessionQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SessionQuery) Limit(limit int) *SessionQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SessionQuery) Offset(offset int) *SessionQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SessionQuery) Unique(unique bool) *SessionQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SessionQuery) Order(o ...OrderFunc) *SessionQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryDevice chains the current query on the "device" edge.
func (sq *SessionQuery) QueryDevice() *DeviceQuery {
	query := &DeviceQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, session.DeviceTable, session.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Session entity from the query.
// Returns a *NotFoundError when no Session was found.
func (sq *SessionQuery) First(ctx context.Context) (*Session, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{session.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SessionQuery) FirstX(ctx context.Context) *Session {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Session ID from the query.
// Returns a *NotFoundError when no Session ID was found.
func (sq *SessionQuery) FirstID(ctx context.Context) (id schema.ID, err error) {
	var ids []schema.ID
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{session.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SessionQuery) FirstIDX(ctx context.Context) schema.ID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Session entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Session entity is found.
// Returns a *NotFoundError when no Session entities are found.
func (sq *SessionQuery) Only(ctx context.Context) (*Session, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{session.Label}
	default:
		return nil, &NotSingularError{session.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SessionQuery) OnlyX(ctx context.Context) *Session {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Session ID in the query.
// Returns a *NotSingularError when more than one Session ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SessionQuery) OnlyID(ctx context.Context) (id schema.ID, err error) {
	var ids []schema.ID
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{session.Label}
	default:
		err = &NotSingularError{session.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SessionQuery) OnlyIDX(ctx context.Context) schema.ID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Sessions.
func (sq *SessionQuery) All(ctx context.Context) ([]*Session, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SessionQuery) AllX(ctx context.Context) []*Session {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Session IDs.
func (sq *SessionQuery) IDs(ctx context.Context) ([]schema.ID, error) {
	var ids []schema.ID
	if err := sq.Select(session.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SessionQuery) IDsX(ctx context.Context) []schema.ID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SessionQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SessionQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SessionQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SessionQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SessionQuery) Clone() *SessionQuery {
	if sq == nil {
		return nil
	}
	return &SessionQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		predicates: append([]predicate.Session{}, sq.predicates...),
		withDevice: sq.withDevice.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SessionQuery) WithDevice(opts ...func(*DeviceQuery)) *SessionQuery {
	query := &DeviceQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withDevice = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *SessionQuery) GroupBy(field string, fields ...string) *SessionGroupBy {
	grbuild := &SessionGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = session.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sq *SessionQuery) Select(fields ...string) *SessionSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SessionSelect{SessionQuery: sq}
	selbuild.label = session.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

func (sq *SessionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !session.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Session, error) {
	var (
		nodes       = []*Session{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withDevice != nil,
		}
	)
	if sq.withDevice != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, session.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Session).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Session{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withDevice; query != nil {
		if err := sq.loadDevice(ctx, query, nodes, nil,
			func(n *Session, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SessionQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*Session, init func(*Session), assign func(*Session, *Device)) error {
	ids := make([]schema.ID, 0, len(nodes))
	nodeids := make(map[schema.ID][]*Session)
	for i := range nodes {
		if nodes[i].device_sessions == nil {
			continue
		}
		fk := *nodes[i].device_sessions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_sessions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SessionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sq *SessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeBytes,
				Column: session.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for i := range fields {
			if fields[i] != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(session.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = session.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SessionGroupBy is the group-by builder for Session entities.
type SessionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SessionGroupBy) Aggregate(fns ...AggregateFunc) *SessionGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SessionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SessionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgb.fields {
		if !session.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SessionGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SessionSelect is the builder for selecting fields of Session entities.
type SessionSelect struct {
	*SessionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SessionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SessionQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SessionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
