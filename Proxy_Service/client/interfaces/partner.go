package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/model"
)

type PartnerClient interface {
	PartnerSingUp(ctx context.Context, req model.PartnerSingUpRequest) (Id uint64, err error)
}
