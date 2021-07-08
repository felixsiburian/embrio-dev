// all interface for usecase register here

package service

import (
	"embrio-dev/service/model"
)

type INasabahUsecase interface {
	CreateNewUser(args model.NasabahRegisterRequest) (err error)
}
