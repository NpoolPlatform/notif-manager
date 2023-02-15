// Code generated by ent, DO NOT EDIT.

package frontendtemplate

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the frontendtemplate type in the database.
	Label = "frontend_template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldLangID holds the string denoting the lang_id field in the database.
	FieldLangID = "lang_id"
	// FieldUsedFor holds the string denoting the used_for field in the database.
	FieldUsedFor = "used_for"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldSender holds the string denoting the sender field in the database.
	FieldSender = "sender"
	// Table holds the table name of the frontendtemplate in the database.
	Table = "frontend_templates"
)

// Columns holds all SQL columns for frontendtemplate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldLangID,
	FieldUsedFor,
	FieldTitle,
	FieldContent,
	FieldSender,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/notif-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultUsedFor holds the default value on creation for the "used_for" field.
	DefaultUsedFor string
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultSender holds the default value on creation for the "sender" field.
	DefaultSender string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
