package middlewares

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAccessMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccessMiddleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
