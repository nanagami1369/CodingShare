package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/model"
)

type UserAccountRepository interface {
	FindOne(id string) (*ent.User, error)
	Create(userId, rowPassword string, accountType user.AccountType, studentNumber *int) (*ent.User, error)
	Exists(id string) (bool, error)
	ExistsStudent(studentNumber int) (bool, error)
}

type SessionRepository interface {
	Set(user *ent.User, dateOfExpiry time.Time) (*ent.Session, error)
	Get(uuid uuid.UUID) (*ent.Session, *ent.User, error)
	Remove(uuid uuid.UUID) (*ent.User, error)
	Exists(uuid uuid.UUID) (bool, error)
}

type VideoRepository interface {
	Add(user *ent.User, title string, language *model.Language, comment string, codingSequence *[]model.CodingSequence) (*model.Video, error)
	FindOne(id int) (*model.Video, error)
	FindFromTitle(title string) ([]*model.Video, error)
	FindFromUserId(id string) ([]*model.Video, error)
	Exists(id int) (bool, error)
}
