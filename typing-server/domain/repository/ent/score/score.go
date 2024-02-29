// Code generated by ent, DO NOT EDIT.

package score

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the score type in the database.
	Label = "score"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKeystrokes holds the string denoting the keystrokes field in the database.
	FieldKeystrokes = "keystrokes"
	// FieldAccuracy holds the string denoting the accuracy field in the database.
	FieldAccuracy = "accuracy"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldStartedAt holds the string denoting the startedat field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the endedat field in the database.
	FieldEndedAt = "ended_at"
	// Table holds the table name of the score in the database.
	Table = "scores"
)

// Columns holds all SQL columns for score fields.
var Columns = []string{
	FieldID,
	FieldKeystrokes,
	FieldAccuracy,
	FieldScore,
	FieldStartedAt,
	FieldEndedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "scores"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_scores",
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

// OrderOption defines the ordering options for the Score queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKeystrokes orders the results by the keystrokes field.
func ByKeystrokes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeystrokes, opts...).ToFunc()
}

// ByAccuracy orders the results by the accuracy field.
func ByAccuracy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccuracy, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByStartedAt orders the results by the startedAt field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByEndedAt orders the results by the endedAt field.
func ByEndedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndedAt, opts...).ToFunc()
}
