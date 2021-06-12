// Code generated by entc, DO NOT EDIT.

package ent

import (
	"goreact/ent/schema"
	"goreact/ent/track"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	trackFields := schema.Track{}.Fields()
	_ = trackFields
	// trackDescLiked is the schema descriptor for liked field.
	trackDescLiked := trackFields[6].Descriptor()
	// track.DefaultLiked holds the default value on creation for the liked field.
	track.DefaultLiked = trackDescLiked.Default.(bool)
}
