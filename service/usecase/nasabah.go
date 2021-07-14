package usecase

import (
	"embrio-dev/service"
	"embrio-dev/service/model"
	"embrio-dev/service/model/db/tableModel"
	"errors"
	"log"
	"strings"
)

type nasabahUsecase struct {
	nasabahRepo service.INasabahRepository
	toolRepo    service.IToolsRepository
}

func NewNasabahUsecase(nasabahRepo service.INasabahRepository, toolRepo service.IToolsRepository) service.INasabahUsecase {
	return nasabahUsecase{
		nasabahRepo: nasabahRepo,
		toolRepo:    toolRepo,
	}
}

func (n nasabahUsecase) CreateNewUser(args model.NasabahRegisterRequest) (err error) {
	var (
		hashedPass        string
		hashedPin         string
		createNasabahArgs tableModel.Nasabah
	)

	// TODO : Validate nasabah request here
	{
		if strings.TrimSpace(args.Password) != "" {
			hashedPass, err = n.toolRepo.SaltAndHash(args.Password)
			if err != nil {
				log.Print(err)
				err = errors.New("[usecase][nasabah] while hashed password")
				return err
			}
		}

		if strings.TrimSpace(args.Pin) != "" {
			hashedPin, err = n.toolRepo.SaltAndHash(args.Pin)
			if err != nil {
				log.Print(err)
				err = errors.New("[usecase][nasabah] while hashed pin")
				return err
			}
		}

		if len(args.Email) <= 0 {
			err = errors.New("invalid email")
			log.Print(err)
			return err
		}

		emailValidation, err := n.toolRepo.CheckEmailValidation(args.Email)
		if err != nil {
			err = errors.New("while check email validation")
			log.Print(err)
			return err
		}

		if !emailValidation {
			err = errors.New("Email is invalid")
			log.Print(err)
			return err
		}

		if len(args.PhoneNumber) <= 0 {
			err = errors.New("invalid phoneNumber")
			log.Print(err)
			return err
		}

		if len(args.Fullname) <= 0 {
			err = errors.New("invalid fullname")
			log.Print(err)
			return err
		}

		if len(args.Username) <= 0 {
			err = errors.New("invalid username")
			log.Print(err)
			return err
		}

		if len(args.Alamat) <= 0 {
			err = errors.New("invalid alamat")
			log.Print(err)
			return err
		}
	}

	createNasabahArgs.Username = args.Username
	createNasabahArgs.Fullname = args.Fullname
	createNasabahArgs.Email = args.Email
	createNasabahArgs.PhoneNumber = args.PhoneNumber
	createNasabahArgs.Pin = hashedPin
	createNasabahArgs.Password = hashedPass
	createNasabahArgs.Alamat = args.Alamat

	err = n.nasabahRepo.CreateNasabah(createNasabahArgs)
	if err != nil {
		log.Print(err)
		return err
	}

	return err
}

func (n nasabahUsecase) Auth(args model.NasbahLoginRequest) (resp model.NasbahLoginResponses, err error) {
	if len(args.Username) <= 0 {
		err = errors.New("invalid username")
		log.Print(err)
		return resp, err
	}

	if len(args.Password) <= 0 {
		err = errors.New("invalid password")
		log.Print(err)
		return resp, err
	}

	resp, err = n.nasabahRepo.SignIn(args)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	return resp, err
}
