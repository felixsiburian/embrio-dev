package model

type CreateTokenArgs struct {
	NasabahID int64
	Fullname  string
	Email     string
}

type TokenResponse struct {
	AccessUUID  string
	RefreshUUID string
	NasabahID   int64
	Fullname    string
	Email       string
}
