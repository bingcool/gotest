package route

import (
	"goTest/domain/controller"
	"goTest/domain/dto/orderDto"
	"goTest/domain/middlewares"

	"github.com/gin-gonic/gin"
)

func SetOrderRouter(router *gin.Engine) {
	// 添加中间件
	v1 := router.Group("/api/v1")
	{
		// 路由中间件
		v1.Use(middlewares.ValidateLogin())

		// 路由处理
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// 路由处理
		v1.Any("/order/list", func(context *gin.Context) {
			listOrderReqDto := &orderDto.ListOrderReqDto{}
			_ = context.ShouldBind(listOrderReqDto)
			(controller.NewOrder(context)).ListOrder(context, listOrderReqDto)
		})

		// 路由处理-最前面的可以是api的中间件
		v1.Any("/order/list1", middlewares.ValidateLogin(), func(context *gin.Context) {
			listOrderReqDto := &orderDto.ListOrderReqDto{}
			_ = context.ShouldBind(listOrderReqDto)
			(controller.NewOrder(context)).ListOrder(context, listOrderReqDto)
		})

		// 路由处理sa
		v1.GET("/order/save", func(context *gin.Context) {
			(controller.NewOrder(context)).SaveOrder(context)
		})
	}

}
