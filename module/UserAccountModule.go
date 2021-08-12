package module

import (
	"errors"

	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
)

type UserAccountModuleImpl struct {
	repository repository.UserAccountRepository
}

func NewUserAccountRepository(r repository.UserAccountRepository) UserAccountModule {
	return &UserAccountModuleImpl{repository: r}
}

func (m *UserAccountModuleImpl) Login(id string, password string) (user *model.User, err error) {
	user, err = m.repository.FindOne(id)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password is incorrect")
	}
	return user, nil
}
