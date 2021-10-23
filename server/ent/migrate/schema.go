// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "date_of_expiry", Type: field.TypeTime},
		{Name: "session_user", Type: field.TypeInt, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_user",
				Columns:    []*schema.Column{SessionsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "account_type", Type: field.TypeEnum, Enums: []string{"student", "teacher", "general"}, Default: "general"},
		{Name: "student_number", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VideosColumns holds the columns for the "videos" table.
	VideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "language_tag", Type: field.TypeJSON},
		{Name: "upload_time", Type: field.TypeTime},
		{Name: "recording_time", Type: field.TypeInt},
		{Name: "coding_sequence", Type: field.TypeJSON},
		{Name: "comment", Type: field.TypeString},
		{Name: "is_removed", Type: field.TypeBool, Default: false},
		{Name: "video_user", Type: field.TypeInt, Nullable: true},
	}
	// VideosTable holds the schema information for the "videos" table.
	VideosTable = &schema.Table{
		Name:       "videos",
		Columns:    VideosColumns,
		PrimaryKey: []*schema.Column{VideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "videos_users_user",
				Columns:    []*schema.Column{VideosColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SessionsTable,
		UsersTable,
		VideosTable,
	}
)

func init() {
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	VideosTable.ForeignKeys[0].RefTable = UsersTable
}
