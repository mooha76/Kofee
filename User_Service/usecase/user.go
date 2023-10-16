package usecase

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
	repo "github.com/mooha76/Kofee/User_Service/repository/interfaces"
	"github.com/mooha76/Kofee/User_Service/usecase/interfaces"
)

type userUsecase struct {
	repo repo.UserRepository
}

func NewUserUsecase(repo repo.UserRepository) interfaces.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (c *userUsecase) SaveUser(ctx context.Context, user model.Users) (userID uint64, err error) {

	return c.repo.SaveUser(ctx, user)
}

func (c *userUsecase) FindUserByEmail(ctx context.Context, email string) (model.Users, error) {
	return c.repo.FindUserByEmail(ctx, email)
}
func (c *userUsecase) FindUserByPhone(ctx context.Context, phone string) (model.Users, error) {
	return c.repo.FindUserByPhone(ctx, phone)
}
