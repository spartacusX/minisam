package main

import (
	//"fmt"
	"github.com/astaxie/beego"
	conf "github.com/spartacusX/minisam/config"
	"github.com/spartacusX/minisam/controllers"
	"github.com/spartacusX/minisam/controllers/admin"
	//"github.com/spartacusX/minisam/dbschema"
)

func main() {
	//beego.SetStaticPath("/myjs", "static/myjs")

	// t := dbschema.GetTableSchemaBySqlName("amSoftLicCounter")
	// fmt.Println(t)

	beego.SessionOn = true

	conf.AppConfigPath = "./conf/sam.conf"
	conf.ParseConfig()

	beego.RegisterController("/index", &controllers.MainController{})
	//beego.RegisterController("/new", &admin.AddController{})
	beego.RegisterController("/dashboard", &admin.DashBoardController{})
	//beego.RegisterController("/unlock", &admin.UnlockController{})

	beego.Run()
}
