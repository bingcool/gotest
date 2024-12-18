package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	Fn    func(phone string)
}

type Student struct {
	Person  // 组合了Person结构体
	School  string
	Address Address
}

type Teacher struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type Address struct {
	City    string
	Country string
}

func NewStudent() *Student {
	s := &Student{}
	return s
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func (s *Student) SetAddress(address Address) {
	s.Address = address
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func TestStruct1(t *testing.T) {
	s := NewStudent()
	address := Address{City: "深圳", Country: "中国"}
	s.SetAddress(address)
	s.SetName("bingcool")
	s.SetAge(18)
	s.Fn = func(phone string) {
		s.Phone = phone
	}
	s.Fn("123456")
	fmt.Println(s.Address.City) // 输出 Tom 18 High School
	fmt.Println(s.Age)
	fmt.Println(s.Phone)
}

func TestStruct2(t *testing.T) {
	var s Student

	fmt.Println(s.Age)

	address := Address{City: "深圳", Country: "中国"}
	s.SetAddress(address)
	s.SetName("bingcool")
	fmt.Println(s.Address.City) // 输出 Tom 18 High School
}

func TestStruct3(t *testing.T) {

}
