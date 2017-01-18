package models

// RegisterForm definiton.
type ChangeroleForm struct {
	Id     string `form:"id"`
	Appid  string `form:"appid"`
	Zoneid string `form:"zoneid"`
}

type AddroleForm struct {
	Rolename   string `form:"zonename"`
	Roleremark string `form:"zoneremark"`
}

type DeleteroleForm struct {
	Id string `form:"id"`
}

type ChangezoneForm struct {
	Id         string `form:"id"`
	Zonename   string `form:"username"`
	Zoneremark string `form:"zoneremark"`
}

type AddzoneForm struct {
	Zonename   string `form:"zonename"`
	Zoneremark string `form:"zoneremark"`
}

type DeletezoneForm struct {
	Id string `form:"id"`
}
