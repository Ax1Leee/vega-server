package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vega-server/api"
	"vega-server/pkg/jwt"
	"vega-server/pkg/log"

	"strings"
)

func UserAuthMiddleware(jwtService *jwt.JWTService, logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.MustGet("request_id")
		// Check if authorization header is present
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			logger.Warn("Authorization header is missing", zap.String("request_id", requestID.(string)))
			api.HandleError(c, 401, "Unauthorized", nil)
			c.Abort()
			return
		}
		// Check if authorization header format is valid
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(tokenString, bearerPrefix) {
			logger.Warn("Authorization header format is invalid", zap.String("request_id", requestID.(string)))
			api.HandleError(c, 401, "Unauthorized", nil)
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, bearerPrefix)
		// Validate token
		claims, err := jwtService.ValidateJWT(tokenString)
		if err != nil {
			logger.Error("ValidateJWT failed", zap.String("request_id", requestID.(string)), zap.Error(err))
			api.HandleError(c, 401, "Unauthorized", nil)
			c.Abort()
			return
		}
		c.Set("id", claims.ID)
		c.Next()
	}
}
