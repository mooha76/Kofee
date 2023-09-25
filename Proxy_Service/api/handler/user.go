package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mooha76/Kofee/Proxy_Service/api/handler/interfaces"
	client "github.com/mooha76/Kofee/Proxy_Service/client/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/utils"
	"github.com/mooha76/Kofee/Proxy_Service/utils/response"
)

type userHandler struct {
	client client.UserClient
}

func NewUserHandler(client client.UserClient) interfaces.UserHandler {
	return &userHandler{
		client: client,
	}
}

func (c *userHandler) GetProfile(ctx *gin.Context) {

	userID := utils.GetUserIDFromContext(ctx)

	user, err := c.client.GetUserProfile(ctx, userID)

	if err != nil {
		response.ErrorResponse(ctx, "failed get user details", err, nil)
		return

	}

	response.SuccessResponse(ctx, "successfully got user details", user)
}
