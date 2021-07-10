package handler

import (
	"embrio-dev/service"
	"embrio-dev/service/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NasabahHandler struct {
	e           *echo.Echo
	nasabahCase service.INasabahUsecase
	tokenCase   service.ITokenUsecase
}

func NewNasabahHandler(
	e *echo.Echo,
	usecase service.INasabahUsecase,
	tokenCase service.ITokenUsecase,
) *NasabahHandler {
	return &NasabahHandler{
		e:           e,
		nasabahCase: usecase,
		tokenCase:   tokenCase,
	}
}

func (ox *NasabahHandler) CreateNewNasabah(ec echo.Context) error {
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

func (ox *NasabahHandler) Auth(ec echo.Context) error {
	var form model.NasbahLoginRequest

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	resp, err := ox.nasabahCase.Auth(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	tokens := map[string]string{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":    "Berhasil",
		"en":    "Success",
		"token": tokens,
	})
}
