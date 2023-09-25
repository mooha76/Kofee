//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/GoGrpcProxy/api"
	"github.com/mooha76/GoGrpcProxy/api/handler"
	"github.com/mooha76/GoGrpcProxy/client"
	"github.com/mooha76/GoGrpcProxy/config"
)

func InitializeApi(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		client.NewAuthClient,
		client.NewUserClient,
		client.NewMerchantClient,

		handler.NewAuthHandler,
		handler.NewUserHandler,
		handler.NewMerchantHandler,

		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
