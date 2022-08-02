// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/node"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// NodeQuery is the builder for querying Node entities.
type NodeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Node
	withPrev   *NodeQuery
	withNext   *NodeQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the NodeQuery builder.
func (nq *NodeQuery) Where(ps ...predicate.Node) *NodeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NodeQuery) Limit(limit int) *NodeQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NodeQuery) Offset(offset int) *NodeQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NodeQuery) Unique(unique bool) *NodeQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NodeQuery) Order(o ...OrderFunc) *NodeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryPrev chains the current query on the "prev" edge.
func (nq *NodeQuery) QueryPrev() *NodeQuery {
	query := &NodeQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := nq.gremlinQuery(ctx)
		fromU = gremlin.InE(node.NextLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryNext chains the current query on the "next" edge.
func (nq *NodeQuery) QueryNext() *NodeQuery {
	query := &NodeQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := nq.gremlinQuery(ctx)
		fromU = gremlin.OutE(node.NextLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first Node entity from the query.
// Returns a *NotFoundError when no Node was found.
func (nq *NodeQuery) First(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{node.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NodeQuery) FirstX(ctx context.Context) *Node {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Node ID from the query.
// Returns a *NotFoundError when no Node ID was found.
func (nq *NodeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{node.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NodeQuery) FirstIDX(ctx context.Context) string {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Node entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Node entity is found.
// Returns a *NotFoundError when no Node entities are found.
func (nq *NodeQuery) Only(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{node.Label}
	default:
		return nil, &NotSingularError{node.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NodeQuery) OnlyX(ctx context.Context) *Node {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Node ID in the query.
// Returns a *NotSingularError when more than one Node ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NodeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{node.Label}
	default:
		err = &NotSingularError{node.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NodeQuery) OnlyIDX(ctx context.Context) string {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nodes.
func (nq *NodeQuery) All(ctx context.Context) ([]*Node, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NodeQuery) AllX(ctx context.Context) []*Node {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Node IDs.
func (nq *NodeQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := nq.Select(node.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NodeQuery) IDsX(ctx context.Context) []string {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NodeQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NodeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NodeQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NodeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NodeQuery) Clone() *NodeQuery {
	if nq == nil {
		return nil
	}
	return &NodeQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		predicates: append([]predicate.Node{}, nq.predicates...),
		withPrev:   nq.withPrev.Clone(),
		withNext:   nq.withNext.Clone(),
		// clone intermediate query.
		gremlin: nq.gremlin.Clone(),
		path:    nq.path,
		unique:  nq.unique,
	}
}

// WithPrev tells the query-builder to eager-load the nodes that are connected to
// the "prev" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithPrev(opts ...func(*NodeQuery)) *NodeQuery {
	query := &NodeQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withPrev = query
	return nq
}

// WithNext tells the query-builder to eager-load the nodes that are connected to
// the "next" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithNext(opts ...func(*NodeQuery)) *NodeQuery {
	query := &NodeQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withNext = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Node.Query().
//		GroupBy(node.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nq *NodeQuery) GroupBy(field string, fields ...string) *NodeGroupBy {
	grbuild := &NodeGroupBy{config: nq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.gremlinQuery(ctx), nil
	}
	grbuild.label = node.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//	}
//
//	client.Node.Query().
//		Select(node.FieldValue).
//		Scan(ctx, &v)
//
func (nq *NodeQuery) Select(fields ...string) *NodeSelect {
	nq.fields = append(nq.fields, fields...)
	selbuild := &NodeSelect{NodeQuery: nq}
	selbuild.label = node.Label
	selbuild.flds, selbuild.scan = &nq.fields, selbuild.Scan
	return selbuild
}

func (nq *NodeQuery) prepareQuery(ctx context.Context) error {
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.gremlin = prev
	}
	return nil
}

func (nq *NodeQuery) gremlinAll(ctx context.Context) ([]*Node, error) {
	res := &gremlin.Response{}
	traversal := nq.gremlinQuery(ctx)
	if len(nq.fields) > 0 {
		fields := make([]interface{}, len(nq.fields))
		for i, f := range nq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := nq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var ns Nodes
	if err := ns.FromResponse(res); err != nil {
		return nil, err
	}
	ns.config(nq.config)
	return ns, nil
}

func (nq *NodeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := nq.gremlinQuery(ctx).Count().Query()
	if err := nq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (nq *NodeQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := nq.gremlinQuery(ctx).HasNext().Query()
	if err := nq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (nq *NodeQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(node.Label)
	if nq.gremlin != nil {
		v = nq.gremlin.Clone()
	}
	for _, p := range nq.predicates {
		p(v)
	}
	if len(nq.order) > 0 {
		v.Order()
		for _, p := range nq.order {
			p(v)
		}
	}
	switch limit, offset := nq.limit, nq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := nq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// NodeGroupBy is the group-by builder for Node entities.
type NodeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NodeGroupBy) Aggregate(fns ...AggregateFunc) *NodeGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NodeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.gremlin = query
	return ngb.gremlinScan(ctx, v)
}

func (ngb *NodeGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ngb.gremlinQuery().Query()
	if err := ngb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ngb.fields)+len(ngb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ngb *NodeGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ngb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ngb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ngb.gremlin.Group().
		By(__.Values(ngb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// NodeSelect is the builder for selecting fields of Node entities.
type NodeSelect struct {
	*NodeQuery
	selector
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NodeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.gremlin = ns.NodeQuery.gremlinQuery(ctx)
	return ns.gremlinScan(ctx, v)
}

func (ns *NodeSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(ns.fields) == 1 {
		if ns.fields[0] != node.FieldID {
			traversal = ns.gremlin.Values(ns.fields...)
		} else {
			traversal = ns.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(ns.fields))
		for i, f := range ns.fields {
			fields[i] = f
		}
		traversal = ns.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ns.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ns.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
