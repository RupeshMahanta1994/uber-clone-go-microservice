package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/rupeshmahanta/auth-service/internal/handler"
	"github.com/rupeshmahanta/auth-service/internal/repository"
	"github.com/rupeshmahanta/auth-service/internal/router"
	"github.com/rupeshmahanta/auth-service/internal/service"
)

func main() {
	// Initialize context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Database connection
	conn, err := pgx.Connect(ctx, "postgres://rupeshmahanta:6206679616@localhost:5432/demodb")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Verify database connection
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("Error pinging database:", err)
	}
	log.Println("✓ Connected to database")
	// Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Verify Redis connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	log.Println("✓ Connected to Redis")

	// Initialize repository
	repo := repository.NewUserRepository(conn)
	tokenRepo := repository.NewTokenRepository(rdb)

	// Initialize service
	authService := service.NewAuthService(repo, tokenRepo)

	// Initialize handler
	authHandler := handler.NewAuthHandler(authService)

	//start the server
	ginRouter := router.SetupRouter(authHandler)

	port := ":8080"
	log.Printf("Starting auth service on port %s\n", port)
	if err := ginRouter.Run(port); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
