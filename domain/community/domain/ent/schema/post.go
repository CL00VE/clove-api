package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("icon"),
		field.String("content"),
		field.Int("like").
			Default(0),
		field.Int("profileID"),
		field.Time("createdAt").
			Default(time.Now).
			SchemaType(map[string]string{
				"mysql": "datetime",
			}),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"mysql": "datetime",
			}),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.From("profile", Profile.Type).
			Ref("posts").
			Unique().
			Field("profileID").Required(),
	}
}
