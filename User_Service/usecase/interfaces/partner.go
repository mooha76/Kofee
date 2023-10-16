package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
)

type PartnerUsecase interface {
	SavePartner(ctx context.Context, Partners model.PARTNERSSINGUP_REQUEST) (PARTNERID uint64, err error)
	FindPartnerByPhone(ctx context.Context, Mobile string) (model.PARTNERS, error)
	FindPartnerByEmail(ctx context.Context, Email string) (model.PARTNERS, error)
}
