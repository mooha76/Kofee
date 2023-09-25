package interfaces

import "github.com/gin-gonic/gin"

type MerchantHandler interface {
	SEND_MERCHAN_API_PURCHASE_REQ(ctx *gin.Context)
}
