// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent/ent/score"
	"ent/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetMailAdress sets the "MailAdress" field.
func (uc *UserCreate) SetMailAdress(s string) *UserCreate {
	uc.mutation.SetMailAdress(s)
	return uc
}

// SetHandleName sets the "HandleName" field.
func (uc *UserCreate) SetHandleName(s string) *UserCreate {
	uc.mutation.SetHandleName(s)
	return uc
}

// SetName sets the "Name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetHashedPassword sets the "HashedPassword" field.
func (uc *UserCreate) SetHashedPassword(s string) *UserCreate {
	uc.mutation.SetHashedPassword(s)
	return uc
}

// SetDepartment sets the "Department" field.
func (uc *UserCreate) SetDepartment(u user.Department) *UserCreate {
	uc.mutation.SetDepartment(u)
	return uc
}

// SetCreatedAt sets the "CreatedAt" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// AddScoreIDs adds the "scores" edge to the Score entity by IDs.
func (uc *UserCreate) AddScoreIDs(ids ...int) *UserCreate {
	uc.mutation.AddScoreIDs(ids...)
	return uc
}

// AddScores adds the "scores" edges to the Score entity.
func (uc *UserCreate) AddScores(s ...*Score) *UserCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddScoreIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.MailAdress(); !ok {
		return &ValidationError{Name: "MailAdress", err: errors.New(`ent: missing required field "User.MailAdress"`)}
	}
	if v, ok := uc.mutation.MailAdress(); ok {
		if err := user.MailAdressValidator(v); err != nil {
			return &ValidationError{Name: "MailAdress", err: fmt.Errorf(`ent: validator failed for field "User.MailAdress": %w`, err)}
		}
	}
	if _, ok := uc.mutation.HandleName(); !ok {
		return &ValidationError{Name: "HandleName", err: errors.New(`ent: missing required field "User.HandleName"`)}
	}
	if v, ok := uc.mutation.HandleName(); ok {
		if err := user.HandleNameValidator(v); err != nil {
			return &ValidationError{Name: "HandleName", err: fmt.Errorf(`ent: validator failed for field "User.HandleName": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "User.Name"`)}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "User.Name": %w`, err)}
		}
	}
	if _, ok := uc.mutation.HashedPassword(); !ok {
		return &ValidationError{Name: "HashedPassword", err: errors.New(`ent: missing required field "User.HashedPassword"`)}
	}
	if v, ok := uc.mutation.HashedPassword(); ok {
		if err := user.HashedPasswordValidator(v); err != nil {
			return &ValidationError{Name: "HashedPassword", err: fmt.Errorf(`ent: validator failed for field "User.HashedPassword": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Department(); !ok {
		return &ValidationError{Name: "Department", err: errors.New(`ent: missing required field "User.Department"`)}
	}
	if v, ok := uc.mutation.Department(); ok {
		if err := user.DepartmentValidator(v); err != nil {
			return &ValidationError{Name: "Department", err: fmt.Errorf(`ent: validator failed for field "User.Department": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "User.CreatedAt"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.MailAdress(); ok {
		_spec.SetField(user.FieldMailAdress, field.TypeString, value)
		_node.MailAdress = value
	}
	if value, ok := uc.mutation.HandleName(); ok {
		_spec.SetField(user.FieldHandleName, field.TypeString, value)
		_node.HandleName = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.HashedPassword(); ok {
		_spec.SetField(user.FieldHashedPassword, field.TypeString, value)
		_node.HashedPassword = value
	}
	if value, ok := uc.mutation.Department(); ok {
		_spec.SetField(user.FieldDepartment, field.TypeEnum, value)
		_node.Department = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uc.mutation.ScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ScoresTable,
			Columns: []string{user.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(score.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
