package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
)

type UserAccountRepository interface {
	FindOne(id string) (*ent.User, error)
	Create(userId, rowPassword string, accountType user.AccountType, studentNumber *int) (*ent.User, error)
	Exists(id string) (bool, error)
	ExistsStudent(studentNumber int) (bool, error)
}

type SessionRepository interface {
	Set(user *ent.User, dateOfExpiry time.Time) (*ent.Session, error)
	Get(uuid uuid.UUID) (*ent.Session, error)
	Remove(uuid uuid.UUID) error
	Exists(uuid uuid.UUID) (bool,error)
}
