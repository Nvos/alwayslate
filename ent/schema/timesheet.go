package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Timesheet holds the schema definition for the Timesheet entity.
type Timesheet struct {
	ent.Schema
}

// Fields of the Timesheet.
func (Timesheet) Fields() []ent.Field {
	return []ent.Field{
		field.Time("started_at").Optional(),
		field.Time("ended_at").Optional(),
		field.Int64("duration").
			GoType(time.Duration(0)).
			NonNegative(),
	}
}

// Edges of the Timesheet.
func (Timesheet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activity", Activity.Type).Ref("timesheets").Unique().Required().Annotations(entgql.Bind()),
	}
}
