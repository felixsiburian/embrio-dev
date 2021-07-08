package model

type NasabahRegisterRequest struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
	Password    string `json:"password"`
}
