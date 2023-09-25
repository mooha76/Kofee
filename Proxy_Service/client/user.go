package client

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/client/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/config"
	"github.com/mooha76/Kofee/Proxy_Service/pb"
	"github.com/mooha76/Kofee/Proxy_Service/utils/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	client pb.UserServiceClient
}

func NewUserClient(cfg *config.Config) (interfaces.UserClient, error) {
	gcc, err := grpc.Dial(cfg.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewUserServiceClient(gcc)

	return &userClient{
		client: client,
	}, nil
}

func (c *userClient) GetUserProfile(ctx context.Context, userID uint64) (response.User, error) {

	res, err := c.client.GetUserProfile(ctx, &pb.GetUserProfileRequest{UserId: userID})
	if err != nil {
		return response.User{}, err
	}

	return response.User{
		FirstName: res.GetFirstName(),
		LastName:  res.GetLastName(),
		Age:       res.GetAge(),
		Email:     res.GetEmail(),
		Phone:     res.GetPhone(),
	}, nil
}
