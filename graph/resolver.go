package graph

//go:generate go install github.com/99designs/gqlgen/cmd
//go:generate go run github.com/99designs/gqlgen

import "alwayslate/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Client *ent.Client
}
