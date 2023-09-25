package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/model"
)

type MerchantClient interface {
	SEND_MERCHAN_API_PURCHASE_REQ(ctx context.Context, req model.Request) (Reqstatus string, err error)
}
