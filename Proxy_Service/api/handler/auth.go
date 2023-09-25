package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mooha76/GoGrpcProxy/api/handler/interfaces"
	client "github.com/mooha76/GoGrpcProxy/client/interfaces"
	"github.com/mooha76/GoGrpcProxy/model"
	"github.com/mooha76/GoGrpcProxy/utils/response"
)

type authHandler struct {
	client client.AuthClient
}

func NewAuthHandler(client client.AuthClient) interfaces.AuthHandler {
	return &authHandler{
		client: client,
	}
}

func (c *authHandler) UserSignup(ctx *gin.Context) {

	var body model.UserSignupRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed bind inputs", err, body)
		return
	}

	userID, err := c.client.UserSignup(context.Background(), body)

	if err != nil {
		response.ErrorResponse(ctx, "failed to signup user", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully singUp created", userID)
}
