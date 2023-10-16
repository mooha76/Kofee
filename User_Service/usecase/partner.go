package usecase

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
	repo "github.com/mooha76/Kofee/User_Service/repository/interfaces"
	"github.com/mooha76/Kofee/User_Service/usecase/interfaces"
)

type PartnerUsecase struct {
	repo repo.PartnerRepository
}

func NewPartnerUsecase(repo repo.PartnerRepository) interfaces.PartnerUsecase {
	return &PartnerUsecase{
		repo: repo,
	}
}

func (c *PartnerUsecase) SavePartner(ctx context.Context, user model.PARTNERSSINGUP_REQUEST) (userID uint64, err error) {

	return c.repo.SaveParner(ctx, user)
}

func (c *PartnerUsecase) FindPartnerByEmail(ctx context.Context, email string) (model.PARTNERS, error) {
	return c.repo.FindPartnerByEmail(ctx, email)
}
func (c *PartnerUsecase) FindPartnerByPhone(ctx context.Context, phone string) (model.PARTNERS, error) {
	return c.repo.FindPartnerByPhone(ctx, phone)
}
