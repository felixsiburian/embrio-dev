package repository

import (
	db2 "embrio-dev/lib/db"
	"embrio-dev/service"
	"embrio-dev/service/model"
	"embrio-dev/service/model/db/tableModel"
	"embrio-dev/service/model/econst"
	"embrio-dev/service/repository/queries"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var db = db2.ConnectionGorm()

type nasabahRepository struct {
	toolRepo  service.IToolsRepository
	tokenRepo service.ITokenRepository
}

func NewNasabahRepository(
	toolRepo service.IToolsRepository,
	tokenRepo service.ITokenRepository,
) service.INasabahRepository {
	return nasabahRepository{
		toolRepo:  toolRepo,
		tokenRepo: tokenRepo,
	}
}

func (n nasabahRepository) CreateNasabah(args tableModel.Nasabah) (err error) {
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

func (n nasabahRepository) SignIn(args model.NasbahLoginRequest) (resp model.NasbahLoginResponses, err error) {
	var (
		nasabah         tableModel.Nasabah
		createTokenArgs model.CreateTokenArgs
	)

	rows := db.Debug().Model(tableModel.Nasabah{}).Where("username = ?", args.Username).Scan(&nasabah)
	if rows.Error != nil {
		log.Println(rows.Error.Error())
		err = errors.New("[repository][nasabah][SignIn] while find user by email")
		return resp, err
	}

	err = n.toolRepo.VerifyPassword(nasabah.Password, args.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Print(err)
		err = errors.New("[repository][nasabah][SignIn] while verify password")
		return resp, err
	}

	createTokenArgs.NasabahID = nasabah.NasabahID
	createTokenArgs.Email = nasabah.Email
	createTokenArgs.Fullname = nasabah.Fullname
	resp, err = n.tokenRepo.CreateToken(createTokenArgs)
	if err != nil {
		log.Print(err)
		err = errors.New("[repository][nasabah][SignIn] while create token")
		return resp, err
	}

	defer db.Close()
	return resp, err
}
