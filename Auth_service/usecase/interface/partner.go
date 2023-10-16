package interfaces

import (
	"context"

	model "github.com/mooha76/Kofee/Auth_service/model"
)

type PartnerUseCase interface {
	PartnerSingUp(ctx context.Context, Partner model.PartnerSingUpRequest) (PartnerId uint64, err error)
}
