package model

type NasabahRegisterRequest struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Alamat      string `json:"alamat"`
	Pin         string `json:"pin"`
	Password    string `json:"password"`
}

type NasbahLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NasbahLoginResponses struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	AccessTokenExpires  int64  `json:"token_expires"`
	RefreshTokenExpires int64  `json:"refresh_expires"`
	AccessTokenUUID     string `json:"access_uuid"`
	RefreshTokenUUID    string `json:"refresh_uuid"`
	// access and refresh UUID is unique so, user can create morethan one token.
	//	this is needed when user want to logged in with different devices.
	// user can also logout from any of the devices without them being logged out from all devices
}
