package interfaces

import "github.com/gin-gonic/gin"

type PartnerHandler interface {
	PartnerSignup(ctx *gin.Context)
}
