//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mooha76/Kofee/Proxy_Service/api"
	"github.com/mooha76/Kofee/Proxy_Service/api/handler"
	"github.com/mooha76/Kofee/Proxy_Service/client"
	"github.com/mooha76/Kofee/Proxy_Service/config"
)

func InitializeApi(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		client.NewAuthClient,
		client.NewUserClient,
		client.NewMerchantClient,
		client.NewPartnerClient,

		handler.NewAuthHandler,
		handler.NewUserHandler,
		handler.NewMerchantHandler,
		handler.NewPartnerHandler,

		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
