package repository

import (
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
)

type UserAccountRepository interface {
	FindOne(id string) (*ent.User, error)
	Create(userId, rowPassword string, accountType user.AccountType, studentNumber *int) (*ent.User, error)
	Exists(id string) (bool, error)
	ExistsStudent(studentNumber int) (bool, error)
}
