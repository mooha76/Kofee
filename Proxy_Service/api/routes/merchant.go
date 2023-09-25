package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/mooha76/GoGrpcProxy/api/handler/interfaces"
)

func SetupMerchantRoutes(merchant *gin.RouterGroup,
	merchantHandler handler.MerchantHandler) {

	mymerchant := merchant.Group("/Api")
	{
		merchnatApi := mymerchant.Group("/merchant")
		{
			merchnatApi.POST("/", merchantHandler.SEND_MERCHAN_API_PURCHASE_REQ)

		}
	}

}
