package controllers

import (
	"io/ioutil"

	"github.com/dafiti-group/${{values.component_id}}/pkg/github"
	"github.com/gin-gonic/gin"
	"github.com/gritzkoo/golang-health-checker-lw/pkg/healthchecker"
	"github.com/sirupsen/logrus"
)

var version, _ = ioutil.ReadFile("rev.txt")
var checker = healthchecker.New(healthchecker.Config{
	Name:    "${{values.component_id}}",
	Version: string(version),
	Integrations: []healthchecker.Check{
		{
			Name:   "GitHub Api Integration",
			Handle: github.Status, // you should write your own tests to pass here!
		},
	},
})

// HealthCheckLiveness show a simple check
func HealthCheckLiveness(c *gin.Context) {
	c.JSON(200, checker.Liveness())
}

// HealthCheckReadiness return a detailed status of all integrations in the list
func HealthCheckReadiness(c *gin.Context) {
	version, _ := ioutil.ReadFile("rev.txt")
	result := healthcheck.HealthCheckerDetailed(healthcheck.ApplicationConfig{
		Name:    "template-golang",
		Version: string(version),
		Integrations: []healthcheck.IntegrationConfig{
			{
				Type: healthcheck.Web,
				Name: "Just a simple test",
				Host: "https://github.com/status",
			},
		},
	})
	if !result.Status {
		// the line below will print the error on logs, so you need to
		// configure grafana or instana alarms with this context
		logrus.WithField("error", result).Warning("Health check fails")
	}
	// the pod will always return a status 200 to prevent shotdown
	c.JSON(200, result)
}
