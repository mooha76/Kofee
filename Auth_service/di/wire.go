//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/Kofee/api"
	"github.com/mooha76/Kofee/api/service"
	"github.com/mooha76/Kofee/client"
	"github.com/mooha76/Kofee/config"

	//"github.com/mooha76/Kofee/db"

	"github.com/mooha76/Kofee/usecase"
)

func InitializeServices(cfg *config.Config) (*api.ServiceServer, error) {

	wire.Build(
		//db.ConnectDatabase,
		//repository.NewAuthRepository,
		//otp.NewTwilioOtpAuth,
		//token.NewJWTAuth,
		client.NewUserClient, usecase.NewAuthUsecase,
		service.NewAuthServiceServer,
		api.NewServerGRPC,
	)
	return &api.ServiceServer{}, nil
}
