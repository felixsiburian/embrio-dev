package repository

import (
	"embrio-dev/service"
	"embrio-dev/service/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type tokenRepository struct {
	toolRepo service.IToolsRepository
}

func NewTokenRepository(toolRepo service.IToolsRepository) service.ITokenRepository {
	return tokenRepository{
		toolRepo: toolRepo,
	}
}

func (t tokenRepository) CreateToken(args model.CreateTokenArgs) (resp model.NasbahLoginResponses, err error) {
	resp.AccessTokenExpires = time.Now().Add(time.Hour * 24).Unix()
	resp.AccessTokenUUID = t.toolRepo.GUID()

	resp.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	resp.RefreshTokenUUID = t.toolRepo.GUID()

	// Create Access Token
	{
		jwtClaims := jwt.MapClaims{}
		jwtClaims["authorized"] = true
		jwtClaims["access_uuid"] = resp.AccessTokenUUID
		jwtClaims["nasabah_id"] = args.NasabahID
		jwtClaims["fullname"] = args.Fullname
		jwtClaims["email"] = args.Email
		jwtClaims["exp"] = resp.AccessTokenExpires

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
		resp.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			log.Print(err)
			err = errors.New("[repository][token] while signed string token")
			return resp, err
		}
	}

	// Create Refresh Token
	{
		rtClaims := jwt.MapClaims{}
		rtClaims["refresh_uuid"] = resp.RefreshTokenUUID
		rtClaims["nasabah_id"] = args.NasabahID
		rtClaims["fullname"] = args.Fullname
		rtClaims["email"] = args.Email
		rtClaims["exp"] = resp.RefreshTokenExpires
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
		resp.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
		if err != nil {
			log.Print(err)
			err = errors.New("[repository][token] while signed string refresh token")
			return resp, err
		}
	}

	return resp, err
}
