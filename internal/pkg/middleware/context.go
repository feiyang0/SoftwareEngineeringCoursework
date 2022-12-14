package middleware

import (
	"SoftwareEngine/internal/pkg/constant"
	"SoftwareEngine/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestID, c.GetString(constant.XRequestIDKey))
		c.Set(log.KeyUserId, c.GetString(constant.XUserIdKey))
		c.Next()
	}
}
