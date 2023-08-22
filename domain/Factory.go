package domain

import (
	"gofy/domain/user"
)

func Show() {
	//myname := user.Name
	userInstance := new(user.User)

	userInstance.SetName("bingcoolhuang")
	userInstance.SetAddress("广东省", "深圳市", "宝安区")

	//fmt.Println(userInstance.GetName())
	//fmt.Println(userInstance.GetAddress().Province)

	userInstance1 := new(user.User)
	userInstance1.SetName("bingcoolhuang")
	userInstance1.SetAddress("福建省", "深圳市", "宝安区")

	//fmt.Println(userInstance1.GetName())
	//fmt.Println(userInstance1.GetAddress().Province)

	userInstance.MyMethod()
}
