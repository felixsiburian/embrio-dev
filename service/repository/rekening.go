package repository

import (
	db2 "embrio-dev/lib/db"
	"embrio-dev/service"
	"embrio-dev/service/model/db/tableModel"
	"embrio-dev/service/repository/queries"
	"errors"
	"log"
	"time"
)

type rekeningRepository struct {
	nasabahRepo service.INasabahRepository
}

func NewRekeningRepository(nasabahRepo service.INasabahRepository) service.IRekeningRepository {
	return rekeningRepository{
		nasabahRepo: nasabahRepo,
	}
}

func (r rekeningRepository) TopUpRekening(args tableModel.TopUpRekeningArgs) (err error) {
	db := db2.ConnectionGorm()

	defer db.Close()
	return err
}

func (r rekeningRepository) CreateRekening(args tableModel.Rekening) (err error) {
	db := db2.ConnectionGorm()

	nasabahInfo, err := r.nasabahRepo.GetNasabahInfo(args.NasabahID)
	if err != nil {
		err = errors.New("while get nasabah info")
		return err
	}

	args.CreatedBy = nasabahInfo.Fullname
	args.CreatedDate = time.Now()
	args.IsActive = true

	resp := db.Debug().Create(&args)
	if resp.Error != nil {
		log.Println(resp.Error)
		err = errors.New("while create rekening")
		return err
	}

	if resp.RowsAffected <= 0 {
		err = errors.New("failed to create rekening")
		return err
	}

	defer db.Close()
	return err
}

func (r rekeningRepository) GetNasabahSaldo(nasabahID int64, noRekening string) (res tableModel.GetSaldoNasabah, err error) {
	db := db2.ConnectionGorm()

	resp := db.Debug().Raw(queries.QueryGetSaldoNasabah, nasabahID, noRekening).Scan(&res)
	if resp.Error != nil {
		log.Println(resp.Error)
		err = errors.New("while get rekening")
		return res, err
	}

	defer db.Close()
	return res, err
}

func (r rekeningRepository) TopUpSaldo(args tableModel.TopUpRekeningArgs) (err error) {
	var (
		db                 = db2.ConnectionGorm()
		transactionHistory tableModel.History
	)

	nasabahSaldo, err := r.GetNasabahSaldo(args.NasabahID, args.NoRekening)
	if err != nil {
		err = errors.New("while get rekening")
		return err
	}

	amount := nasabahSaldo.Saldo + args.Amount
	resp := db.Debug().Exec(queries.QueryUpdateSaldo, amount, args.OperatedBy, args.NasabahID, args.NoRekening)
	if resp.Error != nil {
		log.Println(resp.Error)
		err = errors.New("while update rekening")
		return err
	}

	transactionHistory.NasabahID = args.NasabahID
	transactionHistory.NoRekening = args.NoRekening
	transactionHistory.JumlahTransaksi = args.Amount
	transactionHistory.Action = "Incoming Cash"
	transactionHistory.CreatedBy = args.OperatedBy
	transactionHistory.CreatedDate = time.Now()
	resps := db.Debug().Create(&transactionHistory)
	if resps.Error != nil {
		log.Println(resp.Error)
		err = errors.New("while create transaction history")
		return err
	}

	defer db.Close()
	return err
}

func (r rekeningRepository) DeductionRekening(args tableModel.TopUpRekeningArgs) (err error) {
	var (
		db = db2.ConnectionGorm()
	)

	nasabahSaldo, err := r.GetNasabahSaldo(args.NasabahID, args.NoRekening)
	if err != nil {
		err = errors.New("while get rekening")
		return err
	}

	if nasabahSaldo.Saldo < args.Amount {
		err = errors.New("your balance is not enough")
		log.Println(err)
		return err
	}

	amount := nasabahSaldo.Saldo - args.Amount
	resp := db.Debug().Exec(queries.QueryUpdateSaldo, amount, args.OperatedBy, args.NasabahID, args.NoRekening)
	if resp.Error != nil {
		log.Println(resp.Error)
		err = errors.New("while update rekening")
		return err
	}

	defer db.Close()
	return err
}

func (r rekeningRepository) TarikSaldo(args tableModel.TarikTunaiArgs) (err error) {
	var (
		db                 = db2.ConnectionGorm()
		deductionArgs      tableModel.TopUpRekeningArgs
		transactionHistory tableModel.History
	)

	nasabahRek, err := r.GetNasabahSaldo(args.NasabahID, args.NoRekening)
	if err != nil {
		err = errors.New("while get rekening")
		return err
	}

	if nasabahRek.Saldo < args.Amount {
		err = errors.New("your balance is not enough")
		log.Println(err)
		return err
	}

	deductionArgs.Amount = args.Amount
	deductionArgs.NoRekening = args.NoRekening
	deductionArgs.NasabahID = args.NasabahID
	deductionArgs.OperatedBy = args.OperatedBy
	err = r.DeductionRekening(deductionArgs)
	if err != nil {
		err = errors.New("while deduction rekening")
		return err
	}

	transactionHistory.NasabahID = args.NasabahID
	transactionHistory.NoRekening = args.NoRekening
	transactionHistory.JumlahTransaksi = args.Amount * -1
	transactionHistory.Action = "Cash Out"
	transactionHistory.CreatedBy = args.OperatedBy
	transactionHistory.CreatedDate = time.Now()
	resp := db.Debug().Create(&transactionHistory)
	if resp.Error != nil {
		log.Println(resp.Error)
		err = errors.New("while create transaction history")
		return err
	}

	defer db.Close()
	return err
}
