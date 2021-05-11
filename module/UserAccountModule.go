package module

import (
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
)

type UserAccountModuleImpl struct {
	repository repository.UserAccountRepository
}

func NewUserAccountRepository(r repository.UserAccountRepository) UserAccountModule {
	return &UserAccountModuleImpl{repository: r}
}

func (m *UserAccountModuleImpl) Login(email string) (user *model.User, err error) {
	return m.repository.FindOne(email)
}
