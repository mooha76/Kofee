package client

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/client/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/config"
	"github.com/mooha76/Kofee/Proxy_Service/model"
	"github.com/mooha76/Kofee/Proxy_Service/pb"
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

func (c *authClient) UserSignup(ctx context.Context, req model.UserSignupRequest) (otpID string, err error) {

	res, err := c.client.UserSignup(ctx, &pb.UserSignupRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  req.Password,
	})
	if err != nil {
		return otpID, err
	}
	return res.GetUser_Id(), nil
}
