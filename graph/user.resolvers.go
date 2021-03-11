package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alwayslate/ent"
	"alwayslate/graph/model"
	"context"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	return r.Client.User.Create().
		SetUsername(input.Username).
		SetPassword(input.Password).
		Save(ctx)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.Client.User.Query().
		Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
}
