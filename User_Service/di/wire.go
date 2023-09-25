//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/GoUser_Service-Grpc/api"
	"github.com/mooha76/GoUser_Service-Grpc/api/service"
	"github.com/mooha76/GoUser_Service-Grpc/config"
	"github.com/mooha76/GoUser_Service-Grpc/db"
	"github.com/mooha76/GoUser_Service-Grpc/repository"
	"github.com/mooha76/GoUser_Service-Grpc/usecase"
)

func InitializeService(cfg *config.Config) (*api.ServiceServer, error) {

	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		service.NewUserServiceServer,
		api.NewServerGRPC,
	)
	return &api.ServiceServer{}, nil
}
