package main

import (
	"cody/internal/gin/route"
	"cody/internal/logger"
	"cody/internal/provider"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	config := provider.NewConfig()

	if config.Env != "local" && config.Env != "" {
		BackgroundWorker()
	}

	router := route.SetupRouter(config)
	router.Run(config.Gin.Host + ":" + config.Gin.Port)
}

func BackgroundWorker() {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// CronJob()
}

func CronJob() {
	go func() {
		nyc, _ := time.LoadLocation("UTC")
		cronjob := cron.New(cron.WithSeconds(), cron.WithLocation(nyc))

		cronjob.AddFunc("0 0 20 * * *", func() {
		})

		config := provider.NewConfig()
		if config.Env == "develop" || config.Env == "production" {
			cronjob.Start()
		}
	}()
}
