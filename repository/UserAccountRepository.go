package repository

import (
	"context"

	"github.com/nanagami1369/CodingShare/ent"
	"github.com/nanagami1369/CodingShare/ent/user"
)

type UserAccountRepositoryImpl struct {
	context context.Context
	client  *ent.Client
}

func NewUserAccountRepository(context context.Context, client *ent.Client) UserAccountRepository {
	return &UserAccountRepositoryImpl{context: context, client: client}
}

func (r *UserAccountRepositoryImpl) FindOne(id string) (*ent.User, error) {
	user, err := r.client.User.Query().Where(user.UserID(id)).Only(r.context)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserAccountRepositoryImpl) Create(userId, rowPassword string, accountType user.AccountType, studentNumber *int) (*ent.User, error) {
	if studentNumber == nil {
		return r.client.User.Create().
			SetUserID(userId).
			SetPassword(rowPassword).
			SetAccountType(accountType).
			Save(r.context)
	} else {
		return r.client.User.Create().
			SetUserID(userId).
			SetPassword(rowPassword).
			SetAccountType(accountType).
			SetStudentNumber(*studentNumber).
			Save(r.context)
	}
}

func (r *UserAccountRepositoryImpl) Exists(id string) (bool, error) {
	return r.client.User.Query().Where(user.UserID(id)).Exist(r.context)
}

func (r *UserAccountRepositoryImpl) ExistsStudent(studentNumber int) (bool, error) {
	return r.client.User.Query().Where(user.StudentNumber(studentNumber)).Exist(r.context)
}
