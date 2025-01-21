package test

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/mitchellh/mapstructure"
	"log"
	"testing"
)

func initPeople() []*people {
	userList := []*people{
		{"bing1", 30},
		{"bing2", 31},
	}

	return userList
}

func TestLanJson(t *testing.T) {
	userList := initPeople()
	gutil.Dump(userList)
}

func TestLanJson1(t *testing.T) {
	userList := initPeople()
	map1 := gconv.Maps(userList)
	gutil.Dump(map1[0]["name"])
}

func TestLanJson2(t *testing.T) {
	userList := initPeople()
	values := gutil.ListItemValues(userList, "Name")

	gutil.Dump(values)
}

func TestLanJson3(t *testing.T) {
	userList := initPeople()
	map1 := gconv.Maps(userList)
	values := gutil.ListToMapByKey(map1, "name")
	gutil.Dump(values)
}

// map转换成结构体
func TestLanJson4(t *testing.T) {
	userList := initPeople()
	map1 := gconv.Maps(userList)

	var userList2 []*people
	gconv.Scan(map1, &userList2)

	gutil.Dump(userList2)
}

// map转换成结构体--借助第三方库-mapstructure
func TestLanJson5(t *testing.T) {
	userList := initPeople()
	map1 := gconv.Maps(userList)

	var userList2 []*people

	// 配置解码器
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &userList2,
	})

	// 解码 map 到结构体
	if err := decoder.Decode(map1); err != nil {
		log.Fatalf("Failed to decode map to struct: %v", err)
	}

	gutil.Dump(userList2)

}
