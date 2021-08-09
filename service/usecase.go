// all interface for usecase register here

package service

import (
	"embrio-dev/service/model"
	"embrio-dev/service/model/db/tableModel"
	"github.com/labstack/echo/v4"
)

type ITokenUsecase interface {
	TokenValid(c echo.Context) error
	RefreshTokenValid(c echo.Context) error
	ExtractTokenString(c echo.Context) string
	ExtractTokenResponse(c echo.Context) (resp model.TokenResponse, err error)
	ExtractRefreshTokenResponse(c echo.Context) (resp model.TokenResponse, err error)
	RefreshToken(c echo.Context) (resp model.NasbahLoginResponses, err error)
}

type INasabahUsecase interface {
	CreateNewUser(args model.NasabahRegisterRequest) (err error)
	Auth(args model.NasbahLoginRequest) (resp model.NasbahLoginResponses, err error)
}

type IRekeningUsecase interface {
	CreateRekening(args tableModel.Rekening) (err error)
	GetSaldoNasabah(nasabahID int64, noRekening string) (res tableModel.GetSaldoNasabah, err error)
	TopUpSaldoNasabah(args tableModel.TopUpRekeningArgs) (err error)
}
