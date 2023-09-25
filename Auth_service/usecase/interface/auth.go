package interfaces

import (
	"context"

	model "github.com/mooha76/Kofee/model"
)

type AuthUseCase interface {
	UserSignup(ctx context.Context, user model.UserSignupRequest) (CreatedUser string, err error)
}
