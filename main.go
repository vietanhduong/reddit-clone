package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"reddit-clone/server/handler"
	"reddit-clone/server/topics"
	"reddit-clone/server/users"
)

func main() {
	_ = godotenv.Load()

	server := echo.New()

	server.Use(echoMiddleware.Gzip())
	server.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions, http.MethodPatch},
	}))

	server.HTTPErrorHandler = handler.JSONHTTPErrorHandler

	RegisterClient(server)

	apiGroup := server.Group("/api")
	users.RegisterAPI(apiGroup)
	topics.RegisterAPI(apiGroup)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", GetPort())))
}

func RegisterClient(e *echo.Echo) {
	box := rice.MustFindBox("./client/build")
	client := http.FileServer(box.HTTPBox())
	e.GET("/static/*", echo.WrapHandler(client))
	e.GET("/*", func(c echo.Context) error {
		index, err := box.Open("index.html")
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(index)
		if err != nil {
			return err
		}
		return c.HTMLBlob(http.StatusOK, content)
	})
}

func GetPort() string {
	const DefaultPort = "8080"
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return DefaultPort
}
