// Code generated by entc, DO NOT EDIT.

package timesheet

const (
	// Label holds the string label denoting the timesheet type in the database.
	Label = "timesheet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the ended_at field in the database.
	FieldEndedAt = "ended_at"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// EdgeActivity holds the string denoting the activity edge name in mutations.
	EdgeActivity = "activity"
	// Table holds the table name of the timesheet in the database.
	Table = "timesheets"
	// ActivityTable is the table the holds the activity relation/edge.
	ActivityTable = "timesheets"
	// ActivityInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivityInverseTable = "activities"
	// ActivityColumn is the table column denoting the activity relation/edge.
	ActivityColumn = "activity_timesheets"
)

// Columns holds all SQL columns for timesheet fields.
var Columns = []string{
	FieldID,
	FieldStartedAt,
	FieldEndedAt,
	FieldDuration,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "timesheets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"activity_timesheets",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	DurationValidator func(int64) error
)