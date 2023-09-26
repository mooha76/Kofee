//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/Kofee/User_Service/api"
	"github.com/mooha76/Kofee/User_Service/api/service"
	"github.com/mooha76/Kofee/User_Service/config"
	"github.com/mooha76/Kofee/User_Service/db"
	"github.com/mooha76/Kofee/User_Service/repository"
	"github.com/mooha76/Kofee/User_Service/usecase"
)

func InitializeService(cfg *config.Config) (*api.ServiceServer, error) {

	wire.Build(
		db.InitializeSQLXDatabase,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		service.NewUserServiceServer,
		api.NewServerGRPC,
	)
	return &api.ServiceServer{}, nil
}
