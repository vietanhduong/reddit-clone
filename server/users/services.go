package users

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"reddit-clone/server/common"
)

type (
	UserRepository interface {
		Login(username string, password string) *User
		FindByUsername(username string) *User
	}

	ServiceImpl struct {
		userRepo UserRepository
	}
)

func NewServiceImpl(userRepo UserRepository) *ServiceImpl {
	return &ServiceImpl{userRepo: userRepo}
}

func (s *ServiceImpl) Login(username string, password string) (string, error) {
	user := s.userRepo.Login(username, password)
	if common.IsNil(user) {
		return "", &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "username or password invalid",
		}
	}
	return GenerateToken(user)
}

func (s *ServiceImpl) GetUserByUsername(username string) (*User, error) {
	user := s.userRepo.FindByUsername(username)
	if common.IsNil(user) {
		return nil, &echo.HTTPError{
			Code:    http.StatusNotFound,
			Message: "username does not exist",
		}
	}
	return user, nil
}

func GenerateToken(user *User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	secretKey := common.GetEnv("SECRET_KEY", "very-secret")

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["name"] = user.FullName
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, err
}
