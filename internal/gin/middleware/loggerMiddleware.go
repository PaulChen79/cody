package middleware

import (
	"cody/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitLogger(c *gin.Context) {
	// init logger
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync() // flushes buffer, if any
	c.Next()
}
