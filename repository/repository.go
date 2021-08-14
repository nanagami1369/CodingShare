package repository

import (
	"github.com/nanagami1369/CodingShare/ent"
)

type UserAccountRepository interface {
	FindOne(id string) (selectedUser *ent.User, err error)
}
