package repository

import (
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
)

type UserAccountRepository interface {
	FindOne(id string) (selectedUser *ent.User, err error)
	Create(userId, rowPassword string, accountType user.AccountType, studentNumber *int) (user *ent.User, err error)
	Exists(id string) (bool, error)
}
