package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"vega-server/pkg/log"

	"time"
)

func LoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.MustGet("request_id")

		start := time.Now()
		c.Next()
		end := time.Now()

		logger.Info("HTTP Request",
			zap.String("request_id", requestID.(uuid.UUID).String()),
			zap.String("method", c.Request.Method),
			zap.String("url_path", c.Request.URL.Path),
			zap.Duration("duration", end.Sub(start)),
		)
	}
}
