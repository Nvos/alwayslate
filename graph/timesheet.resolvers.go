package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alwayslate/ent"
	"alwayslate/graph/generated"
	"alwayslate/graph/model"
	"context"
	"time"
)

func (r *mutationResolver) CreateTimesheet(ctx context.Context, input model.CreateTimesheetInput) (*ent.Timesheet, error) {
	return r.Client.Timesheet.Create().
		SetActivityID(input.ActivityID).
		SetDuration(time.Duration(input.Duration)).
		SetNillableStartedAt(input.StartedAt).
		SetNillableEndedAt(input.EndedAt).
		Save(ctx)
}

func (r *timesheetResolver) Duration(ctx context.Context, obj *ent.Timesheet) (int, error) {
	return int(obj.Duration), nil
}

// Timesheet returns generated.TimesheetResolver implementation.
func (r *Resolver) Timesheet() generated.TimesheetResolver { return &timesheetResolver{r} }

type timesheetResolver struct{ *Resolver }
