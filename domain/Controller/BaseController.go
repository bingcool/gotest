package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type Interface interface {
	Construct(c *gin.Context) *BaseController
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

// ResponseStruct returnJson 响应结构体
type ResponseStruct struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (b *BaseController) returnJson(data any) {
	valueType := reflect.ValueOf(data).Kind()
	switch valueType {
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		if reflect.ValueOf(data).Len() == 0 {
			data = make([]any, 0)
		}
	default:
	}
	// 响应结构体
	responseStruct := &ResponseStruct{}
	responseStruct.Code = 0
	responseStruct.Msg = "success"
	responseStruct.Data = data
	b.context.JSON(http.StatusOK, responseStruct)
}
