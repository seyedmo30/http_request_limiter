package main

import (
	"github.com/seyedmo30/http_request_limiter/internal/config"
	"github.com/seyedmo30/http_request_limiter/internal/controller"
	"github.com/seyedmo30/http_request_limiter/internal/repository"
	"github.com/seyedmo30/http_request_limiter/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize dependencies
	repo := repository.NewLimiterRepository()
	limiterService := service.NewLimiterService(repo, *cfg)
	limiterController := controller.NewLimiterController(limiterService)

	// Setup Gin
	router := gin.Default()
	router.GET("/request", limiterController.HandleRequest)

	// Start the server
	router.Run(":8088")
}
