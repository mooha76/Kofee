//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/Kofee/Auth_service/api"
	"github.com/mooha76/Kofee/Auth_service/api/service"
	"github.com/mooha76/Kofee/Auth_service/client"
	"github.com/mooha76/Kofee/Auth_service/config"

	//"github.com/mooha76/Kofee/Auth_service/db"

	"github.com/mooha76/Kofee/Auth_service/usecase"
)

func InitializeServices(cfg *config.Config) (*api.ServiceServer, error) {

	wire.Build(
		//db.ConnectDatabase,
		//repository.NewAuthRepository,
		//otp.NewTwilioOtpAuth,
		//token.NewJWTAuth,
		client.NewUserClient,
		usecase.NewAuthUsecase,

		client.NewPartnerClient,
		usecase.NewPartnerUsecase,

		service.NewAuthServiceServer,

		api.NewServerGRPC,
	)
	return &api.ServiceServer{}, nil
}
