package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Int("like").
			Default(0),
		field.Int("postID"),
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

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("comments").
			Unique().
			Field("postID").Required(),
		edge.From("profile", Profile.Type).
			Ref("comments").
			Unique().
			Field("profileID").Required(),
	}
}
