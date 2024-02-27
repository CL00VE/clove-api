// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFileName holds the string denoting the filename field in the database.
	FieldFileName = "file_name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldInstance holds the string denoting the instance field in the database.
	FieldInstance = "instance"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the image in the database.
	Table = "images"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldFileName,
	FieldURL,
	FieldSize,
	FieldInstance,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// InstanceValidator is a validator for the "instance" field. It is called by the builders before save.
	InstanceValidator func([]byte) error
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Image queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFileName orders the results by the fileName field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
