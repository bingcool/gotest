package test

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gutil"
	"testing"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	Fn    func(phone string)
}

type Student struct {
	*Person   // 组合了Person结构体
	Education Education
	School    string
	Address   Address
	Cars      []string
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

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type B bool

func TestStruct3(t *testing.T) {

	s := B(true)

	fmt.Println(s)

	//peoples := []people{
	//	{Name: "a", Age: 20},
	//	{Name: "b", Age: 10},
	//	{Name: "c", Age: 17},
	//	{Name: "d", Age: 8},
	//	{Name: "e", Age: 28},
	//}
	//s1 := "aaa"
	//s := string(s1)
	//fmt.Println(s)
	//
	//fmt.Println(peoples)
}

func TestStruct4(t *testing.T) {
	var s Student
	s.Name = "bingcool"
	s.Age = 18
	s.Address = Address{City: "深圳", Country: "中国"}

	s1 := s

	s1.Age = 19

	fmt.Println(s.Age, s1.Age)

}

type Education interface {
	Get() string
}

type MeEducation struct {
	XueLi string
}

func (m *MeEducation) Get() string {
	return `11`
}

func TestStruct5(t *testing.T) {
	s := &Student{
		Person: &Person{
			Name: "b",
			Age:  18,
		},
		Education: &MeEducation{XueLi: "大学"},
		Cars:      make([]string, 0),
		Address: Address{
			City:    "深圳",
			Country: "中国",
		},
	}

	s1 := s

	s1.Age = 19

	s1.Cars = append(s1.Cars, "a")

	CopyStudent(s)
	gutil.Dump(s.Cars)

}

func CopyStudent(s *Student) {
	s.Cars = append(s.Cars, "b")
	gutil.Dump(s.Cars)
}
