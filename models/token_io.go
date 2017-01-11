package models

// RegisterForm definiton.
type CreateTokenForm struct {
	Appid  string `form:"appid"`
	Secret string `form:"secret"`
}
