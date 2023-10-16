package service

import (
	"context"

	model "github.com/mooha76/Kofee/Auth_service/model"
	"github.com/mooha76/Kofee/Auth_service/pb"
	"github.com/mooha76/Kofee/Auth_service/utils"

	//"github.com/mooha76/Kofee/Auth_service/token"
	usecase "github.com/mooha76/Kofee/Auth_service/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServer struct {
	userusecase    usecase.AuthUseCase
	partnerusecase usecase.PartnerUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(userusecase usecase.AuthUseCase, partnerusecase usecase.PartnerUseCase) pb.AuthServiceServer {
	return &authServiceServer{
		userusecase:    userusecase,
		partnerusecase: partnerusecase,
	}
}

// User Sginup
func (c *authServiceServer) UserSignup(ctx context.Context, req *pb.UserSignupRequest) (*pb.UserSignupResponse, error) {
	utils.LogMessage(utils.Cyan, "SaveUser Invoked")
	signupRequest := model.UserSignupRequest{
		FirstName:  req.GetFirstName(),
		MiddleName: req.GetMiddleName(),
		LastName:   req.GetLastName(),
		Age:        req.GetAge(),
		Gender:     req.GetGender(),
		Email:      req.GetEmail(),
		Phone:      req.GetPhone(),
		Account:    req.GetAccount(),
		Password:   req.GetPassword(),
	}

	UserId, err := c.userusecase.UserSignup(ctx, signupRequest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserSignupResponse{User_Id: UserId}, nil
}
