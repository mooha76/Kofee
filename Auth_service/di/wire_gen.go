// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/mooha76/Kofee/Auth_service/api"
	"github.com/mooha76/Kofee/Auth_service/api/service"
	"github.com/mooha76/Kofee/Auth_service/client"
	"github.com/mooha76/Kofee/Auth_service/config"
	"github.com/mooha76/Kofee/Auth_service/usecase"
)

// Injectors from wire.go:

func InitializeServices(cfg *config.Config) (*api.ServiceServer, error) {
	userClient, err := client.NewUserClient(cfg)
	if err != nil {
		return nil, err
	}
	authUseCase := usecase.NewAuthUsecase(userClient)
	partnerClient, err := client.NewPartnerClient(cfg)
	if err != nil {
		return nil, err
	}
	partnerUseCase := usecase.NewPartnerUsecase(partnerClient)
	authServiceServer := service.NewAuthServiceServer(authUseCase, partnerUseCase)
	serviceServer, err := api.NewServerGRPC(cfg, authServiceServer)
	if err != nil {
		return nil, err
	}
	return serviceServer, nil
}
