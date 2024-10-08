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

type UserServiceServer struct {
	usecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(usecase usecase.UserUsecase) pb.UserServiceServer {
	return &UserServiceServer{
		usecase: usecase,
	}
}
func (c *UserServiceServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	utils.LogMessage(utils.Cyan, "SaveUser Invoked")
	userID, err := c.usecase.SaveUser(context.Background(), model.Users{
		FirstName:  req.GetFirstName(),
		MiddleName: req.GetMiddleName(),
		LastName:   req.GetLastName(),
		Age:        req.GetAge(),
		Gender:     req.GetGender(),
		Email:      req.GetEmail(),
		Phone:      req.GetPhone(),
		Account:    req.GetAccount(),
		Password:   req.GetPassword(),
	})

	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully User Details Saved")
	return &pb.SaveUserResponse{UserId: userID}, nil
}

func (c *UserServiceServer) FindUserByEmail(ctx context.Context, req *pb.FindUserByEmailRequest) (*pb.FindUserByEmailResponse, error) {
	utils.LogMessage(utils.Cyan, "FindUserByEmail Invoked")
	user, err := c.usecase.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully Found User By Email")
	return &pb.FindUserByEmailResponse{
		UserId:      user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Phone:       user.Phone,
		Email:       user.Email,
		Password:    user.Password,
		Verified:    user.Verified,
		BlockStatus: user.BlockStatus,
	}, nil
}

func (c *UserServiceServer) FindUserByPhone(ctx context.Context, req *pb.FindUserByPhoneRequest) (*pb.FindUserByPhoneResponse, error) {
	utils.LogMessage(utils.Cyan, "FindUserByPhone Invoked")

	user, err := c.usecase.FindUserByPhone(ctx, req.GetPhone())
	if err != nil {
		utils.LogMessage(utils.Red, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}
	utils.LogMessage(utils.Green, "Successfully Found User By Email")
	return &pb.FindUserByPhoneResponse{
		UserId:      user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Phone:       user.Phone,
		Email:       user.Email,
		Password:    user.Password,
		Verified:    user.Verified,
		BlockStatus: user.BlockStatus,
	}, nil
}
