package repository

import (
	"fmt"
	"main/data/entity"
	"time"

	"github.com/google/uuid"
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

func (ur *UserRepo) GetUserByID(id uuid.UUID) (user *entity.User, found bool){
	fmt.Println(id)
	err := ur.userDB.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, false
	}
	return user, true
}

func (ur *UserRepo) GetBookByTitle(title string) (book *entity.Book, found bool) {
	book = new(entity.Book)
	err := ur.userDB.Where("title = ?", title).First(book).Error
	if err != nil {
		return nil, false
	}
	return book, true
}

func (ur *UserRepo) AddToCart(shoppingCart *entity.ShoppingCart) (err error) {
	err = ur.userDB.Create(shoppingCart).Error
	return
}

func (ur *UserRepo) GetCart(userID uuid.UUID) (shoppingCart []entity.ShoppingCart, err error) {
	err = ur.userDB.Where("user_id = ?", userID).Find(&shoppingCart).Error
	return
}

func (ur *UserRepo) RemoveFromCart(id uuid.UUID, title string) (err error) {
	err = ur.userDB.Where("user_id = ? AND title = ?", id, title).Delete(&entity.ShoppingCart{}).Error
	return
}

func (ur *UserRepo) CheckDailySellRecord() (records []entity.SoldRecord, err error) {
	// query for cuurent date
	today := time.Now().Format("2006-01-02")
	err = ur.userDB.Where("created_at::date = ?", today).Find(&records).Error
	return
}