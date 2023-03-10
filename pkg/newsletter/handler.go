package newsletter

import "github.com/gin-gonic/gin"

type Handler interface {
	Get(ctx *gin.Context)
}
