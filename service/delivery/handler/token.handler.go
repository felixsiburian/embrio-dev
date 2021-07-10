package handler

import (
	"embrio-dev/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type TokenHandler struct {
	e         *echo.Echo
	tokenCase service.ITokenUsecase
}

func NewTokenHandler(
	e *echo.Echo,
	tokenCase service.ITokenUsecase,
) *TokenHandler {
	return &TokenHandler{
		e:         e,
		tokenCase: tokenCase,
	}
}

func (ox *TokenHandler) Ping(ec echo.Context) error {
	resp, err := ox.tokenCase.ExtractTokenResponse(ec)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"ping": "PING!!!",
		"id":   "Berhasil",
		"en":   "Success",
		"resp": resp,
	})
}

func (ox *TokenHandler) PingRefreshToken(ec echo.Context) error {
	resp, err := ox.tokenCase.ExtractRefreshTokenResponse(ec)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"ping": "PING!!!",
		"id":   "Berhasil",
		"en":   "Success",
		"resp": resp,
	})
}

func (ox *TokenHandler) Refresh(ec echo.Context) error {
	resp, err := ox.tokenCase.RefreshToken(ec)
	if err != nil {
		log.Println(err)
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
