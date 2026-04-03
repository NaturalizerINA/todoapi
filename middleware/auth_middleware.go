package middleware

import (
	"strings"
	"todoapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "MyJwtSecretKey"

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Authorization header is required",
				Data:    fiber.Map{},
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid authorization header format",
				Data:    fiber.Map{},
			})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(models.APIResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid or expired token",
				Data:    fiber.Map{},
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["user_id"])
		c.Locals("email", claims["email"])

		return c.Next()
	}
}
