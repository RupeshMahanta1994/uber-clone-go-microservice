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
	app.Post("/auth/logout", proxy.ReverseProxy("localhost:8080"))

	//Protected routes group
	protected := app.Group("/", middleware.AuthMiddleware())
	//User service routes
	protected.All("/profile/*", proxy.ReverseProxy("localhost:3000"))
	protected.All("/drivers/*", proxy.ReverseProxy("localhost:4000"))

	app.Listen(":8081")
}
