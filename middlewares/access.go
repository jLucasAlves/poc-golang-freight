package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
			"method":   c.Request.Method,
			"path":     c.Request.RequestURI,
			"status":   c.Writer.Status(),
			"referrer": c.Request.Referer(),
			// TODO: inject new request id interface
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("AccessMiddleware")
		}
	}
}
