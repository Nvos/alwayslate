package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("timesheets", Timesheet.Type).Annotations(entgql.Bind()),
		edge.From("project", Project.Type).Ref("activities").Unique().Required().Annotations(entgql.Bind()),
	}
}
