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

type merchantHandler struct {
	client client.MerchantClient
}

func NewMerchantHandler(client client.MerchantClient) interfaces.MerchantHandler {
	return &merchantHandler{
		client: client,
	}
}

func (c *merchantHandler) SEND_MERCHAN_API_PURCHASE_REQ(ctx *gin.Context) {

	var body model.Request

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
	userID, err := c.client.SEND_MERCHAN_API_PURCHASE_REQ(context.Background(), body)

	if err != nil {
		response.ErrorResponse(ctx, "Filled to send merchnat API Request ", err, nil)
		return
	}

	response.SuccessResponse(ctx, "successfully singUp created", userID)
}
