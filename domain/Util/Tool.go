package Util

import "reflect"

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

func ContainsInSlice(slice []string, targetItem string) bool {
	for _, item := range slice {
		if item == targetItem {
			return true
		}
	}
	return false
}
