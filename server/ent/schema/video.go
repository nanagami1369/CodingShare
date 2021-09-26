package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/nanagami1369/CodingShare/model"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.JSON("language_tag", &model.Language{}),
		field.Time("upload_time").
			Immutable().
			Default(time.Now),
		field.Int("recording_time").
			NonNegative(),
		field.JSON("coding_sequence", &[]model.CodingSequence{}),
		field.String("comment"),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required(),
	}
}
