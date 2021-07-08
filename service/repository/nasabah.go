package repository

import (
	db2 "embrio-dev/lib/db"
	"embrio-dev/service"
	"embrio-dev/service/model/db/tableModel"
	"embrio-dev/service/model/econst"
	"embrio-dev/service/repository/queries"
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

	getUserByQuery := db.Exec(queries.QueryGetUserByEmail, args.Email).RowsAffected
	if getUserByQuery > 0 {
		err = errors.New("[repository][nasabah][CreateNasabah] nasabah already exists")
		return err
	}

	res := db.Debug().Create(&args)
	if res.Error != nil {
		err = errors.New("[repository][nasabah][CreateNasabah] while CreateNasabah")
		return err
	}

	defer db.Close()
	return err
}
