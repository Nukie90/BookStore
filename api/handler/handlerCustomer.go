package handler

import (
	"main/data/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddToCart(c *fiber.Ctx, db *gorm.DB) error {
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



	if userType != "Customer" {
		return c.Status(403).JSON(fiber.Map{
			"message": "You are not allowed to access this area",
		})
	}
	
	shoppingCart := new(model.ShoppingCart)

	if err := c.BodyParser(shoppingCart); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	//parse only the title and quantity
	shoppingCart.ID = uuid.New()
	user := new(model.User)
	db.Where("username = ?", claims["username"]).First(&user)
	shoppingCart.UserID = user.ID

	book := new(model.Book)
	db.Where("title = ?", shoppingCart.Title).First(&book)
	shoppingCart.BookID = book.ID
	if shoppingCart.Quantity > book.Stock {
		return c.Status(400).JSON(fiber.Map{
			"error": "Not enough stock",
		})
	}
	shoppingCart.Cost = book.Price * float64(shoppingCart.Quantity)

	db.Create(&shoppingCart)

	return c.JSON(shoppingCart)

}

func GetCart(c *fiber.Ctx, db *gorm.DB) error {
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
	if claims["user_type"] != "Customer" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	user := new(model.User)
	db.Where("username = ?", claims["username"]).First(&user)

	var shoppingCart []model.ShoppingCart
	db.Where("user_id = ?", user.ID).Find(&shoppingCart)

	return c.JSON(shoppingCart)
}

func Checkout(c *fiber.Ctx, db *gorm.DB) error {
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
	if claims["user_type"] != "Customer" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	user := new(model.User)
	db.Where("username = ?", claims["username"]).First(&user)

	var shoppingCart []model.ShoppingCart
	db.Where("user_id = ?", user.ID).Find(&shoppingCart)

	bookList := make([]model.Book, len(shoppingCart))
	cost := 0.0
	for i, item := range shoppingCart {
		book := new(model.Book)
		db.Where("id = ?", item.BookID).First(&book)
		book.Stock = book.Stock - item.Quantity
		if book.Stock < 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Not enough stock",
			})
		}
		db.Save(&book)
		bookList[i] = *book
		cost += item.Cost
	}
	record := new(model.SoldRecord)
	record.ID = uuid.New()
	record.BuyerID = user.ID
	record.BookList = bookList
	record.TotalPrice = cost
	db.Create(&record)

	db.Where("user_id = ?", user.ID).First(&shoppingCart)
	db.Delete(&shoppingCart)

	return c.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}

func RemoveItem(c *fiber.Ctx, db *gorm.DB) error {
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
	if claims["user_type"] != "Customer" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	user := new(model.User)
	db.Where("username = ?", claims["username"]).First(&user)

	title := c.Params("title")
	shoppingCart := new(model.ShoppingCart)
	db.Where("title = ? AND user_id = ?", title, user.ID).First(&shoppingCart)
	db.Delete(&shoppingCart)

	return c.JSON(fiber.Map{
		"message": "Item removed from cart",
	})

}