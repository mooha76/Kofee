// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/mooha76/GoGrpcProxy/api"
	"github.com/mooha76/GoGrpcProxy/api/handler"
	"github.com/mooha76/GoGrpcProxy/client"
	"github.com/mooha76/GoGrpcProxy/config"
)

// Injectors from wire.go:

func InitializeApi(cfg *config.Config) (*api.Server, error) {
	authClient, err := client.NewAuthClient(cfg)
	if err != nil {
		return nil, err
	}
	authHandler := handler.NewAuthHandler(authClient)
	merchantClient, err := client.NewMerchantClient(cfg)
	if err != nil {
		return nil, err
	}
	merchantHandler := handler.NewMerchantHandler(merchantClient)
	userClient, err := client.NewUserClient(cfg)
	if err != nil {
		return nil, err
	}
	userHandler := handler.NewUserHandler(userClient)
	server := api.NewServerHTTP(cfg, authHandler, merchantHandler, userHandler)
	return server, nil
}
