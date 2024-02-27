// Code generated by ent, DO NOT EDIT.

package ent

import (
	"clove-api/domain/image/domain/ent/image"
	"clove-api/domain/image/domain/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescInstance is the schema descriptor for instance field.
	imageDescInstance := imageFields[5].Descriptor()
	// image.InstanceValidator is a validator for the "instance" field. It is called by the builders before save.
	image.InstanceValidator = imageDescInstance.Validators[0].(func([]byte) error)
	// imageDescCreatedAt is the schema descriptor for createdAt field.
	imageDescCreatedAt := imageFields[6].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the createdAt field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
}