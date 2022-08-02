// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/sid"
)

// IntSID is the model entity for the IntSID schema.
type IntSID struct {
	config
	// ID of the ent.
	ID sid.ID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IntSIDQuery when eager-loading is set.
	Edges          IntSIDEdges `json:"edges"`
	int_sid_parent *sid.ID
}

// IntSIDEdges holds the relations/edges for other nodes in the graph.
type IntSIDEdges struct {
	// Parent holds the value of the parent edge.
	Parent *IntSID `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*IntSID `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e IntSIDEdges) ParentOrErr() (*IntSID, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: intsid.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e IntSIDEdges) ChildrenOrErr() ([]*IntSID, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IntSID) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case intsid.FieldID:
			values[i] = new(sid.ID)
		case intsid.ForeignKeys[0]: // int_sid_parent
			values[i] = &sql.NullScanner{S: new(sid.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type IntSID", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IntSID fields.
func (is *IntSID) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case intsid.FieldID:
			if value, ok := values[i].(*sid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				is.ID = *value
			}
		case intsid.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field int_sid_parent", values[i])
			} else if value.Valid {
				is.int_sid_parent = new(sid.ID)
				*is.int_sid_parent = *value.S.(*sid.ID)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the IntSID entity.
func (is *IntSID) QueryParent() *IntSIDQuery {
	return (&IntSIDClient{config: is.config}).QueryParent(is)
}

// QueryChildren queries the "children" edge of the IntSID entity.
func (is *IntSID) QueryChildren() *IntSIDQuery {
	return (&IntSIDClient{config: is.config}).QueryChildren(is)
}

// Update returns a builder for updating this IntSID.
// Note that you need to call IntSID.Unwrap() before calling this method if this IntSID
// was returned from a transaction, and the transaction was committed or rolled back.
func (is *IntSID) Update() *IntSIDUpdateOne {
	return (&IntSIDClient{config: is.config}).UpdateOne(is)
}

// Unwrap unwraps the IntSID entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (is *IntSID) Unwrap() *IntSID {
	_tx, ok := is.config.driver.(*txDriver)
	if !ok {
		panic("ent: IntSID is not a transactional entity")
	}
	is.config.driver = _tx.drv
	return is
}

// String implements the fmt.Stringer.
func (is *IntSID) String() string {
	var builder strings.Builder
	builder.WriteString("IntSID(")
	builder.WriteString(fmt.Sprintf("id=%v", is.ID))
	builder.WriteByte(')')
	return builder.String()
}

// IntSIDs is a parsable slice of IntSID.
type IntSIDs []*IntSID

func (is IntSIDs) config(cfg config) {
	for _i := range is {
		is[_i].config = cfg
	}
}
