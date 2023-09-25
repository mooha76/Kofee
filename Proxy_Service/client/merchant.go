package client

import (
	"context"

	"github.com/mooha76/GoGrpcProxy/client/interfaces"
	"github.com/mooha76/GoGrpcProxy/config"
	"github.com/mooha76/GoGrpcProxy/model"
	"github.com/mooha76/GoGrpcProxy/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type merchantClient struct {
	client pb.Merchan_PURCHISE_SERVICEClient
}

func NewMerchantClient(cfg *config.Config) (interfaces.MerchantClient, error) {

	gcc, err := grpc.Dial(cfg.MerchantServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewMerchan_PURCHISE_SERVICEClient(gcc)

	return &merchantClient{
		client: client,
	}, nil
}

func (c *merchantClient) SEND_MERCHAN_API_PURCHASE_REQ(ctx context.Context, req model.Request) (status string, err error) {

	res, err := c.client.SEND_MERCHAN_API_PURCHASE_REQ(ctx, &pb.Request{
		SchemaVersion: req.SchemaVersion,
		RequestId:     req.RequestId,
		Timestamp:     req.Timestamp,
		ChannelName:   req.ChannelName,
		ServiceName:   req.ServiceName,
		ServiceParams: &pb.ServiceParams{
			MerchantUid:   req.ServiceParams.MerchantUid,
			PaymentMethod: req.ServiceParams.PaymentMethod,
			ApiKey:        req.ServiceParams.ApiKey,
			ApiUserId:     req.ServiceParams.ApiUserId,
			PayerInfo: &pb.PayerInfo{
				AccountNo: req.ServiceParams.PayerInfo.AccountNo,
			},
		},
	})
	if err != nil {
		return status, err
	}
	return res.GetStatus(), nil
}
