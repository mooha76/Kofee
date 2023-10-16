package interfaces

import (
	"context"

	model "github.com/mooha76/Kofee/Auth_service/model"
)

type PartnerClient interface {
	FindPartnerByEmail(ctx context.Context, email string) (model.Partner, error)
	FindPartnerByPhone(ctx context.Context, phone string) (model.Partner, error)
	SavePartner(ctx context.Context, user model.PartnerSingUpRequest) (PARNERID uint64, err error)
}
