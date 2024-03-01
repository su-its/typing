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
	// scoreDescID is the schema descriptor for id field.
	scoreDescID := scoreFields[0].Descriptor()
	// score.DefaultID holds the default value on creation for the id field.
	score.DefaultID = scoreDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescMailAdress is the schema descriptor for MailAdress field.
	userDescMailAdress := userFields[1].Descriptor()
	// user.MailAdressValidator is a validator for the "MailAdress" field. It is called by the builders before save.
	user.MailAdressValidator = func() func(string) error {
		validators := userDescMailAdress.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_MailAdress string) error {
			for _, fn := range fns {
				if err := fn(_MailAdress); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescHandleName is the schema descriptor for HandleName field.
	userDescHandleName := userFields[2].Descriptor()
	// user.HandleNameValidator is a validator for the "HandleName" field. It is called by the builders before save.
	user.HandleNameValidator = func() func(string) error {
		validators := userDescHandleName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_HandleName string) error {
			for _, fn := range fns {
				if err := fn(_HandleName); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescName is the schema descriptor for Name field.
	userDescName := userFields[3].Descriptor()
	// user.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_Name string) error {
			for _, fn := range fns {
				if err := fn(_Name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescHashedPassword is the schema descriptor for HashedPassword field.
	userDescHashedPassword := userFields[4].Descriptor()
	// user.HashedPasswordValidator is a validator for the "HashedPassword" field. It is called by the builders before save.
	user.HashedPasswordValidator = func() func(string) error {
		validators := userDescHashedPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_HashedPassword string) error {
			for _, fn := range fns {
				if err := fn(_HashedPassword); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
