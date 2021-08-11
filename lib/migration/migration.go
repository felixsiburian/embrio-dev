// create table here
// Add this function to main.go if you want to create new table

package migration

import (
	db2 "embrio-dev/lib/db"
	"embrio-dev/service/model/db/tableModel"
	"errors"
)

func InitTable() (err error) {
	db := db2.ConnectionGorm()

	res := db.AutoMigrate(&tableModel.Nasabah{}, &tableModel.Rekening{}, &tableModel.History{})
	if res.Error != nil {
		err = errors.New("[migration][initTable] while migrate table")
		return err
	}

	defer db.Close()

	return err
}
