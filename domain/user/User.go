package user

import "fmt"

const Name = "增冰"

type User struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func (user *User) SetName(name string) {
	user.Name = name
}

func (user *User) GetName() (name string) {
	return user.Name
}

func (user *User) SetAddress(province string, city string, district string) {
	user.Address.Province = province
	user.Address.City = city
	user.Address.District = district
}

func (user *User) GetAddress() Address {
	return user.Address
}

type MyInterface interface {
	MyMethod()
}

func (user *User) MyMethod() {
	fmt.Println("MyMethod")
}
