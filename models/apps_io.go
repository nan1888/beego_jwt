package models

// RegisterForm definiton.
type ChangeappsForm struct {
	Id        string `form:"id"`
	Appname   string `form:"appname"`
	Appremark string `form:"appremark"`
	Token     string `form:"token"`
}

type AddappsForm struct {
	Appname   string `form:"appname"`
	Appremark string `form:"appremark"`
	Token     string `form:"token"`
}

type DeleteappsForm struct {
	Id    string `form:"id"`
	Token string `form:"token"`
}
