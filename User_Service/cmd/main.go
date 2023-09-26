package main

import (
	"log"

	"github.com/mooha76/Kofee/User_Service/config"
	"github.com/mooha76/Kofee/User_Service/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	service, err := di.InitializeService(cfg)
	if err != nil {
		log.Fatalf("failed initialize service error:%s", err.Error())
	}

	service.Start()
}
