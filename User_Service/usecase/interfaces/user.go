package interfaces

import (
	"context"

	"github.com/mooha76/GoUser_Service-Grpc/model"
)

type UserUsecase interface {
	SaveUser(ctx context.Context, user model.User) (userID uint64, err error)
	FindUserByPhone(ctx context.Context, phone string) (model.User, error)
	FindUserByEmail(ctx context.Context, email string) (model.User, error)
}
