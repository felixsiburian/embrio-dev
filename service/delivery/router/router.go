// register router here

package router

import (
	"embrio-dev/service"
	"embrio-dev/service/config"
	"embrio-dev/service/delivery/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, nasabahCase service.INasabahUsecase, tokenCase service.ITokenUsecase) {
	c := controller.NewNasabahController(e, nasabahCase, tokenCase)
	m := config.NewMiddlewareController(e, tokenCase)

	r := e.Group("nasabah/v1")
	r.POST("/create", c.CreateNewNasabah)
	r.POST("/auth", c.Auth)

	//example how to use token as header or you can use this to check user token translation
	r.GET("/ping", c.Ping, m.SetAuthentication)
}
