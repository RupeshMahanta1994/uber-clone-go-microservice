package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rupeshmahanta/auth-service/internal/handler"
)

// SetupRouter configures and returns a Gin router with authentication routes
func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Authentication routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", func(c *gin.Context) {
			var req handler.RegisterRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			resp, err := authHandler.Register(c.Request.Context(), &req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, resp)
		})

		auth.POST("/login", func(c *gin.Context) {
			var req handler.LoginRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			resp, err := authHandler.Login(c.Request.Context(), &req)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, resp)
		})
	}

	return router
}
