package module

import (
	"errors"
	"regexp"

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

func (m *UserAccountModuleImpl) Confirm(request *model.LoginRequest) (*ent.User, error) {
	user, err := m.repository.FindOne(request.Id)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.RowPassword)); err != nil {
		return nil, errors.New("password is incorrect")
	}
	return user, nil
}

func (m *UserAccountModuleImpl) SignIn(request *model.SignInRequest) (*ent.User, error) {
	if err := checkSignInRequestValidation(request); err != nil {
		return nil, err
	}
	isExists, err := m.repository.Exists(request.Id)
	if err != nil {
		return nil, err
	}
	if isExists {
		return nil, errors.New("sign in request error request id is Exists")
	}
	if request.AccountType == user.AccountTypeStudent {
		isExistsStudent, err := m.repository.ExistsStudent(*request.StudentNumber)
		if err != nil {
			return nil, err
		}
		if isExistsStudent {
			return nil, errors.New("sign in request error request student number is Exists")
		}
	} else {
		if request.StudentNumber != nil {
			return nil, errors.New("sign in request error request has student number,but request is not student")
		}
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
	r := regexp.MustCompile(`^[0-9a-zA-Z]*$`)
	if !r.MatchString(request.Id) {
		return errors.New("sign in request error only half-width alphanumeric characters can be used in the user id")
	}
	return nil
}