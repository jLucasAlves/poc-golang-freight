package configs

import (
	"os"

	"github.com/gin-gonic/gin"
	instana "github.com/instana/go-sensor"
	"github.com/instana/go-sensor/instrumentation/instagin"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func initAPM(router *gin.Engine) {
	if os.Getenv("ENVIRONMENT") != "development" {
		tracer := instana.NewTracerWithOptions(instana.DefaultOptions())
		sensor := instana.NewSensorWithTracer(tracer)

		opentracing.InitGlobalTracer(tracer)
		instana.SetLogger(logrus.StandardLogger())

		instagin.AddMiddleware(sensor, router)
	}
}
