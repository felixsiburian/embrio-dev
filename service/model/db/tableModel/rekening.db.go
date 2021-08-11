package tableModel

import "time"

type (
	Rekening struct {
		ID           int64     `gorm:"primary_key;auto_increment" json:"id"`
		NasabahID    int64     `json:"nasabah_id"`
		NoRekening   string    `json:"no_rekening"`
		CabangBank   string    `json:"cabang_bank"`
		Saldo        int64     `json:"saldo"`
		IsActive     bool      `json:"is_active"`
		CreatedBy    string    `json:"created_by"`
		CreatedDate  time.Time `json:"created_date"`
		ModifiedBy   string    `json:"modified_by"`
		ModifiedDate time.Time `json:"modified_date"`
		DeletedBy    string    `json:"deleted_by"`
		DeletedDate  time.Time `json:"deleted_date"`
	}

	TopUpRekeningArgs struct {
		NasabahID  int64  `json:"nasabah_id"`
		NoRekening string `json:"no_rekening"`
		Amount     int64  `json:"amount"`
		OperatedBy string `json:"operated_by"`
	}

	GetSaldoNasabah struct {
		Fullname    string
		NoRekening  string
		CabangBank  string
		Saldo       int64
		IsActive    bool
		CreatedDate time.Time
	}

	TarikTunaiArgs struct {
		NasabahID   int64     `json:"nasabah_id"`
		NoRekening  string    `json:"no_rekening"`
		Amount      int64     `json:"amount"`
		OperatedBy  string    `json:"operated_by"`
		CreatedDate time.Time `json:"created_date"`
	}
)
