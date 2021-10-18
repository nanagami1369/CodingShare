// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAccountType holds the string denoting the account_type field in the database.
	FieldAccountType = "account_type"
	// FieldStudentNumber holds the string denoting the student_number field in the database.
	FieldStudentNumber = "student_number"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// Table holds the table name of the user in the database.
	Table = "users"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "session_user"
	// VideosTable is the table that holds the videos relation/edge.
	VideosTable = "videos"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "video_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAccountType,
	FieldStudentNumber,
	FieldPassword,
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
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)

// AccountType defines the type for the "account_type" enum field.
type AccountType string

// AccountTypeGeneral is the default value of the AccountType enum.
const DefaultAccountType = AccountTypeGeneral

// AccountType values.
const (
	AccountTypeStudent AccountType = "student"
	AccountTypeTeacher AccountType = "teacher"
	AccountTypeGeneral AccountType = "general"
)

func (at AccountType) String() string {
	return string(at)
}

// AccountTypeValidator is a validator for the "account_type" field enum values. It is called by the builders before save.
func AccountTypeValidator(at AccountType) error {
	switch at {
	case AccountTypeStudent, AccountTypeTeacher, AccountTypeGeneral:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for account_type field: %q", at)
	}
}
