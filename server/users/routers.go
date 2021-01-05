package users

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"reddit-clone/server/common"
	"reddit-clone/server/handler"
)

type (
	UserService interface {
		Login(username string, password string) (string, error)
		GetUserByUsername(username string) (*User, error)
	}

	ServerImpl struct {
		userService UserService
	}
)

func RegisterAPI(api *echo.Group) {
	userRepo := NewRepository()
	userService := NewServiceImpl(userRepo)
	server := ServerImpl{userService: userService}

	// Register Login API outside `userEndpoints`
	api.POST("/login", server.login)

	// User endpoints
	userEndpoints := api.Group("/users")
	userEndpoints.GET("/me", server.me, handler.IsLoggedIn)
}

func (s *ServerImpl) login(ctx echo.Context) error {
	var user *UserRequest
	if err := ctx.Bind(&user); err != nil {
		return err
	}

	if common.IsEmpty(user.Username) || common.IsEmpty(user.Password) {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "username and password is required",
		}
	}
	accessToken, err := s.userService.Login(user.Username, user.Password)
	if err != nil {
		return err
	}
	response := map[string]interface{}{
		"access_token": accessToken,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (s *ServerImpl) me(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	userModel, err := s.userService.GetUserByUsername(username)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		Code:    http.StatusOK,
		Content: userModel,
	})
}
