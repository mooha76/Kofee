package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/utils/response"
)

type UserClient interface {
	GetUserProfile(ctx context.Context, userID uint64) (response.User, error)
}

type UsertestClint interface {
	Getmyusercline(ctx context.Context, userID uint64) (response.User, error)
}
