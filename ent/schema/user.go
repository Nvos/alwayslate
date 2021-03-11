package schema

import (
	"alwayslate/ent/schema/pulid"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Immutable().Unique().
			Annotations(entgql.OrderField("USERNAME")),
		field.String("password").
			Sensitive(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pulid.MixinWithPrefix("user"),
	}
}