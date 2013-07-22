package main

import (
	"github.com/astaxie/beego"
	"minisam/controllers"
	"minisam/controllers/admin"
)

func main() {
	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/new", &admin.AddController{})

	beego.Run()
}
