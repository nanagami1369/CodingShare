// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/nanagami1369/CodingShare/ent/schema"
	"github.com/nanagami1369/CodingShare/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserID is the schema descriptor for user_id field.
	userDescUserID := userFields[0].Descriptor()
	// user.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	user.UserIDValidator = userDescUserID.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
