// Code generated by entc, DO NOT EDIT.

package ent

import (
	"alwayslate/ent/activity"
	"alwayslate/ent/predicate"
	"alwayslate/ent/project"
	"alwayslate/ent/timesheet"
	"alwayslate/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   activity.Table,
			Columns: activity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: activity.FieldID,
			},
		},
		Type: "Activity",
		Fields: map[string]*sqlgraph.FieldSpec{
			activity.FieldName: {Type: field.TypeString, Column: activity.FieldName},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: project.FieldID,
			},
		},
		Type: "Project",
		Fields: map[string]*sqlgraph.FieldSpec{
			project.FieldName: {Type: field.TypeString, Column: project.FieldName},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   timesheet.Table,
			Columns: timesheet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timesheet.FieldID,
			},
		},
		Type: "Timesheet",
		Fields: map[string]*sqlgraph.FieldSpec{
			timesheet.FieldStartedAt: {Type: field.TypeTime, Column: timesheet.FieldStartedAt},
			timesheet.FieldEndedAt:   {Type: field.TypeTime, Column: timesheet.FieldEndedAt},
			timesheet.FieldDuration:  {Type: field.TypeInt64, Column: timesheet.FieldDuration},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreatedAt: {Type: field.TypeTime, Column: user.FieldCreatedAt},
			user.FieldCreatedBy: {Type: field.TypeInt, Column: user.FieldCreatedBy},
			user.FieldUpdatedAt: {Type: field.TypeTime, Column: user.FieldUpdatedAt},
			user.FieldUpdatedBy: {Type: field.TypeInt, Column: user.FieldUpdatedBy},
			user.FieldUsername:  {Type: field.TypeString, Column: user.FieldUsername},
			user.FieldPassword:  {Type: field.TypeString, Column: user.FieldPassword},
		},
	}
	graph.MustAddE(
		"timesheets",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activity.TimesheetsTable,
			Columns: []string{activity.TimesheetsColumn},
			Bidi:    false,
		},
		"Activity",
		"Timesheet",
	)
	graph.MustAddE(
		"project",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activity.ProjectTable,
			Columns: []string{activity.ProjectColumn},
			Bidi:    false,
		},
		"Activity",
		"Project",
	)
	graph.MustAddE(
		"activities",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ActivitiesTable,
			Columns: []string{project.ActivitiesColumn},
			Bidi:    false,
		},
		"Project",
		"Activity",
	)
	graph.MustAddE(
		"activity",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timesheet.ActivityTable,
			Columns: []string{timesheet.ActivityColumn},
			Bidi:    false,
		},
		"Timesheet",
		"Activity",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *ActivityQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ActivityQuery builder.
func (aq *ActivityQuery) Filter() *ActivityFilter {
	return &ActivityFilter{aq}
}

// addPredicate implements the predicateAdder interface.
func (m *ActivityMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ActivityMutation builder.
func (m *ActivityMutation) Filter() *ActivityFilter {
	return &ActivityFilter{m}
}

// ActivityFilter provides a generic filtering capability at runtime for ActivityQuery.
type ActivityFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ActivityFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ActivityFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(activity.FieldID))
}

// WhereName applies the entql string predicate on the name field.
func (f *ActivityFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(activity.FieldName))
}

// WhereHasTimesheets applies a predicate to check if query has an edge timesheets.
func (f *ActivityFilter) WhereHasTimesheets() {
	f.Where(entql.HasEdge("timesheets"))
}

// WhereHasTimesheetsWith applies a predicate to check if query has an edge timesheets with a given conditions (other predicates).
func (f *ActivityFilter) WhereHasTimesheetsWith(preds ...predicate.Timesheet) {
	f.Where(entql.HasEdgeWith("timesheets", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasProject applies a predicate to check if query has an edge project.
func (f *ActivityFilter) WhereHasProject() {
	f.Where(entql.HasEdge("project"))
}

// WhereHasProjectWith applies a predicate to check if query has an edge project with a given conditions (other predicates).
func (f *ActivityFilter) WhereHasProjectWith(preds ...predicate.Project) {
	f.Where(entql.HasEdgeWith("project", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (pq *ProjectQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ProjectQuery builder.
func (pq *ProjectQuery) Filter() *ProjectFilter {
	return &ProjectFilter{pq}
}

// addPredicate implements the predicateAdder interface.
func (m *ProjectMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ProjectMutation builder.
func (m *ProjectMutation) Filter() *ProjectFilter {
	return &ProjectFilter{m}
}

// ProjectFilter provides a generic filtering capability at runtime for ProjectQuery.
type ProjectFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ProjectFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ProjectFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(project.FieldID))
}

// WhereName applies the entql string predicate on the name field.
func (f *ProjectFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(project.FieldName))
}

// WhereHasActivities applies a predicate to check if query has an edge activities.
func (f *ProjectFilter) WhereHasActivities() {
	f.Where(entql.HasEdge("activities"))
}

// WhereHasActivitiesWith applies a predicate to check if query has an edge activities with a given conditions (other predicates).
func (f *ProjectFilter) WhereHasActivitiesWith(preds ...predicate.Activity) {
	f.Where(entql.HasEdgeWith("activities", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TimesheetQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TimesheetQuery builder.
func (tq *TimesheetQuery) Filter() *TimesheetFilter {
	return &TimesheetFilter{tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TimesheetMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TimesheetMutation builder.
func (m *TimesheetMutation) Filter() *TimesheetFilter {
	return &TimesheetFilter{m}
}

// TimesheetFilter provides a generic filtering capability at runtime for TimesheetQuery.
type TimesheetFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *TimesheetFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TimesheetFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(timesheet.FieldID))
}

// WhereStartedAt applies the entql time.Time predicate on the started_at field.
func (f *TimesheetFilter) WhereStartedAt(p entql.TimeP) {
	f.Where(p.Field(timesheet.FieldStartedAt))
}

// WhereEndedAt applies the entql time.Time predicate on the ended_at field.
func (f *TimesheetFilter) WhereEndedAt(p entql.TimeP) {
	f.Where(p.Field(timesheet.FieldEndedAt))
}

// WhereDuration applies the entql int64 predicate on the duration field.
func (f *TimesheetFilter) WhereDuration(p entql.Int64P) {
	f.Where(p.Field(timesheet.FieldDuration))
}

// WhereHasActivity applies a predicate to check if query has an edge activity.
func (f *TimesheetFilter) WhereHasActivity() {
	f.Where(entql.HasEdge("activity"))
}

// WhereHasActivityWith applies a predicate to check if query has an edge activity with a given conditions (other predicates).
func (f *TimesheetFilter) WhereHasActivityWith(preds ...predicate.Activity) {
	f.Where(entql.HasEdgeWith("activity", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *UserFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *UserFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreatedAt))
}

// WhereCreatedBy applies the entql int predicate on the created_by field.
func (f *UserFilter) WhereCreatedBy(p entql.IntP) {
	f.Where(p.Field(user.FieldCreatedBy))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *UserFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdatedAt))
}

// WhereUpdatedBy applies the entql int predicate on the updated_by field.
func (f *UserFilter) WhereUpdatedBy(p entql.IntP) {
	f.Where(p.Field(user.FieldUpdatedBy))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *UserFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(user.FieldUsername))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}
