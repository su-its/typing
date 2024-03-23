package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Score struct {
	ent.Schema
}

func (Score) Fields() []ent.Field {
	//カラムとしてint型のkeystrokes,float型のaccuracy，float型のscore,datetime型のstartedAt,datetime型のendedAtを持たせる
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("id").Unique(),
		field.Int("keystrokes"),
		field.Float("accuracy"),
		field.Time("created_at").Immutable().Default(time.Now)}
}

// Edges of the Score.
func (Score) Edges() []ent.Edge {
	return []ent.Edge{
		//ScoreとUserの関係をInverseで持たせる。ScoreはUserを必須とする
		edge.From("user", User.Type).
			Ref("scores").
			Unique().
			Required(),
	}
}
