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

type partnerClient struct {
	client pb.AuthServiceClient
}

func NewPartnerClient(cfg *config.Config) (interfaces.PartnerClient, error) {

	gcc, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewAuthServiceClient(gcc)

	return &partnerClient{
		client: client,
	}, nil
}

func (c *partnerClient) PartnerSingUp(ctx context.Context, req model.PartnerSingUpRequest) (Id uint64, err error) {
	utils.LogMessage(utils.Cyan, "PartnerSingUp Request")
	res, err := c.client.PartnerSingUp(ctx, &pb.PartnerSingUpRequest{
		FIRSTNAME:               req.FIRSTNAME,
		MIDDLENAME:              req.MIDDLENAME,
		LASTNAME:                req.FIRSTNAME,
		MOBILE:                  req.MOBILE,
		ACCOUNTNO:               req.ACCOUNTNO,
		WITHDRAWALREQUESTTYPEID: req.WITHDRAWALREQUESTTYPEID,
		AGE:                     req.AGE,
		GENDER:                  req.GENDER,
		EMAIL:                   req.EMAIL,
		USERNAME:                req.EMAIL,
		PASSWORD:                req.PASSWORD,
	})

	println("Log trace =================")
	if err != nil {
		return Id, err

	}
	return res.GetPARTNERID(), nil
}
