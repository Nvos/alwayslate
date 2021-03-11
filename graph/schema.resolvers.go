package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alwayslate/ent"
	"alwayslate/ent/schema/pulid"
	"alwayslate/graph/generated"
	"context"
)

func (r *queryResolver) Node(ctx context.Context, id pulid.ID) (ent.Noder, error) {
	return r.Client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []pulid.ID) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
