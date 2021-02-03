package configs

import (
	"dafiti-group/microservice/controllers"
	"dafiti-group/microservice/middlewares"

	"github.com/gin-gonic/gin"
)

// GetServer returns the default application configuration
func GetServer() *gin.Engine {
	server := gin.New()

	// default middlewares section
	server.Use(middlewares.AccessMiddleware())
	server.Use(gin.Recovery())
	server.Use(middlewares.InstanaTracerMiddleware())

	// routes section
	server.GET("/", controllers.Example)
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)

	return server
}
