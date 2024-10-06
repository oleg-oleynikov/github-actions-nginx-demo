package main

import (
	"github-actions-nginx-demo/config"
	"github-actions-nginx-demo/logger"
	"github-actions-nginx-demo/server"
	"log"
)

const (
	pathToCfg = "config/config.yml"
)

func main() {
	cfg, err := config.LoadConfig(pathToCfg)
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.Fatal(server.NewServer(appLogger, cfg).Run())
}
