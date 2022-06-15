package configs

import (
	"github.com/dafiti-group/template-golang/controllers"
	"github.com/dafiti-group/template-golang/middlewares"

	"github.com/gin-gonic/gin"
)

// GetServer returns the default application configuration
func GetServer() *gin.Engine {
	server := gin.New()
	initAPM(server)
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
