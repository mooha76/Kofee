package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mooha76/Kofee/Proxy_Service/api/handler/interfaces"
	client "github.com/mooha76/Kofee/Proxy_Service/client/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/model"
	"github.com/mooha76/Kofee/Proxy_Service/utils/response"
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

	var body model.PartnerSingUpRequest

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
