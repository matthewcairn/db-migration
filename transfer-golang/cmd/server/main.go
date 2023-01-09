package main

import (
	"context"
	"gitlab.com/goxp/cloud0/logger"
	"os"
	"transfer/conf"
	"transfer/pkg/service"
)

func main() {
	logger.Init("transfer")
	_ = os.Setenv("PORT", conf.GetConfig().Port)
	//_ = os.Setenv("ENABLE_DB", conf.GetConfig().EnableDB)
	_ = os.Setenv("KAFKA_HOST", conf.GetConfig().KafkaHost)

	conf.LoadConfig()
	app := service.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
