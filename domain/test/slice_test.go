package test

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/juju/errors"
	"reflect"
	"sort"
	"testing"
)

type CMAP[T int | float64] []T

// TestSelect3 泛型处理
func TestSelect3(t *testing.T) {
	// 声明一个有容量的切片，长度为1，容量为5，这时候0~长度的位置的值将初始化为零值
	cmap1 := make(CMAP[int], 1, 5)
	cmap1 = append(cmap1, 1)
	fmt.Println(cmap1)

	// 声明一个空的切片
	var cmap2 CMAP[float64]
	fmt.Println(cmap2)
	cmap2 = append(cmap2, 1.1)
	fmt.Println(cmap2)

	// 初始化切片并且初始化值
	cmap3 := CMAP[int]{1, 2, 3, 4}
	fmt.Println(cmap3)
}

func TestSelect4(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[0])

	s1 := make([]int, 1, 5)

	fmt.Println(s1)

	var s2 []int

	fmt.Println(s2)

}

func TestSelect5(t *testing.T) {
	arr := [2]int{1, 2}
	fmt.Println(arr)

	var arr1 [2]int

	fmt.Println(arr1)

	reflectType := reflect.TypeOf(arr)
	reflectValue := reflect.ValueOf(arr)
	fmt.Printf("[typeof:%v]\n", reflectType)   // string
	fmt.Printf("[valueof:%v]\n", reflectValue) // abc
}

func TestSlice6(t *testing.T) {
	slice := make([]int, 3, 5) // 创建长度为3，容量为5的切片
	slice = append(slice, 1)
	fmt.Println(slice)

	fmt.Println(len(slice)) // 返回长度

	fmt.Println(cap(slice)) // 返回容量
}

// 创建空切片[]
func TestSlice7(t *testing.T) {
	slice1 := make([]int, 0)

	fmt.Println(slice1)
}

func TestSlice8(t *testing.T) {
	//s1 := gconv.SliceAny(123)
	//fmt.Println(s1)

	//s2 := gconv.Convert(123, "[]int")
	//fmt.Println(s2)
	//fmt.Println(reflect.TypeOf(s2).String())
	//
	//s3 := gconv.Int8(100)
	//fmt.Println(s3)

	map1 := make(map[string]any)
	map1["name"] = "zhangsan"

	fmt.Println(map1)

	s4, found := gutil.ItemValue(map1, "name")
	if found {
		fmt.Println(s4)
	}

	keys := gutil.Keys(map1)

	fmt.Println(keys)

}

func TestSlice9(t *testing.T) {
	sliceMap1 := make(map[string]any)
	sliceMap1["id"] = 12
	sliceMap1["name"] = "zhangsan12"

	sliceMap2 := make(map[string]any)
	sliceMap2["id"] = 34
	sliceMap2["name"] = "zhangsan34"

	SliceArr := make([]map[string]any, 0)
	SliceArr = append(SliceArr, sliceMap1, sliceMap2)
	a := gutil.SliceToMapWithColumnAsKey(SliceArr, "id")

	gutil.DumpWithType(a)

	sort.Slice(SliceArr, func(i, j int) bool {
		v1, ok1 := SliceArr[i]["id"].(int)
		v2, ok2 := SliceArr[j]["id"].(int)
		if !ok1 || !ok2 {
			errors.New("id is not int")
		}
		return v1 > v2
	})

	gutil.DumpWithType(SliceArr)

}

func TestSlice10(t *testing.T) {
	p1 := &Person{
		Name: "zhangsan12",
		Age:  12,
	}

	p2 := &Person{
		Name: "zhangsan34",
		Age:  34,
	}

	SliceArr := make([]*Person, 0)
	SliceArr = append(SliceArr, p1, p2)

	gutil.DumpWithType(SliceArr)

	for _, item := range SliceArr {
		gutil.DumpWithType(item)
	}

	list := gutil.SliceToMapWithColumnAsKey(SliceArr, "Name")
	for _, item1 := range list {
		switch v := item1.(type) {
		case *Person:
			gutil.DumpWithType(v.Name)
		}
	}

	if item3, ok := list["zhangsan12"]; ok {
		switch v := item3.(type) {
		case *Person:
			gutil.DumpWithType(v.Name)
		}
	}

}

// CompareType 可比较的泛型类型限归集
type CompareType interface {
	string | int | int16 | int32 | int64 | int8 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func CompareAny[V CompareType](a, b V) int {
	if a == b {
		return 0
	}

	if a < b {
		return -1
	}

	return 1
}

func TestSlice11(t *testing.T) {
	p1 := &Person{
		Name: "zhangsan12",
		Age:  12,
	}

	p2 := &Person{
		Name: "zhangsan34",
		Age:  34,
	}

	SliceArr := make([]*Person, 0)
	SliceArr = append(SliceArr, p1, p2)

	index, found := sort.Find(len(SliceArr), func(i int) int {
		a := 34
		b := SliceArr[i].Age
		return CompareAny[int](a, b)
	})

	gutil.DumpWithType(index)
	gutil.DumpWithType(found)

	sort.Slice(SliceArr, func(i, j int) bool {
		return SliceArr[i].Age > SliceArr[j].Age
	})

	gutil.DumpWithType(SliceArr)

}

func TestSlice12(t *testing.T) {
	p1 := &Person{
		Name: "zhangsan12",
		Age:  12,
	}

	p2 := &Person{
		Name: "zhangsan34",
		Age:  34,
	}

	SliceArr := make([]*Person, 0)
	SliceArr = append(SliceArr, p1, p2)

	values := gutil.ListItemValues(SliceArr, "Name")

	fmt.Println(values)
}

func TestSlice13(t *testing.T) {
	jsonContent := `{"name":"john", "score":"100"}`
	j := gjson.New(jsonContent)
	fmt.Println(j.Get("name"))
	fmt.Println(j.Get("score"))

	j.Set("score", 150)
	s1, _ := j.ToJson()
	s := string(s1)
	fmt.Println(s)
}

func TestSlice14(t *testing.T) {
	p1 := &Teacher{
		Name: "zhangsan12",
		Age:  12,
	}

	p2 := &Teacher{
		Name: "zhangsan34",
		Age:  34,
	}

	SliceArr := make([]*Teacher, 0)
	SliceArr = append(SliceArr, p1, p2)

	j := gjson.New(SliceArr)

	fmt.Println(j.ToJsonString())

	//gutil.Dump(j.Get("0"))
}
