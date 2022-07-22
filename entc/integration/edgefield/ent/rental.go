// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"github.com/google/uuid"
)

// Rental is the model entity for the Rental schema.
type Rental struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// CarID holds the value of the "car_id" field.
	CarID uuid.UUID `json:"car_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RentalQuery when eager-loading is set.
	Edges RentalEdges `json:"edges"`
}

// RentalEdges holds the relations/edges for other nodes in the graph.
type RentalEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Car holds the value of the car edge.
	Car *Car `json:"car,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RentalEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CarOrErr returns the Car value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RentalEdges) CarOrErr() (*Car, error) {
	if e.loadedTypes[1] {
		if e.Car == nil {
			// The edge car was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: car.Label}
		}
		return e.Car, nil
	}
	return nil, &NotLoadedError{edge: "car"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Rental) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case rental.FieldID, rental.FieldUserID:
			values[i] = new(sql.NullInt64)
		case rental.FieldDate:
			values[i] = new(sql.NullTime)
		case rental.FieldCarID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Rental", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Rental fields.
func (r *Rental) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rental.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case rental.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				r.Date = value.Time
			}
		case rental.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				r.UserID = int(value.Int64)
			}
		case rental.FieldCarID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field car_id", values[i])
			} else if value != nil {
				r.CarID = *value
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Rental entity.
func (r *Rental) QueryUser() *UserQuery {
	return (&RentalClient{config: r.config}).QueryUser(r)
}

// QueryCar queries the "car" edge of the Rental entity.
func (r *Rental) QueryCar() *CarQuery {
	return (&RentalClient{config: r.config}).QueryCar(r)
}

// Update returns a builder for updating this Rental.
// Note that you need to call Rental.Unwrap() before calling this method if this Rental
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Rental) Update() *RentalUpdateOne {
	return (&RentalClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Rental entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Rental) Unwrap() *Rental {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Rental is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Rental) String() string {
	var builder strings.Builder
	builder.Grow(35)
	builder.WriteString("Rental(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("date=")
	builder.WriteString(r.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("car_id=")
	builder.WriteString(fmt.Sprintf("%v", r.CarID))
	builder.WriteByte(')')
	return builder.String()
}

// Rentals is a parsable slice of Rental.
type Rentals []*Rental

func (r Rentals) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
