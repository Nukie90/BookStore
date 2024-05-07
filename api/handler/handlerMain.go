package handler

import (
	"fmt"
	"main/data/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx, db *gorm.DB) error {
	user := new(model.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid input",
		})
	}

	var existingUser model.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if user.Password != existingUser.Password {
		return c.Status(400).JSON(fiber.Map{
			"password": user.Password,
			"db": existingUser.Password,
			"message": "Invalid password",
		})
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"user_type": existingUser.UserType,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 72),
	})

	return c.JSON(fiber.Map{
		"token": signedToken,
		"user_type": existingUser.UserType,
	})
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("jwt")

	return c.JSON(fiber.Map{
		"message": "Logout successfully",
	})
}

func AccessibleArea(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	userType, ok := claims["user_type"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "User type not found in claims",
		})
	}

	if userType != "Owner" && userType != "Customer" {
		return c.Status(403).JSON(fiber.Map{
			"message": "You are not allowed to access this area",
		})
	}

	return c.JSON(fiber.Map{
		"message": "This is an accessible area",
		"user_type":  userType,
	})
}

func RestrictedArea(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "Could not parse claims",
		})
	}

	userType, ok := claims["user_type"].(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"error": "User type not found in claims",
		})
	}

	if userType != "Owner" {
		return c.Status(403).JSON(fiber.Map{
			"message": "You are not allowed to access this area",
		})
	}

	return c.JSON(fiber.Map{
		"message": "This is a restricted area for owners only",
		"user_type": userType,
	})
}

func BrowseBooks(c *fiber.Ctx, db *gorm.DB) error {
	how := c.Params("how")
	switch how {
	case "all":
		var books []model.Book
		if err := db.Find(&books).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(books)
	case "available":
		var books []model.Book
		if err := db.Where("quantity > 0").Find(&books).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(books)
	case "by_price":
		var books []model.Book
		if err := db.Order("price desc").Find(&books).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(books)
	case "by_author":
		var books []model.Book
		if err := db.Order("author").Find(&books).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(books)
	case "by_title":
		var books []model.Book
		if err := db.Order("title").Find(&books).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(books)
	default:
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid input",
			"message": "Invalid input",
		})
	}
}