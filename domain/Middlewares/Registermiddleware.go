package Middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetGlobalMiddleware(router *gin.Engine) {
	SetGlobalRecovery(router)
	SetLoggerWithWriter(router)
}
