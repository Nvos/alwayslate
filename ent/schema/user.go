package schema

import (
	"alwayslate/ent/schema/mixin"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Immutable().
			Unique().
			Annotations(entgql.OrderField("USERNAME")),

		field.String("password").
			Sensitive(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}