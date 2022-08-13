// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/gremlin"
)

// License is the model entity for the License schema.
type License struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
}

// FromResponse scans the gremlin response data into License.
func (l *License) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanl struct {
		ID         int   `json:"id,omitempty"`
		CreateTime int64 `json:"create_time,omitempty"`
		UpdateTime int64 `json:"update_time,omitempty"`
	}
	if err := vmap.Decode(&scanl); err != nil {
		return err
	}
	l.ID = scanl.ID
	l.CreateTime = time.Unix(0, scanl.CreateTime)
	l.UpdateTime = time.Unix(0, scanl.UpdateTime)
	return nil
}

// Update returns a builder for updating this License.
// Note that you need to call License.Unwrap() before calling this method if this License
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *License) Update() *LicenseUpdateOne {
	return (&LicenseClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the License entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *License) Unwrap() *License {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: License is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *License) String() string {
	var builder strings.Builder
	builder.WriteString("License(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("create_time=")
	builder.WriteString(l.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(l.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Licenses is a parsable slice of License.
type Licenses []*License

// FromResponse scans the gremlin response data into Licenses.
func (l *Licenses) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanl []struct {
		ID         int   `json:"id,omitempty"`
		CreateTime int64 `json:"create_time,omitempty"`
		UpdateTime int64 `json:"update_time,omitempty"`
	}
	if err := vmap.Decode(&scanl); err != nil {
		return err
	}
	for _, v := range scanl {
		*l = append(*l, &License{
			ID:         v.ID,
			CreateTime: time.Unix(0, v.CreateTime),
			UpdateTime: time.Unix(0, v.UpdateTime),
		})
	}
	return nil
}

func (l Licenses) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}