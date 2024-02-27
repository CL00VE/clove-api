// Code generated by ent, DO NOT EDIT.

package ent

import (
	"clove-api/domain/community/domain/ent/comment"
	"clove-api/domain/community/domain/ent/post"
	"clove-api/domain/community/domain/ent/profile"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Like holds the value of the "like" field.
	Like int `json:"like"`
	// PostID holds the value of the "postID" field.
	PostID int `json:"postID,omitempty"`
	// ProfileID holds the value of the "profileID" field.
	ProfileID int `json:"profileID,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges        CommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[0] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[1] {
		if e.Profile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldLike, comment.FieldPostID, comment.FieldProfileID:
			values[i] = new(sql.NullInt64)
		case comment.FieldText:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt, comment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case comment.FieldLike:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like", values[i])
			} else if value.Valid {
				c.Like = int(value.Int64)
			}
		case comment.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field postID", values[i])
			} else if value.Valid {
				c.PostID = int(value.Int64)
			}
		case comment.FieldProfileID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field profileID", values[i])
			} else if value.Valid {
				c.ProfileID = int(value.Int64)
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (c *Comment) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryPost queries the "post" edge of the Comment entity.
func (c *Comment) QueryPost() *PostQuery {
	return NewCommentClient(c.config).QueryPost(c)
}

// QueryProfile queries the "profile" edge of the Comment entity.
func (c *Comment) QueryProfile() *ProfileQuery {
	return NewCommentClient(c.config).QueryProfile(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("text=")
	builder.WriteString(c.Text)
	builder.WriteString(", ")
	builder.WriteString("like=")
	builder.WriteString(fmt.Sprintf("%v", c.Like))
	builder.WriteString(", ")
	builder.WriteString("postID=")
	builder.WriteString(fmt.Sprintf("%v", c.PostID))
	builder.WriteString(", ")
	builder.WriteString("profileID=")
	builder.WriteString(fmt.Sprintf("%v", c.ProfileID))
	builder.WriteString(", ")
	builder.WriteString("createdAt=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment