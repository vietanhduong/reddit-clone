package common

import (
	"os"
	"reflect"

	"github.com/labstack/echo/v4"
)

type (
	Pair struct {
		Key   int
		Value int
	}
	Pairs []Pair
)

// Reference: https://golang.org/pkg/sort/#Interface
func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[j].Value < p[i].Value }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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
