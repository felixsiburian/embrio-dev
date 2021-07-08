package controller

import (
	"embrio-dev/service"
	"embrio-dev/service/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NasabahController struct {
	e           *echo.Echo
	nasabahCase service.INasabahUsecase
}

func NewNasabahController(e *echo.Echo, usecase service.INasabahUsecase) *NasabahController {
	return &NasabahController{
		e:           e,
		nasabahCase: usecase,
	}
}

func (ox *NasabahController) CreateNewNasabah(ec echo.Context) error {
	var form model.NasabahRegisterRequest

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	err = ox.nasabahCase.CreateNewUser(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id": "Berhasil",
		"en": "Success",
	})
}
