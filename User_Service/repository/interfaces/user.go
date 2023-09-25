package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (model.User, error)
	FindUserByPhone(ctx context.Context, phone string) (model.User, error)

	SaveUser(ctx context.Context, user model.User) (userID uint64, err error)
}
