// Code generated by ent, DO NOT EDIT.

package ent

import (
	"clove-api/domain/community/domain/ent/comment"
	"clove-api/domain/community/domain/ent/post"
	"clove-api/domain/community/domain/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescLike is the schema descriptor for like field.
	commentDescLike := commentFields[1].Descriptor()
	// comment.DefaultLike holds the default value on creation for the like field.
	comment.DefaultLike = commentDescLike.Default.(int)
	// commentDescCreatedAt is the schema descriptor for createdAt field.
	commentDescCreatedAt := commentFields[4].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the createdAt field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updatedAt field.
	commentDescUpdatedAt := commentFields[5].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescLike is the schema descriptor for like field.
	postDescLike := postFields[2].Descriptor()
	// post.DefaultLike holds the default value on creation for the like field.
	post.DefaultLike = postDescLike.Default.(int)
	// postDescCreatedAt is the schema descriptor for createdAt field.
	postDescCreatedAt := postFields[4].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the createdAt field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updatedAt field.
	postDescUpdatedAt := postFields[5].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
}
