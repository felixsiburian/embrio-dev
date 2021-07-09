// register router here

package router

import (
	"embrio-dev/service"
	"embrio-dev/service/delivery/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, nasabahCase service.INasabahUsecase) {
	c := controller.NewNasabahController(e, nasabahCase)

	r := e.Group("nasabah/v1")
	r.POST("/create", c.CreateNewNasabah)
	r.POST("/auth", c.Auth)
}
