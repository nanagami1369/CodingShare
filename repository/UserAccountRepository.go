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

func (r *UserAccountRepositoryImpl) FindOne(id string) (selectedUser *ent.User, err error) {
	user, err := r.client.User.Query().Where(user.UserID(id)).Only(r.context)
	if err != nil {
		return nil, err
	}
	return user, err
}
