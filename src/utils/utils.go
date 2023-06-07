package utils

import (
	"reflect"
	"runtime"
)

// GetFunctionName returns the name of a function
func GetFunctionName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	dotIndex := len(fullName) - 1
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			dotIndex = i
			break
		}
	}
	return fullName[dotIndex+1:]
}
