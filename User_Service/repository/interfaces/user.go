package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (model.Users, error)
	FindUserByPhone(ctx context.Context, phone string) (model.Users, error)

	SaveUser(ctx context.Context, user model.Users) (userID uint64, err error)
}
