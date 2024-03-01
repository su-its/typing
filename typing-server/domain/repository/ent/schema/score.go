package schema

import (
	"entgo.io/ent"
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
		field.Float("score").Comment("スコアはaccuracyとkeystrokesの積で計算される"),
		field.Time("startedAt"),
		field.Time("endedAt"),
	}
}
