package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rupeshmahanta/api-gateway/internal/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "Invalid Token"})
		}
		//store value in context
		c.Locals("user_id", claims.UserId)
		c.Locals("role", claims.Role)
		return c.Next()
	}
}
