package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/mooha76/Kofee/Proxy_Service/api/handler/interfaces"
)

func SetupPartnerRoutes(user *gin.RouterGroup, partnerHandler handler.PartnerHandler) {

	mypartner := user.Group("/api")
	{
		partnerApi := mypartner.Group("/partner/create")
		{
			partnerApi.POST("/", partnerHandler.PartnerSignup)

		}
	}

}
