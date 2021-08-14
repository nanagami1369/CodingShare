package module

import (
	"errors"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/model"
	"github.com/nanagami1369/CodingShare/repository"
	"golang.org/x/crypto/bcrypt"
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

func (m *UserAccountModuleImpl) SignIn(request *model.SignInRequest) (user *ent.User, err error) {
	if err := checkSignInRequestValidation(request); err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(request.RowPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	password := string(hash)
	return m.repository.Create(request.Id, password, request.AccountType, request.StudentNumber)
}

// サインイン用のバリデーション
func checkSignInRequestValidation(request *model.SignInRequest) error {
	if request.Id == "" {
		return errors.New("sign in request error id is empty")
	}
	if len(request.RowPassword) < 8 {
		return errors.New("sign in request error password is 8 characters or more")
	}
	if request.AccountType == "" {
		return errors.New("sign in request error accountType is empty")
	}
	if request.AccountType == user.AccountTypeStudent && request.StudentNumber == nil {
		return errors.New("sign in request error request is student ,but request have not a student Number")
	}
	return nil
}
