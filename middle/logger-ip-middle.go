package middle

import (
	"common-web-framework/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware(context *gin.Context) {
	defer func() {
		start := time.Now()

		context.Next()

		end := time.Now()

		latency := end.Sub(start)

		var ms = fmt.Sprintf("%.2f/ms", float64(latency.Milliseconds()))

		config.LOGGER.Info("Request",
			zap.String("path", context.Request.URL.Path),
			zap.String("method", context.Request.Method),
			zap.String("user_agent", context.Request.UserAgent()),
			zap.String("latency", ms),
		)
	}()
}
