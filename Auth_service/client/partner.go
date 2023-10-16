package client

import (
	"context"
	"fmt"

	interfaces "github.com/mooha76/Kofee/Auth_service/client/interface"
	"github.com/mooha76/Kofee/Auth_service/config"
	model "github.com/mooha76/Kofee/Auth_service/model"
	"github.com/mooha76/Kofee/Auth_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type partnerClient struct {
	client pb.UserServiceClient
}

func NewPartnerClient(cfg *config.Config) (interfaces.PartnerClient, error) {

	gcc, err := grpc.Dial(cfg.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(gcc)

	return &partnerClient{
		client: client,
	}, nil
}

func (c *partnerClient) FindPartnerByEmail(ctx context.Context, email string) (model.Partner, error) {

	res, err := c.client.FindPartnerByEmail(ctx, &pb.FindPartnerEmailRequest{
		EMAIL: email,
	})

	if err != nil {
		return model.Partner{
			PARTNERID:  res.GetPARTNERID(),
			FIRSTNAME:  res.GetFIRSTNAME(),
			MIDDLENAME: res.GetMIDDLENAME(),
			LASTNAME:   res.GetLASTNAME(),
			MOBILE:     res.GetMOBILE(),
		}, err
	}

	return model.Partner{}, nil
}

func (c *partnerClient) FindPartnerByPhone(ctx context.Context, phone string) (model.Partner, error) {

	res, err := c.client.FindPartnerByPhone(ctx, &pb.FindPartnerByPhoneRequest{
		MOBILE: phone,
	})
	if err != nil {
		return model.Partner{
			PARTNERID:  res.GetPARTNERID(),
			FIRSTNAME:  res.GetFIRSTNAME(),
			MIDDLENAME: res.GetMIDDLENAME(),
			LASTNAME:   res.GetLASTNAME(),
			MOBILE:     res.GetMOBILE(),
		}, err
	}

	return model.Partner{}, nil
}

func (c *partnerClient) SavePartner(ctx context.Context, partner model.PartnerSingUpRequest) (userID uint64, err error) {
	res, err := c.client.SavePartner(ctx, &pb.SavePartnerRequest{
		FIRSTNAME:               partner.FIRSTNAME,
		MIDDLENAME:              partner.MIDDLENAME,
		LASTNAME:                partner.LASTNAME,
		MOBILE:                  partner.MOBILE,
		ACCOUNTNO:               partner.ACCOUNTNO,
		WITHDRAWALREQUESTTYPEID: partner.WITHDRAWALREQUESTTYPEID,
	})
	if err != nil {
		return 0, err
	}
	fmt.Println("user client res", res)
	return res.GetPARTNERID(), nil
}
