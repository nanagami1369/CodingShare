package module

import "github.com/nanagami1369/CodingShare/model"

type UserAccountModule interface {
	Login(email string) (user *model.User, err error)
}
