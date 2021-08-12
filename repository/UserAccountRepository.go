package repository

import (
	"github.com/nanagami1369/CodingShare/model"
	"gorm.io/gorm"
)

type UserAccountRepositoryImpl struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) UserAccountRepository {
	return &UserAccountRepositoryImpl{db: db}
}

func (r *UserAccountRepositoryImpl) FindOne(id string) (user *model.User, err error) {
	user = &model.User{}
	err = r.db.Where("user_id = ?", id).First(user).Error
	return user, err
}
