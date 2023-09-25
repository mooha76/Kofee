package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/mooha76/Kofee/Proxy_Service/api/handler/interfaces"
)

func SetupUserRoutes(user *gin.RouterGroup,
	userHandler handler.UserHandler, authHandler handler.AuthHandler) {

	auth := user.Group("/auth")
	{
		signup := auth.Group("/signup")
		{
			signup.POST("/", authHandler.UserSignup)

		}
	}

	profile := user.Group("/profile")
	{
		profile.GET("/", userHandler.GetProfile)
	}

}
