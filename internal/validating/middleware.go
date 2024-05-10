package validating

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// SetupMiddleware function is used to set up the middleware for the application
func SetupMiddleware(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey :jwtware.SigningKey{Key: []byte("secret")},
	}))
}

func IsCustomer(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.Status(400).JSON(fiber.Map{
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

	c.Locals("username", claims["username"])
	c.Locals("user_type", userType)

	return c.Next()
}

func IsOwner(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.Status(400).JSON(fiber.Map{
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

	c.Locals("username", claims["username"])
	c.Locals("user_type", userType)

	return c.Next()
}