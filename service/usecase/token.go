package usecase

import (
	"embrio-dev/service"
	"embrio-dev/service/model"
	"embrio-dev/service/tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"strings"
)

type tokenUsecase struct {
	tokenRepo service.ITokenRepository
}

func NewTokenUsecase(tokenRepo service.ITokenRepository) service.ITokenUsecase {
	return tokenUsecase{
		tokenRepo: tokenRepo,
	}
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(b))
}

func (ox tokenUsecase) TokenValid(c echo.Context) error {
	tokenString := ox.ExtractTokenString(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func (ox tokenUsecase) RefreshTokenValid(c echo.Context) error {
	tokenString := ox.ExtractTokenString(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func (ox tokenUsecase) ExtractTokenString(c echo.Context) string {
	keys := c.Request().URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	bearToken := c.Request().Header.Get("Authorization")
	if len(strings.Split(bearToken, " ")) == 2 {
		return strings.Split(bearToken, " ")[1]
	}

	return ""
}

func (ox tokenUsecase) ExtractTokenResponse(c echo.Context) (resp model.TokenResponse, err error) {
	tokenString := ox.ExtractTokenString(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println(err)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		log.Println(err)
		err = errors.New("[repository][token][ExtractTokenResponse] while parse jwt")
		return resp, err
	}

	jwtClaims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID := jwtClaims["access_uuid"]
		nasabahID := jwtClaims["nasabah_id"]
		fullname := jwtClaims["fullname"]
		email := jwtClaims["email"]

		resp.AccessUUID = fmt.Sprintf("%s", accessUUID)
		resp.NasabahID = tools.StringToInt64(fmt.Sprintf("%.0f", nasabahID))
		resp.Fullname = fmt.Sprintf("%s", fullname)
		resp.Email = fmt.Sprintf("%s", email)

		return resp, err
	}

	return resp, err
}

func (ox tokenUsecase) ExtractRefreshTokenResponse(c echo.Context) (resp model.TokenResponse, err error) {
	tokenString := ox.ExtractTokenString(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println(err)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		log.Println(err)
		err = errors.New("[repository][token][ExtractTokenResponse] while parse jwt")
		return resp, err
	}

	jwtClaims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID := jwtClaims["refresh_uuid"]
		nasabahID := jwtClaims["nasabah_id"]
		fullname := jwtClaims["fullname"]
		email := jwtClaims["email"]

		resp.RefreshUUID = fmt.Sprintf("%s", refreshUUID)
		resp.NasabahID = tools.StringToInt64(fmt.Sprintf("%.0f", nasabahID))
		resp.Fullname = fmt.Sprintf("%s", fullname)
		resp.Email = fmt.Sprintf("%s", email)

		return resp, err
	}

	return resp, err
}

func (ox tokenUsecase) RefreshToken(c echo.Context) (resp model.NasbahLoginResponses, err error) {
	var (
		tokenInfo model.CreateTokenArgs
	)

	tokenString := ox.ExtractTokenString(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println(err)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		log.Println(err)
		err = errors.New("[usecase][token][RefreshToken] while parse token")
		return resp, err
	}

	//check token valid or not
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		err = errors.New("Invalid Token")
		return resp, err
	}

	// if token valid, claims all info
	jwtClaims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid := jwtClaims["refresh_uuid"]
		nasabahID := jwtClaims["nasabah_id"]
		fullname := jwtClaims["fullname"]
		email := jwtClaims["email"]

		resp.RefreshTokenUUID = fmt.Sprintf("%s", refreshUuid)
		tokenInfo.NasabahID = tools.StringToInt64(fmt.Sprintf("%.0f", nasabahID))
		tokenInfo.Fullname = fmt.Sprintf("%s", fullname)
		tokenInfo.Email = fmt.Sprintf("%s", email)
	}

	//create new token here
	newToken, err := ox.tokenRepo.CreateToken(tokenInfo)
	if err != nil {
		log.Println(err)
		err = errors.New("[usecase][token][RefreshToken] while CreateToken")
		return resp, err
	}

	resp.AccessToken = newToken.AccessToken
	resp.RefreshToken = newToken.RefreshToken

	return resp, err
}
