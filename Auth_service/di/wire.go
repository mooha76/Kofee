//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/GoGrpcAuth/api"
	"github.com/mooha76/GoGrpcAuth/api/service"
	"github.com/mooha76/GoGrpcAuth/client"
	"github.com/mooha76/GoGrpcAuth/config"

	//"github.com/mooha76/GoGrpcAuth/db"

	"github.com/mooha76/GoGrpcAuth/usecase"
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
