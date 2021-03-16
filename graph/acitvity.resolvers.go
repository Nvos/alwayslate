package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alwayslate/ent"
	"alwayslate/graph/model"
	"context"
)

func (r *mutationResolver) CreateActivity(ctx context.Context, input model.CreateActivityInput) (*ent.Activity, error) {
	return r.Client.Activity.Create().
		SetName(input.Name).
		SetProjectID(input.ProjectID).
		Save(ctx)
}

func (r *queryResolver) Activities(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ActivityOrder) (*ent.ActivityConnection, error) {
	return r.Client.Debug().Activity.Query().
		Paginate(ctx, after, first, before, last, ent.WithActivityOrder(orderBy))
}
