package util

import (
	"github.com/mozillazg/go-pinyin"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func IsMap(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Map
}

func IsStruct(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Struct
}

func IsSlice(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Slice
}

func IsChan(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Chan
}

func ContainsInSliceString(slice []string, targetItem string) (bool, int) {
	for i, item := range slice {
		i := i
		if item == targetItem {
			return true, i
		}
	}
	return false, -1
}

func ContainsInSliceInt(slice []int, targetItem int) (bool, int) {
	for i, item := range slice {
		i := i
		if item == targetItem {
			return true, i
		}
	}
	return false, -1
}

func ContainsInSliceFloat(slice []int, targetItem int) (bool, int) {
	for i, item := range slice {
		i := i
		if item == targetItem {
			return true, i
		}
	}
	return false, -1
}

func IsInt(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func IsFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func IsNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// SnakeString 驼峰转蛇形
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// CamelString 蛇形转驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// Cors 如果用了网关，不要在代码实现可以
func Cors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") // header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    // 设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             // 允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              // 返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

// ChinesConvertPinyin 中文转拼音(支持大写)
func ChinesConvertPinyin(s string, uppercase bool) (py []string) {
	//var py []string
	pyStr := s
	newPy := pinyin.NewArgs()
	rs := []rune(s)
	for _, r := range rs {
		if unicode.Is(unicode.Han, r) {
			singlePy := pinyin.SinglePinyin(r, newPy)
			py = append(py, singlePy...)
		} else {
			singlePy := string(r)
			py = append(py, singlePy)
		}
	}
	pyStr = strings.Join(py, "")

	if uppercase {
		pyStr = strings.ToUpper(pyStr)
	}
	py = strings.Split(pyStr, " ")
	return
}

func IsChinese(str string) bool {
	for _, char := range str {
		if !unicode.Is(unicode.Han, char) {
			return false
		}
	}
	return true
}

// GetStrPointerVal 获取字符串指针的值
func GetStrPointerVal(str *string) (resp string) {
	resp = ""
	if str != nil {
		resp = *str
	}
	return
}

// ToLowerFirstSting 首字母小写
func ToLowerFirstSting(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	rs[0] = unicode.ToLower(rs[0])
	return string(rs)
}

// ToUpperFirstSting 首字母大写
func ToUpperFirstSting(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}

// ReplaceAllBySliceSting 批量替换
func ReplaceAllBySliceSting(s string, old []string, new []string) string {
	if s == "" {
		return ""
	}
	for i, v := range old {
		if len(new) == 1 {
			s = strings.ReplaceAll(s, v, new[0])
		} else {
			s = strings.ReplaceAll(s, v, new[i])
		}
	}
	return s
}
