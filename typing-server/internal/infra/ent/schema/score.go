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
		field.UUID("user_id", uuid.UUID{}).StorageKey("user_id").Unique(),
		field.Int("keystrokes"),
		field.Float("accuracy"),
		// TODO: is_max系は不要になったので、削除する https://github.com/su-its/typing/issues/208
		field.Bool("is_max_keystrokes").Optional().Comment("条件を満たす結果のうち、Userのkeystrokesが最大のもの"),
		field.Bool("is_max_accuracy").Optional().Comment("条件を満たす結果のうち、Userのaccuracyが最大のもの"),
		field.Time("created_at").Immutable().Default(time.Now)}
}

// Edges of the Score.
func (Score) Edges() []ent.Edge {
	return []ent.Edge{
		//ScoreとUserの関係をInverseで持たせる。ScoreはUserを必須とする
		edge.From("user", User.Type).
			Ref("scores").
			Unique().
			Field("user_id"). // ここでリンクするカラムを指定
			Required(),
	}
}
