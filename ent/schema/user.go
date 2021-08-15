package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Match(regexp.MustCompile(`^[0-9a-zA-Z]*$`)).
			Unique().
			NotEmpty(),
		field.Enum("account_type").
			Values("student", "teacher", "general").
			Default("general"),
		field.Int("student_number").
			Nillable().
			Unique().
			Optional(),
		field.String("password").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
