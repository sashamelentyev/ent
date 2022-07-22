// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
)

// Goods is the model entity for the Goods schema.
type Goods struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Goods.
func (_go *Goods) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_go struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scan_go); err != nil {
		return err
	}
	_go.ID = scan_go.ID
	return nil
}

// Update returns a builder for updating this Goods.
// Note that you need to call Goods.Unwrap() before calling this method if this Goods
// was returned from a transaction, and the transaction was committed or rolled back.
func (_go *Goods) Update() *GoodsUpdateOne {
	return (&GoodsClient{config: _go.config}).UpdateOne(_go)
}

// Unwrap unwraps the Goods entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_go *Goods) Unwrap() *Goods {
	_tx, ok := _go.config.driver.(*txDriver)
	if !ok {
		panic("ent: Goods is not a transactional entity")
	}
	_go.config.driver = _tx.drv
	return _go
}

// String implements the fmt.Stringer.
func (_go *Goods) String() string {
	var builder strings.Builder
	builder.Grow(8)
	builder.WriteString("Goods(")
	builder.WriteString(fmt.Sprintf("id=%v", _go.ID))
	builder.WriteByte(')')
	return builder.String()
}

// GoodsSlice is a parsable slice of Goods.
type GoodsSlice []*Goods

// FromResponse scans the gremlin response data into GoodsSlice.
func (_go *GoodsSlice) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_go []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scan_go); err != nil {
		return err
	}
	for _, v := range scan_go {
		*_go = append(*_go, &Goods{
			ID: v.ID,
		})
	}
	return nil
}

func (_go GoodsSlice) config(cfg config) {
	for _i := range _go {
		_go[_i].config = cfg
	}
}
