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
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"user_type": userDetail.UserType,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("could not login")
	}

	return signedToken, nil
}