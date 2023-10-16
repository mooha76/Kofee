package usecase

import (
	"context"
	"fmt"

	client "github.com/mooha76/Kofee/Auth_service/client/interface"
	"github.com/mooha76/Kofee/Auth_service/model"
	interfaces "github.com/mooha76/Kofee/Auth_service/usecase/interface"
	"github.com/mooha76/Kofee/Auth_service/utils"
)

type partnerUsecase struct {
	partnerrClient client.PartnerClient
}

func NewPartnerUsecase(partnerrClient client.PartnerClient) interfaces.PartnerUseCase {
	return &partnerUsecase{

		partnerrClient: partnerrClient,
	}
}

// Signup for user
func (c *partnerUsecase) PartnerSingUp(ctx context.Context, Partner model.PartnerSingUpRequest) (PartnerId uint64, err error) {

	existUser, err := c.partnerrClient.FindPartnerByEmail(ctx, Partner.EMAIL)
	if err != nil {
		return 0, fmt.Errorf("failed to check Partner already exist \nerror:%s", err.Error())
	}
	// if user exist and verified also then return error
	if existUser.PARTNERID != 0 && existUser.ISACTIVE {
		return 0, fmt.Errorf("partner already exist with given details and verified")
	}

	PARTNERID := existUser.PARTNERID
	// if user not exist then save the user
	if PARTNERID == 0 {

		Partner.PASSWORD, err = utils.GenerateHashFromPassword(Partner.PASSWORD)
		if err != nil {
			return 0, fmt.Errorf("failed to hash user password")
		}

		PARTNERID, err = c.partnerrClient.SavePartner(ctx, Partner)
		if err != nil {
			return 0, fmt.Errorf("failed to save user \nerror%s", err.Error())
		}
	}

	return PARTNERID, err
}
