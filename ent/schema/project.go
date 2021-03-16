package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", Activity.Type).Annotations(entgql.Bind()),
	}
}
