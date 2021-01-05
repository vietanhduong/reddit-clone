package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/vietanhduong/reddit-clone/server/handler"
	"github.com/vietanhduong/reddit-clone/server/topics"
	"github.com/vietanhduong/reddit-clone/server/users"
)

func main() {
	_ = godotenv.Load()

	server := echo.New()

	server.Use(echoMiddleware.Gzip())
	server.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions, http.MethodPatch},
	}))

	server.HTTPErrorHandler = handler.JSONHTTPErrorHandler

	apiGroup := server.Group("/api")
	users.RegisterAPI(apiGroup)
	topics.RegisterAPI(apiGroup)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", GetPort())))
}

func GetPort() string {
	const DefaultPort = "8080"
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return DefaultPort
}
