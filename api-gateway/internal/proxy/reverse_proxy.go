package proxy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ReverseProxy(target string) fiber.Handler {
	return proxy.Balancer(proxy.Config{
		Servers: []string{
			"http://" + target,
		},
		ModifyRequest: func(c *fiber.Ctx) error {

			// SAFE extraction
			if userID, ok := c.Locals("user_id").(string); ok {
				c.Request().Header.Set("X-User-ID", userID)
			}

			if role, ok := c.Locals("role").(string); ok {
				c.Request().Header.Set("X-User-Role", role)
			}

			return nil
		},
	})
}
