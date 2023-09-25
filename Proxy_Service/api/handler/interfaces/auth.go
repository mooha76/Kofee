package interfaces

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	UserSignup(ctx *gin.Context)
}
