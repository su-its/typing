// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMailAdress holds the string denoting the mailadress field in the database.
	FieldMailAdress = "mail_adress"
	// FieldHandleName holds the string denoting the handlename field in the database.
	FieldHandleName = "handle_name"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHashedPassword holds the string denoting the hashedpassword field in the database.
	FieldHashedPassword = "hashed_password"
	// FieldDepartment holds the string denoting the department field in the database.
	FieldDepartment = "department"
	// EdgeScores holds the string denoting the scores edge name in mutations.
	EdgeScores = "scores"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ScoresTable is the table that holds the scores relation/edge.
	ScoresTable = "scores"
	// ScoresInverseTable is the table name for the Score entity.
	// It exists in this package in order to avoid circular dependency with the "score" package.
	ScoresInverseTable = "scores"
	// ScoresColumn is the table column denoting the scores relation/edge.
	ScoresColumn = "user_scores"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldMailAdress,
	FieldHandleName,
	FieldName,
	FieldHashedPassword,
	FieldDepartment,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// MailAdressValidator is a validator for the "MailAdress" field. It is called by the builders before save.
	MailAdressValidator func(string) error
	// HandleNameValidator is a validator for the "HandleName" field. It is called by the builders before save.
	HandleNameValidator func(string) error
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// HashedPasswordValidator is a validator for the "HashedPassword" field. It is called by the builders before save.
	HashedPasswordValidator func(string) error
)

// Department defines the type for the "Department" enum field.
type Department string

// Department values.
const (
	DepartmentCS Department = "CS"
	DepartmentBI Department = "BI"
	DepartmentIA Department = "IA"
)

func (_department Department) String() string {
	return string(_department)
}

// DepartmentValidator is a validator for the "Department" field enum values. It is called by the builders before save.
func DepartmentValidator(_department Department) error {
	switch _department {
	case DepartmentCS, DepartmentBI, DepartmentIA:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for Department field: %q", _department)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMailAdress orders the results by the MailAdress field.
func ByMailAdress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMailAdress, opts...).ToFunc()
}

// ByHandleName orders the results by the HandleName field.
func ByHandleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHandleName, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHashedPassword orders the results by the HashedPassword field.
func ByHashedPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashedPassword, opts...).ToFunc()
}

// ByDepartment orders the results by the Department field.
func ByDepartment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartment, opts...).ToFunc()
}

// ByScoresCount orders the results by scores count.
func ByScoresCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newScoresStep(), opts...)
	}
}

// ByScores orders the results by scores terms.
func ByScores(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScoresStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newScoresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScoresInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ScoresTable, ScoresColumn),
	)
}
