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
	tokenCase   service.ITokenUsecase
}

func NewNasabahController(e *echo.Echo, usecase service.INasabahUsecase, tokenCase service.ITokenUsecase) *NasabahController {
	return &NasabahController{
		e:           e,
		nasabahCase: usecase,
		tokenCase:   tokenCase,
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

func (ox *NasabahController) Auth(ec echo.Context) error {
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

func (ox *NasabahController) Ping(ec echo.Context) error {
	s := ox.nasabahCase.TestExtractToken()
	resp, err := ox.tokenCase.ExtractTokenResponse(ec)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"resp": resp,
		"s":    s,
	})
}
