package client

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/client/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/config"
	"github.com/mooha76/Kofee/Proxy_Service/model"
	"github.com/mooha76/Kofee/Proxy_Service/pb"
	"github.com/mooha76/Kofee/Proxy_Service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(cfg *config.Config) (interfaces.AuthClient, error) {

	gcc, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewAuthServiceClient(gcc)

	return &authClient{
		client: client,
	}, nil
}

func (c *authClient) UserSignup(ctx context.Context, req model.PartnerSingUpRequest) (ID string, err error) {

	res, err := c.client.UserSingUp(ctx, &pb.UserSignupRequest{})
	utils.LogMessage(utils.Cyan, "User SingUp Request")
	if err != nil {
		return ID, err
	}

	return res.GetUser_Id(), nil
}
