package configs

import (
	"github.com/dafiti-group/golang-template-project/controllers"
	"github.com/dafiti-group/golang-template-project/middlewares"

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
