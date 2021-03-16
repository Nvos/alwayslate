package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alwayslate/ent"
	"alwayslate/graph/model"
	"context"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProjectInput) (*ent.Project, error) {
	return r.Client.Project.Create().
		SetName(input.Name).
		Save(ctx)
}
