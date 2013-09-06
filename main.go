package main

import (
	"github.com/astaxie/beego"
	conf "github.com/spartacusX/minisam/config"
	"github.com/spartacusX/minisam/controllers"
	"github.com/spartacusX/minisam/controllers/admin"
)

func main() {
	//beego.SetStaticPath("/myjs", "static/myjs")

	beego.SessionOn = true

	conf.AppConfigPath = "./conf/sam.conf"
	conf.ParseConfig()

	beego.RegisterController("/index", &controllers.MainController{})
	beego.RegisterController("/new", &admin.AddController{})
	beego.RegisterController("/dashboard", &admin.DashBoardController{})
	beego.RegisterController("/unlock", &admin.UnlockController{})

	beego.Run()
}
