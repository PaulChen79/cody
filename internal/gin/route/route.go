package route

import (
	"cody/config"
	middleware "cody/internal/gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	if config.Gin.Mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.InitLogger)
	r.Use(middleware.Recovery)

	return r
}
