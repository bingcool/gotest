package Middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetGlobalRecovery(router *gin.Engine) {
	router.Use(gin.Recovery())
}
