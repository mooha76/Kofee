package service

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
	"github.com/mooha76/Kofee/User_Service/pb"
	usecase "github.com/mooha76/Kofee/User_Service/usecase/interfaces"
	"github.com/mooha76/Kofee/User_Service/utils"

	//"github.com/mooha76/Kofee/User_Service/utils/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PartnerServiceServer struct {
	usecase usecase.PartnerUsecase
	pb.UnimplementedUserServiceServer
}

func NewPartnerServiceServer(usecase usecase.PartnerUsecase) pb.UserServiceServer {
	return &PartnerServiceServer{
		usecase: usecase,
	}
}

func (c *PartnerServiceServer) SavePartner(ctx context.Context, req *pb.SavePartnerRequest) (*pb.SavePartnerResponse, error) {
	utils.LogMessage(utils.Cyan, "SavePartner Request Invoked")
	PARTNERID, err := c.usecase.SavePartner(context.Background(), model.PARTNERSSINGUP_REQUEST{
		FIRSTNAME:               req.GetFIRSTNAME(),
		MIDDLENAME:              req.GetMIDDLENAME(),
		LASTNAME:                req.GetLASTNAME(),
		MOBILE:                  req.GetMOBILE(),
		ACCOUNTNO:               req.GetACCOUNTNO(),
		WITHDRAWALREQUESTTYPEID: req.GetWITHDRAWALREQUESTTYPEID(),
		AGE:                     req.GetAGE(),
		GENDER:                  req.GetGENDER(),
		EMAIL:                   req.GetEMAIL(),
		PASSWORD:                req.GetPASSWORD(),
	})

	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully User Details Saved")
	return &pb.SavePartnerResponse{PARTNERID: PARTNERID}, nil
}

func (c *PartnerServiceServer) FindPartnerByPhone(ctx context.Context, req *pb.FindPartnerByPhoneRequest) (*pb.FindPartnerByPhoneResponse, error) {
	utils.LogMessage(utils.Cyan, "FindUserByPhone Invoked")

	Partner, err := c.usecase.FindPartnerByPhone(ctx, req.GetMOBILE())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully Found User By Email")
	return &pb.FindPartnerByPhoneResponse{
		PARTNERID: Partner.PARTNERID,
	}, nil
}

func (c *PartnerServiceServer) FindPartnerByEmail(ctx context.Context, req *pb.FindPartnerEmailRequest) (*pb.FindPartnerByEmailResponse, error) {
	utils.LogMessage(utils.Cyan, "FindUserByPhone Invoked")

	Partner, err := c.usecase.FindPartnerByEmail(ctx, req.GetEMAIL())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully Found User By Email")
	return &pb.FindPartnerByEmailResponse{
		PARTNERID: Partner.PARTNERID,
	}, nil
}
