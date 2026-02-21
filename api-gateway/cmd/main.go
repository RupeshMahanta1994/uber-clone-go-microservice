package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rupeshmahanta/api-gateway/internal/middleware"
	"github.com/rupeshmahanta/api-gateway/internal/proxy"
)

func main() {
	app := fiber.New()

	//public Routes
	app.Post("/auth/login", proxy.ReverseProxy("localhost:8080"))
	app.Post("/auth/register", proxy.ReverseProxy("localhost:8080"))

	//Protected routes group
	protected := app.Group("/api", middleware.AuthMiddleware())
	//User service routes
	protected.All("/profile/*", proxy.ReverseProxy("user-service:3000"))
	app.Listen(":8081")
}
