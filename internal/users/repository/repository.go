package repository

import (
	"main/data/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	userDB *gorm.DB
}

type UserDetail struct {
	entity.User
	ShoppingCart []entity.ShoppingCart
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{userDB: db}
}

func (ur *UserRepo) AddUser(user *entity.User) (err error) {
	err = ur.userDB.Create(user).Error
	return
}

func (ur *UserRepo) Login(user *entity.User) (*UserDetail, error) {
	var userDetail UserDetail
	err := ur.userDB.Where("username = ? AND password = ?", user.Username, user.Password).First(&userDetail.User).Error
	if err != nil {
		return nil, err
	}

	err = ur.userDB.Where("user_id = ?", userDetail.ID).Find(&userDetail.ShoppingCart).Error
	if err != nil {
		return nil, err
	}

	return &userDetail, nil
}