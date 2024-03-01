// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ent/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// MailAdress holds the value of the "MailAdress" field.
	MailAdress string `json:"MailAdress,omitempty"`
	// HandleName holds the value of the "HandleName" field.
	HandleName string `json:"HandleName,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// HashedPassword holds the value of the "HashedPassword" field.
	HashedPassword string `json:"HashedPassword,omitempty"`
	// Department holds the value of the "Department" field.
	Department user.Department `json:"Department,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Scores holds the value of the scores edge.
	Scores []*Score `json:"scores,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ScoresOrErr returns the Scores value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ScoresOrErr() ([]*Score, error) {
	if e.loadedTypes[0] {
		return e.Scores, nil
	}
	return nil, &NotLoadedError{edge: "scores"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldMailAdress, user.FieldHandleName, user.FieldName, user.FieldHashedPassword, user.FieldDepartment:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldMailAdress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MailAdress", values[i])
			} else if value.Valid {
				u.MailAdress = value.String
			}
		case user.FieldHandleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field HandleName", values[i])
			} else if value.Valid {
				u.HandleName = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldHashedPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field HashedPassword", values[i])
			} else if value.Valid {
				u.HashedPassword = value.String
			}
		case user.FieldDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Department", values[i])
			} else if value.Valid {
				u.Department = user.Department(value.String)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryScores queries the "scores" edge of the User entity.
func (u *User) QueryScores() *ScoreQuery {
	return NewUserClient(u.config).QueryScores(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("MailAdress=")
	builder.WriteString(u.MailAdress)
	builder.WriteString(", ")
	builder.WriteString("HandleName=")
	builder.WriteString(u.HandleName)
	builder.WriteString(", ")
	builder.WriteString("Name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("HashedPassword=")
	builder.WriteString(u.HashedPassword)
	builder.WriteString(", ")
	builder.WriteString("Department=")
	builder.WriteString(fmt.Sprintf("%v", u.Department))
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
