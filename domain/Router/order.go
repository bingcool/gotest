package Router

import (
	"github.com/gin-gonic/gin"
	"goTest/domain/Controller"
)

func SetOrderRouter(router *gin.Engine) {
	// 添加中间件
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.Any("/order/list", func(context *gin.Context) {
			(Controller.NewOrder(context)).ListOrder(context)
		})

		v1.GET("/order/save", func(context *gin.Context) {
			(Controller.NewOrder(context)).SaveOrder(context)
		})
	}

}
