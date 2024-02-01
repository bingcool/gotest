package util

import (
	"reflect"
	"strconv"
)

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

func ContainsInSliceString(slice []string, targetItem string) (bool, int) {
	for i, item := range slice {
		i := i
		if item == targetItem {
			return true, i
		}
	}
	return false, -1
}

func ContainsInSliceInt(slice []int, targetItem int) (bool, int) {
	for i, item := range slice {
		i := i
		if item == targetItem {
			return true, i
		}
	}
	return false, -1
}

func ContainsInSliceFloat(slice []int, targetItem int) (bool, int) {
	for i, item := range slice {
		i := i
		if item == targetItem {
			return true, i
		}
	}
	return false, -1
}

func IsInt(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func IsFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func IsNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}
