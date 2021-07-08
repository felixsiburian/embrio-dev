// all interface for repository register here

package service

import (
	"embrio-dev/service/model/db/tableModel"
)

type INasabahRepository interface {
	CreateNasabah(args tableModel.Nasabah) (err error)
}

type IToolsRepository interface {
	SaltAndHash(pwd string) (s string, err error)
	CheckEmailValidation(email string) (res bool, err error)
}
