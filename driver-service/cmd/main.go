package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/rupeshmahanta/driver-service/internal/handler"
	"github.com/rupeshmahanta/driver-service/internal/repository"
	"github.com/rupeshmahanta/driver-service/internal/service"
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
	//Injection
	driverRepository := repository.NewDriverRepository(conn)
	driverService := service.NewDriverProfile(driverRepository)
	driverHandler := handler.NewDriverHandler(driverService)

	app := fiber.New()
	driverHandler.RegisterDriverRoutes(app)
	fmt.Println("Driver Server started on port 4000")
	app.Listen(":4000")

}
