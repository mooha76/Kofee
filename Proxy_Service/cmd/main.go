package main

import (
	"log"

	"github.com/mooha76/GoGrpcProxy/config"
	"github.com/mooha76/GoGrpcProxy/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	service, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatalf("failed initialize api error:%s", err.Error())
	}

	service.Start()
}
