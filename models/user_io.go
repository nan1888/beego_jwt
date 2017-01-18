package models

// RegisterForm definiton.
type RegisterForm struct {
	Email    string `form:"email"    valid:"Required"`
	Password string `form:"password" valid:"Required"`
	Username string `form:"username"`
}

// LoginForm definiton.
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password" valid:"Required"`
}
type LoginInfo struct {
	Code     int   `json:"code"`
	UserInfo *User `json:"user"`
}
type ChangeuserroleForm struct {
	Id       string `form:"id"`
	Username string `form:"username"`
}
