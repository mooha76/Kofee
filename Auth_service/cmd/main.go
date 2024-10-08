package main

import (
	"log"

	"github.com/mooha76/Kofee/Auth_service/config"
	"github.com/mooha76/Kofee/Auth_service/di"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	service, err := di.InitializeServices(cfg)
	if err != nil {
		log.Fatalf("failed initialize service error:%s", err.Error())
	}

	service.Start()

}
