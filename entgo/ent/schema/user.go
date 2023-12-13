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
	return []ent.Field{
		// entgo는 validator가 내장으로 포함되어 있어, NotEmtpy 메서드를 실행할 수 있다.
		field.String("eamil").
			NotEmpty(),
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
	}
}
