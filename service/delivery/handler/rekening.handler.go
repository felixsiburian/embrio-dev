package handler

import (
	"embrio-dev/service"
	"embrio-dev/service/model/db/tableModel"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RekeningHandler struct {
	e            *echo.Echo
	tokenCase    service.ITokenUsecase
	rekeningCase service.IRekeningUsecase
}

func NewRekeningHandler(
	e *echo.Echo,
	tokenCase service.ITokenUsecase,
	rekeningCase service.IRekeningUsecase,
) *RekeningHandler {
	return &RekeningHandler{
		e:            e,
		tokenCase:    tokenCase,
		rekeningCase: rekeningCase,
	}
}

func (ox *RekeningHandler) CreateRekening(ec echo.Context) error {
	nasabahInfo, err := ox.tokenCase.ExtractTokenResponse(ec)
	if err != nil {
		return ec.JSON(http.StatusUnauthorized, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	var form tableModel.Rekening
	err = json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	form.NasabahID = nasabahInfo.NasabahID
	err = ox.rekeningCase.CreateRekening(form)
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

func (ox *RekeningHandler) GetSaldoNasabah(ec echo.Context) error {
	noRekening := ec.QueryParam("no_rekening")
	nasabahInfo, err := ox.tokenCase.ExtractTokenResponse(ec)
	if err != nil {
		return ec.JSON(http.StatusUnauthorized, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.rekeningCase.GetSaldoNasabah(nasabahInfo.NasabahID, noRekening)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})

	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"data": res,
	})
}

func (ox *RekeningHandler) TopUpSaldoNasabah(ec echo.Context) error {
	nasabahInfo, err := ox.tokenCase.ExtractTokenResponse(ec)
	if err != nil {
		return ec.JSON(http.StatusUnauthorized, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	var form tableModel.TopUpRekeningArgs
	err = json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	form.NasabahID = nasabahInfo.NasabahID
	form.OperatedBy = nasabahInfo.Fullname
	err = ox.rekeningCase.TopUpSaldoNasabah(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id":           "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en":           "Something went wrong. Please try again later",
			"internal_msg": err.Error(),
		})
	}
	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id": "Berhasil",
		"en": "Success",
	})
}
