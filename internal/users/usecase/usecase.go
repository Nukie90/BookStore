package usecase

import (
	"fmt"
	"log"
	"main/data/entity"
	"main/internal/users/repository"
	"strconv"
	"time"



	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo *repository.UserRepo
}

func NewUserUsecase(repo *repository.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (uu *UserUsecase) AddUser(user *entity.User) (err error) {
	user.ID = uuid.New()

	switch user.UserType {
		case strconv.Itoa(1):
			user.UserType = "Customer"
		case strconv.Itoa(2):
			user.UserType = "Owner"
		default:
			return fmt.Errorf("invalid user type")
	}

	if err := uu.userRepo.AddUser(user); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uu *UserUsecase) Login(user *entity.User) (string, error) {
	userDetail, err := uu.userRepo.Login(user)
	if err != nil {
		return "", fmt.Errorf("could not login")
	}

	claims := jwt.MapClaims{
		"user_id": userDetail.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
		"user_type": userDetail.UserType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("could not login")
	}

	return signedToken, nil
}

func (uu *UserUsecase) AddToCart(shoppingCart *entity.ShoppingCart) (error) {
	shoppingCart.ID = uuid.New()
	book, found := uu.userRepo.GetBookByTitle(shoppingCart.Title)
	if !found {
		return fmt.Errorf("book not found")
	}
	shoppingCart.BookID = book.ID

	if shoppingCart.Quantity > book.Stock {
		return fmt.Errorf("not enough stock")
	}

	shoppingCart.Cost = book.Price * float64(shoppingCart.Quantity)

	if err := uu.userRepo.AddToCart(shoppingCart); err != nil {
		return fmt.Errorf("could not add to cart")
	}

	return nil
}

func (uu *UserUsecase) GetCart(userID string) ([]entity.ShoppingCart, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}

	shoppingCart, err := uu.userRepo.GetCart(parsedUserID)
	if err != nil {
		return nil, fmt.Errorf("could not get cart")
	}

	return shoppingCart, nil
}

func (uu *UserUsecase) RemoveFromCart(shoppingCart *entity.ShoppingCart) error {
	userID := shoppingCart.UserID
	title := shoppingCart.Title

	if err := uu.userRepo.RemoveFromCart(userID, title); err != nil {
		return fmt.Errorf("could not remove from cart")
	}

	return nil
}

func (uu *UserUsecase) CheckDailySellRecord() ([]entity.SoldRecord, error) {
	records, err := uu.userRepo.CheckDailySellRecord()
	if err != nil {
		return nil, fmt.Errorf("could not check daily sell record")
	}

	return records, nil
}