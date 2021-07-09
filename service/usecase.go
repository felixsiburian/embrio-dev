// all interface for usecase register here

package service

import (
	"embrio-dev/service/model"
	"github.com/labstack/echo/v4"
)

type ITokenUsecase interface {
	TokenValid(c echo.Context) error
	ExtractTokenString(c echo.Context) string
	ExtractTokenResponse(c echo.Context) (resp model.TokenResponse, err error)
}

type INasabahUsecase interface {
	CreateNewUser(args model.NasabahRegisterRequest) (err error)
	Auth(args model.NasbahLoginRequest) (resp model.NasbahLoginResponses, err error)
	TestExtractToken() (s string)
}
