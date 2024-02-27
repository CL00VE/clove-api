package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("description"),
		field.String("fileName").
			Unique(),
		field.String("url").
			Unique(),
		field.String("size"),
		field.Bytes("instance").
			MaxLen(1000000), // 1 mb
		field.Time("createdAt").
			Default(time.Now).
			SchemaType(map[string]string{
				"mysql": "datetime",
			}),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
