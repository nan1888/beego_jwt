package models

// RegisterForm definiton.
type ChangeroleForm struct {
	Id     string `form:"id"`
	Appid  string `form:"appid"`
	Zoneid string `form:"zoneid"`
	Token  string `form:"token"`
}

type AddroleForm struct {
	Appid  string `form:"appid"`
	Zoneid string `form:"zoneid"`
	Token  string `form:"token"`
}

type DeleteroleForm struct {
	Id    string `form:"id"`
	Token string `form:"token"`
}

type ChangezoneForm struct {
	Id         string `form:"id"`
	Zonename   string `form:"username"`
	Zoneremark string `form:"zoneremark"`
	Token      string `form:"token"`
}

type AddzoneForm struct {
	Zonename   string `form:"zonename"`
	Zoneremark string `form:"zoneremark"`
	Token      string `form:"token"`
}

type DeletezoneForm struct {
	Id    string `form:"id"`
	Token string `form:"token"`
}
