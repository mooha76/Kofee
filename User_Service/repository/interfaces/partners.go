package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/User_Service/model"
)

type PartnerRepository interface {
	FindPartnerByEmail(ctx context.Context, EMAIL string) (model.PARTNERS, error)
	FindPartnerByPhone(ctx context.Context, MOBILE string) (model.PARTNERS, error)
	SaveParner(ctx context.Context, Partner model.PARTNERSSINGUP_REQUEST) (PARTNERID uint64, err error)
}
