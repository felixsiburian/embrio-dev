// all interface for repository register here

package service

import (
	"embrio-dev/service/model"
	"embrio-dev/service/model/db/tableModel"
)

type INasabahRepository interface {
	SignIn(args model.NasbahLoginRequest) (resp model.NasbahLoginResponses, err error)
	CreateNasabah(args tableModel.Nasabah) (err error)
	GetNasabahInfo(nasabahID int64) (res tableModel.GetNasabahList, err error)
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

type IRekeningRepository interface {
	TopUpRekening(args tableModel.TopUpRekeningArgs) (err error)
	DeductionRekening(args tableModel.TopUpRekeningArgs) (err error)
	CreateRekening(args tableModel.Rekening) (err error)
	GetNasabahSaldo(nasabahID int64, noRekening string) (res tableModel.GetSaldoNasabah, err error)
	TopUpSaldo(args tableModel.TopUpRekeningArgs) (err error)
	TarikSaldo(args tableModel.TarikTunaiArgs) (err error)
}
