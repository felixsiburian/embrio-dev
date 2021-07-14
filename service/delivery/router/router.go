// register router here

package router

import (
	"embrio-dev/service"
	"embrio-dev/service/config"
	"embrio-dev/service/delivery/handler"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, nasabahCase service.INasabahUsecase, tokenCase service.ITokenUsecase) {
	c := handler.NewNasabahHandler(e, nasabahCase, tokenCase)
	t := handler.NewTokenHandler(e, tokenCase)
	m := config.NewMiddlewareController(e, tokenCase)

	r := e.Group("nasabah/v1")
	r.POST("/create", c.CreateNewNasabah)
	r.POST("/auth", c.Auth)

	r.POST("/refresh", t.Refresh, m.SetRefreshAuthentication)

	//example how to use token as header or you can use this to check user token translation
	r.GET("/ping/access/token", t.Ping, m.SetAuthentication)
	r.GET("/ping/refresh/token", t.PingRefreshToken, m.SetRefreshAuthentication)
}
