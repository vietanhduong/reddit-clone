package common

import (
	"os"
	"reflect"

	"github.com/labstack/echo/v4"
)

func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func IsEmpty(s string) bool {
	return len(s) == 0
}

func GetEnv(key string, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultValue
}

func HttpError(code int, message interface{}) *echo.HTTPError {
	return &echo.HTTPError{
		Code:    code,
		Message: message,
	}
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
