package module

import "github.com/nanagami1369/CodingShare/ent"

type UserAccountModule interface {
	Login(id string, password string) (user *ent.User, err error)
}
