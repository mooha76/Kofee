package interfaces

import (
	"context"

	"github.com/mooha76/Kofee/Proxy_Service/model"
)

type AuthClient interface {
	UserSignup(ctx context.Context, req model.UserSignupRequest) (otpID string, err error)
}
