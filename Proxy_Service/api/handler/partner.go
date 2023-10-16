package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mooha76/Kofee/Proxy_Service/api/handler/interfaces"
	client "github.com/mooha76/Kofee/Proxy_Service/client/interfaces"
	"github.com/mooha76/Kofee/Proxy_Service/model"
	"github.com/mooha76/Kofee/Proxy_Service/utils/response"
)

type partnerhHandler struct {
	client client.PartnerClient
}

func NewPartnerHandler(client client.PartnerClient) interfaces.PartnerHandler {
	return &partnerhHandler{
		client: client,
	}
}

func (c *partnerhHandler) PartnerSignup(ctx *gin.Context) {

	var body model.PartnerSingUpRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, "failed bind inputs", err, body)
		return
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
	userID, err := c.client.PartnerSingUp(context.Background(), body)

	if err != nil {
		response.ErrorResponse(ctx, "Filled To create Partner ", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully PARTNER created", userID)
}
