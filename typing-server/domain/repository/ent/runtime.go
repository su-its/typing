// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/repository/ent/schema"
	"github.com/su-its/typing/typing-server/domain/repository/ent/score"
	"github.com/su-its/typing/typing-server/domain/repository/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	scoreFields := schema.Score{}.Fields()
	_ = scoreFields
	// scoreDescCreatedAt is the schema descriptor for created_at field.
	scoreDescCreatedAt := scoreFields[6].Descriptor()
	// score.DefaultCreatedAt holds the default value on creation for the created_at field.
	score.DefaultCreatedAt = scoreDescCreatedAt.Default.(func() time.Time)
	// scoreDescID is the schema descriptor for id field.
	scoreDescID := scoreFields[0].Descriptor()
	// score.DefaultID holds the default value on creation for the id field.
	score.DefaultID = scoreDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescStudentNumber is the schema descriptor for student_number field.
	userDescStudentNumber := userFields[1].Descriptor()
	// user.StudentNumberValidator is a validator for the "student_number" field. It is called by the builders before save.
	user.StudentNumberValidator = userDescStudentNumber.Validators[0].(func(string) error)
	// userDescHandleName is the schema descriptor for handle_name field.
	userDescHandleName := userFields[2].Descriptor()
	// user.HandleNameValidator is a validator for the "handle_name" field. It is called by the builders before save.
	user.HandleNameValidator = func() func(string) error {
		validators := userDescHandleName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(handle_name string) error {
			for _, fn := range fns {
				if err := fn(handle_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
