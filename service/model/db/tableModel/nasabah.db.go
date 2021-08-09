package tableModel

import "time"

type (
	Nasabah struct {
		NasabahID    int64     `gorm:"primary_key;auto_increment" json:"nasabah_id"`
		Username     string    `json:"username"`
		Fullname     string    `json:"fullname"`
		Email        string    `json:"email"`
		PhoneNumber  string    `json:"phone_number"`
		Alamat       string    `json:"alamat"`
		Pin          string    `json:"pin"`
		Password     string    `json:"password"`
		IsFirstLogin bool      `json:"is_first_login"`
		IsActive     bool      `json:"is_active"`
		CreatedBy    string    `json:"created_by"`
		CreatedDate  time.Time `json:"created_date"`
		ModifiedBy   string    `json:"modified_by"`
		ModifiedDate time.Time `json:"modified_date"`
		DeletedBy    string    `json:"deleted_by"`
		DeletedDate  time.Time `json:"deleted_date"`
	}

	GetNasabahList struct {
		Username    string
		Fullname    string
		Email       string
		PhoneNumber string
		Alamat      string
		CreatedDate time.Time
	}
)
