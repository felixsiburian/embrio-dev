package repository

import (
	db2 "embrio-dev/lib/db"
	"embrio-dev/service"
	"embrio-dev/service/model/db/tableModel"
	"embrio-dev/service/model/econst"
	"errors"
	"time"
)

type nasabahRepository struct {
}

func NewNasabahRepository(toolRepo service.IToolsRepository) service.INasabahRepository {
	return nasabahRepository{}
}

func (n nasabahRepository) CreateNasabah(args tableModel.Nasabah) (err error) {
	db := db2.ConnectionGorm()

	args.IsActive = true
	args.CreatedBy = econst.AppName
	args.ModifiedBy = econst.AppName
	args.CreatedAt = time.Now()
	args.ModifiedAt = time.Now()

	res := db.Debug().Create(&args)
	if res.Error != nil {
		err = errors.New("[repository][nasabah] while CreateNasabah")
		return err
	}

	defer db.Close()
	return err
}
