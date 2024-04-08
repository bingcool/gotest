package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	str := `{name:'bingcool'}`
	strByte := []byte(str)
	var a any
	err := json.Unmarshal(strByte, &a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(a)
}
