package module

import (
	"errors"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
)

type UserAccountModuleImpl struct {
	repository repository.UserAccountRepository
}

func NewUserAccountModule(r repository.UserAccountRepository) UserAccountModule {
	return &UserAccountModuleImpl{repository: r}
}

func (m *UserAccountModuleImpl) Login(request *model.LoginRequest) (user *ent.User, err error) {
	user, err = m.repository.FindOne(request.Id)
	if err != nil {
		return nil, err
	}
	if user.Password != request.Password {
		return nil, errors.New("password is incorrect")
	}
	return user, nil
}
