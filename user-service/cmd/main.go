package main

import (
	"context"
	"fmt"
	"github/rupeshmahanta/user-service/internal/handler"
	"github/rupeshmahanta/user-service/internal/repository"
	"github/rupeshmahanta/user-service/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Database connection
	conn, err := pgx.Connect(ctx, "postgres://rupeshmahanta:6206679616@localhost:5432/demodb")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	fmt.Println("Conected to DB successfully")
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	app := fiber.New()
	profileRepository := repository.NewProfileRepository(conn)
	ProfileService := service.NewProfileService(profileRepository)
	ProfileHandler := handler.NewProfileHandler(ProfileService)
	ProfileHandler.RegisterProfileRoutes(app)
	fmt.Println("Server started on port 3000")
	app.Listen(":3000")

}
