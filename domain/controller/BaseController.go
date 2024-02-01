package controller

import (
	"github.com/gin-gonic/gin"
	"goTest/domain/middlewares"
	"reflect"
	"strconv"
)

type Interface interface {
	Construct(c *gin.Context) *BaseController
}

func (b *BaseController) bindToReqDtoStruct(obj any) {
	if reflect.ValueOf(obj).Kind().String() != "ptr" {
		panic("bindToReqDtoStruct error")
	}

	_ = b.context.ShouldBind(obj)
}

type BaseController struct {
	UserId  int
	context *gin.Context
}

// Construct 实现ControllerInterface的Construct方法
func (b *BaseController) Construct(c *gin.Context) *BaseController {
	b.UserId, _ = strconv.Atoi(c.DefaultQuery("user_id", "0"))
	b.context = c
	return b
}

func (b *BaseController) returnJson(data any) {
	response := middlewares.Response{}
	response.Write(0, data, "success")(b.context)
}
