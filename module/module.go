package module

import (
	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/model"
)

type UserAccountModule interface {
	Login(request *model.LoginRequest) (*ent.User, error)
	SignIn(request *model.SignInRequest) (*ent.User, error)
}
