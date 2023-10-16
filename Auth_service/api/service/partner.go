package service

import (
	"context"

	model "github.com/mooha76/Kofee/Auth_service/model"
	"github.com/mooha76/Kofee/Auth_service/pb"
	"github.com/mooha76/Kofee/Auth_service/utils"

	//"github.com/mooha76/Kofee/Auth_service/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Partner Sginup
func (c *authServiceServer) PartnerSingUp(ctx context.Context, req *pb.PartnerSingUpRequest) (*pb.PartnerSingUpResponse, error) {
	utils.LogMessage(utils.Cyan, "SavePartner Invoked")
	partnersignupRequest := model.PartnerSingUpRequest{
		FIRSTNAME:               req.GetFIRSTNAME(),
		MIDDLENAME:              req.GetMIDDLENAME(),
		LASTNAME:                req.GetLASTNAME(),
		MOBILE:                  req.GetMOBILE(),
		ACCOUNTNO:               req.GetACCOUNTNO(),
		WITHDRAWALREQUESTTYPEID: req.GetWITHDRAWALREQUESTTYPEID(),
		AGE:                     req.GetAGE(),
		GENDER:                  req.GetGENDER(),
		EMAIL:                   req.GetEMAIL(),
		USERNAME:                req.GetUSERNAME(),
		PASSWORD:                req.GetPASSWORD(),
	}
	println("Log trace =================")
	UserId, err := c.partnerusecase.PartnerSingUp(ctx, partnersignupRequest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.PartnerSingUpResponse{PARTNERID: UserId}, nil
}
