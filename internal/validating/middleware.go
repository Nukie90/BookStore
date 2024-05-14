package validating

import (
	_ "fmt"
	"strings"

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

func JwtAuth() fiber.Handler {
	return (func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user_type", claims["user_type"])
		// fmt.Println(claims["user_type"])
		c.Locals("user_id", claims["user_id"])
		// fmt.Println(claims["user_id"])
		return c.Next()
	})
}
		

func IsCustomer(c *fiber.Ctx) error {
	userType := c.Locals("user_type")
	if userType != "Customer" {
		return c.Status(403).JSON(fiber.Map{
			"message": "You are not allowed to access this area",
		})
	}

	return c.Next()
}

func IsOwner(c *fiber.Ctx) error {
	userType := c.Locals("user_type")
	if userType != "Owner" {
		return c.Status(403).JSON(fiber.Map{
			"message": "You are not allowed to access this area",
		})
	}

	return c.Next()
}