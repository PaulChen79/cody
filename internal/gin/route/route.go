package route

import (
	"cody/config"
	websocket "cody/internal/gin/handler/websocket"
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

	api := r.Group("/v1")
	{
		websocketAPI := api.Group("/ws")
		websocketAPI.GET("/run_code", websocket.RunCode)
	}

	return r
}
