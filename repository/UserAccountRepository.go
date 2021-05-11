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

func (r *UserAccountRepositoryImpl) FindOne(email string) (user *model.User, err error) {
	user = &model.User{UserId: -1}
	err = r.db.Where("email = ?", email).First(user).Error
	return user, err
}
