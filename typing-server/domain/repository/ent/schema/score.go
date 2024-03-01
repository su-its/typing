// user.goと同じように、score.goも作成します。リレーションはuser.go→score.goでone to zero or oneの関係です。
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
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID { return uuid.Must(uuid.NewRandom()) }),
		field.Int("keystrokes"),
		field.Float("accuracy"),
		field.Float("score"),
		field.Time("startedAt"),
		field.Time("endedAt"),
	}
}
