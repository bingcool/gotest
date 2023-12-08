package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type ControllerConstruct interface {
	Construct(c *gin.Context) *BaseController
}

type BaseController struct {
	UserId  int
	context *gin.Context
}

func (b *BaseController) Construct(c *gin.Context) *BaseController {
	b.UserId, _ = strconv.Atoi(c.DefaultQuery("user_id", "0"))
	b.context = c
	return b
}

// returnJson 响应结构体
func (b *BaseController) returnJson(data any) {
	// 响应结构体
	responseStruct := struct {
		Code int
		Msg  string
		Data any
	}{}
	responseStruct.Code = 0
	responseStruct.Msg = "success"
	responseStruct.Data = data
	b.context.JSON(http.StatusOK, responseStruct)
}

func structToMap(obj interface{}) any {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// 确保 obj 是一个结构体类型
	if t.Kind() != reflect.Struct {
		return v
	}

	// 创建一个空的 map
	m := make(map[string]interface{})

	// 遍历结构体的字段并将字段名和字段值存储在 map 中
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		m[field.Name] = value
	}

	return m
}
