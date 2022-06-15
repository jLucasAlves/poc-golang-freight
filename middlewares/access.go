package middlewares

import (
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var version, _ = ioutil.ReadFile("rev.txt")

// AccessMiddleware is responsable to use logrus in gin log interface requests
func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := time.Since(start)

		entry := logrus.WithFields(logrus.Fields{
			"duration": duration.Milliseconds(),
			"version":  string(version),
			"method":   c.Request.Method,
			"path":     c.Request.RequestURI,
			"status":   c.Writer.Status(),

			// TODO: inject new request id interface
			"requestId":     c.Writer.Header().Get("Request-Id"),
			"correlationId": c.Writer.Header().Get("Correlation-Id"),
			"headers":       c.Writer.Header(),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("AccessMiddleware")
		}
	}
}
