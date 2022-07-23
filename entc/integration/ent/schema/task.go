// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/ent/schema/task"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("priority").
			GoType(task.Priority(0)).
			Default(int(task.PriorityMid)).
			Validate(func(i int) error {
				return task.Priority(i).Validate()
			}),
		field.JSON("priorities", map[string]task.Priority{}).
			Optional(),
	}
}
