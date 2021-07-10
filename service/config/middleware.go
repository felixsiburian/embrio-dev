// all middleware register here

package config

import (
	"embrio-dev/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MiddlewareController struct {
	e         *echo.Echo
	tokenCase service.ITokenUsecase
}

func NewMiddlewareController(e *echo.Echo, tokenCase service.ITokenUsecase) *MiddlewareController {
	return &MiddlewareController{
		e:         e,
		tokenCase: tokenCase,
	}
}

func (ox *MiddlewareController) SetAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := ox.tokenCase.TokenValid(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}

func (ox *MiddlewareController) SetRefreshAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := ox.tokenCase.RefreshTokenValid(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}
