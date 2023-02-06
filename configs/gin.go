package configs

import (
	"github.com/dafiti-group/poc-golang-freight/controllers"
	"github.com/dafiti-group/poc-golang-freight/middlewares"
	"github.com/dafiti-group/poc-golang-freight/service"

	"github.com/gin-gonic/gin"
)

var (
	freightService    = service.NewFreightService()
	freightController = controllers.NewFreightController(freightService)
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
	server.POST("/", freightController.FilterShipmentType)
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)

	return server
}
