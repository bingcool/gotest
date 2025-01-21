package entity

import "goTest/domain/dao/model"

const TableNamePerson = "people"

// Person mapped from table <people>
type Person struct {
	model.Person
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}
