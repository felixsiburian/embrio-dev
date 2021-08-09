package usecase

import (
	"embrio-dev/service"
	"embrio-dev/service/model/db/tableModel"
	"errors"
	"log"
)

type rekeningUsecase struct {
	rekeningRepo service.IRekeningRepository
}

func NewRekeningUsecase(
	rekeningRepo service.IRekeningRepository,
) service.IRekeningUsecase {
	return rekeningUsecase{
		rekeningRepo: rekeningRepo,
	}
}

func (r rekeningUsecase) CreateRekening(args tableModel.Rekening) (err error) {
	if args.NasabahID <= 0 {
		err = errors.New("invalid nasabah id")
		return err
	}

	if len(args.CabangBank) <= 0 {
		err = errors.New("invalid cabang bank")
		return err
	}

	if len(args.NoRekening) < 12 {
		err = errors.New("invalid no rekening")
		return err
	}

	if args.Saldo < 200000 {
		err = errors.New("invalid saldo")
		return err
	}

	err = r.rekeningRepo.CreateRekening(args)
	if err != nil {
		log.Println(err)
		err = errors.New("failed create rekening")
		return err
	}

	return err
}

func (r rekeningUsecase) GetSaldoNasabah(nasabahID int64, noRekening string) (res tableModel.GetSaldoNasabah, err error) {
	res, err = r.rekeningRepo.GetNasabahSaldo(nasabahID, noRekening)
	if err != nil {
		log.Println(err)
		err = errors.New("failed get saldo")
		return res, err
	}

	return res, err
}

func (r rekeningUsecase) TopUpSaldoNasabah(args tableModel.TopUpRekeningArgs) (err error) {
	if len(args.NoRekening) < 12 {
		err = errors.New("invalid no rekening")
		return err
	}

	if args.Amount <= 0 {
		err = errors.New("invalid amount")
		return err
	}

	err = r.rekeningRepo.TopUpSaldo(args)
	if err != nil {
		err = errors.New("Failed Top Up Saldo")
		return err
	}

	return err
}
