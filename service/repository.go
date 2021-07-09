// all interface for repository register here

package service

import (
	"embrio-dev/service/model"
	"embrio-dev/service/model/db/tableModel"
)

type INasabahRepository interface {
	SignIn(args model.NasbahLoginRequest) (resp model.NasbahLoginResponses, err error)
	CreateNasabah(args tableModel.Nasabah) (err error)
}

type IToolsRepository interface {
	SaltAndHash(pwd string) (s string, err error)
	CheckEmailValidation(email string) (res bool, err error)
	VerifyPassword(hashedPwd, pwd string) (err error)
	GUID() string
}

type ITokenRepository interface {
	CreateToken(args model.CreateTokenArgs) (resp model.NasbahLoginResponses, err error)
	//ExtractTokenString(c echo.Context) string
}
