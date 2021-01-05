package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var msg interface{}

	msg = "internal server error"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message

	}
	response := map[string]interface{}{
		"code":    code,
		"message": msg,
	}
	if err := c.JSON(code, response); err != nil {
		c.Logger().Error(err)
	}

	c.Logger().Error(err)
}
