package tableModel

import "time"

type Nasabah struct {
	NasabahID   int64     `gorm:"primary_key;auto_increment" json:"nasabah_id"`
	Username    string    `json:"username"`
	Fullname    string    `json:"fullname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Pin         string    `json:"pin"`
	Password    string    `json:"password"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedBy  string    `json:"modified_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	DeletedBy   string    `json:"deleted_by"`
	DeletedAt   time.Time `json:"deleted_at"`
}
