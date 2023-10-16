package interfaces

import (
	"context"

	model "github.com/mooha76/Kofee/Auth_service/model"
)

type AuthUseCase interface {
	UserSignup(ctx context.Context, user model.UserSignupRequest) (CreatedUser string, err error)
	//PartnerSingUp(ctx context.Context, Partner model.PartnerSingUpRequest) (PartnerId uint64, err error)
}
