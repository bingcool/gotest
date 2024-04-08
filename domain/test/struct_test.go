package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person  // 组合了Person结构体
	School  string
	Address Address
}

type Address struct {
	City    string
	Country string
}

func NewStudent() *Student {
	s := Student{}
	return &s
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) SetAddress(address Address) {
	s.Address = address
}

func TestStruct1(t *testing.T) {
	s := NewStudent()
	address := Address{City: "深圳", Country: "中国"}
	s.SetAddress(address)
	s.SetName("bingcool")
	fmt.Println(s.Address.City) // 输出 Tom 18 High School
}
