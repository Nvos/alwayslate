// Code generated by entc, DO NOT EDIT.

package ent

import (
	"alwayslate/ent/activity"
	"alwayslate/ent/timesheet"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Timesheet is the model entity for the Timesheet schema.
type Timesheet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt time.Time `json:"started_at,omitempty"`
	// EndedAt holds the value of the "ended_at" field.
	EndedAt time.Time `json:"ended_at,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration time.Duration `json:"duration,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TimesheetQuery when eager-loading is set.
	Edges               TimesheetEdges `json:"edges"`
	activity_timesheets *int
}

// TimesheetEdges holds the relations/edges for other nodes in the graph.
type TimesheetEdges struct {
	// Activity holds the value of the activity edge.
	Activity *Activity `json:"activity,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ActivityOrErr returns the Activity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TimesheetEdges) ActivityOrErr() (*Activity, error) {
	if e.loadedTypes[0] {
		if e.Activity == nil {
			// The edge activity was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: activity.Label}
		}
		return e.Activity, nil
	}
	return nil, &NotLoadedError{edge: "activity"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Timesheet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case timesheet.FieldID, timesheet.FieldDuration:
			values[i] = &sql.NullInt64{}
		case timesheet.FieldStartedAt, timesheet.FieldEndedAt:
			values[i] = &sql.NullTime{}
		case timesheet.ForeignKeys[0]: // activity_timesheets
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Timesheet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Timesheet fields.
func (t *Timesheet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case timesheet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case timesheet.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				t.StartedAt = value.Time
			}
		case timesheet.FieldEndedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ended_at", values[i])
			} else if value.Valid {
				t.EndedAt = value.Time
			}
		case timesheet.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				t.Duration = time.Duration(value.Int64)
			}
		case timesheet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field activity_timesheets", value)
			} else if value.Valid {
				t.activity_timesheets = new(int)
				*t.activity_timesheets = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryActivity queries the "activity" edge of the Timesheet entity.
func (t *Timesheet) QueryActivity() *ActivityQuery {
	return (&TimesheetClient{config: t.config}).QueryActivity(t)
}

// Update returns a builder for updating this Timesheet.
// Note that you need to call Timesheet.Unwrap() before calling this method if this Timesheet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Timesheet) Update() *TimesheetUpdateOne {
	return (&TimesheetClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Timesheet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Timesheet) Unwrap() *Timesheet {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Timesheet is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Timesheet) String() string {
	var builder strings.Builder
	builder.WriteString("Timesheet(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", started_at=")
	builder.WriteString(t.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ended_at=")
	builder.WriteString(t.EndedAt.Format(time.ANSIC))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", t.Duration))
	builder.WriteByte(')')
	return builder.String()
}

// Timesheets is a parsable slice of Timesheet.
type Timesheets []*Timesheet

func (t Timesheets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
