package repository

import "github.com/nanagami1369/CodingShare/model"

type UserAccountRepository interface {
	FindOne(id string) (user *model.User, err error)
}
