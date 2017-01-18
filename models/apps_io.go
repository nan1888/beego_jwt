package models

// RegisterForm definiton.
type ChangeappsForm struct {
	Id        string `form:"id"`
	Appname   string `form:"appname"`
	Appremark string `form:"appremark"`
}

type AddappsForm struct {
	Appname   string `form:"appname"`
	Appremark string `form:"appremark"`
}

type DeleteappsForm struct {
	Id string `form:"id"`
}
