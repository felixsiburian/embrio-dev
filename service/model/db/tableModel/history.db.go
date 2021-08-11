package tableModel

import "time"

type History struct {
	HistoryID       int64     `gorm:"primary_key;auto_increment" json:"history_id"`
	NasabahID       int64     `json:"nasabah_id"`
	NoRekening      string    `json:"no_rekening"`
	JumlahTransaksi int64     `json:"jumlah_transaksi"`
	Action          string    `json:"action"`
	IsActive        bool      `json:"is_active"`
	CreatedBy       string    `json:"created_by"`
	CreatedDate     time.Time `json:"created_date"`
	ModifiedBy      string    `json:"modified_by"`
	ModifiedDate    time.Time `json:"modified_date"`
	DeletedBy       string    `json:"deleted_by"`
	DeletedDate     time.Time `json:"deleted_date"`
}
