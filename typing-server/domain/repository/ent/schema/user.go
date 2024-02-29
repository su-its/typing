package schema

import (
	"entgo.io/ent"

	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	//カラムとしてvarcher型(255)のMailAdress,varcher型(255)のHandleName,varcher型(255)のName,varcher型(255)のHashedPassword,enum型のDepartmentを持たせる
	return []ent.Field{
		field.String("MailAdress").NotEmpty().MaxLen(255),
		field.String("HandleName").NotEmpty().MaxLen(255),
		field.String("Name").NotEmpty().MaxLen(255),
		field.String("HashedPassword").NotEmpty().MaxLen(255),
		field.Enum("Department").Values("CS", "BI", "IA"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//UserとScoreの関係をOne to Zero or Moreで持たせる
		edge.To("scores", Score.Type),
	}
}
