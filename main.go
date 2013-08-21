package main

import (
	"github.com/astaxie/beego"
	"github.com/spartacusX/minisam/controllers"
	"github.com/spartacusX/minisam/controllers/admin"
)

func main() {
	//beego.SetStaticPath("/myjs", "static/myjs")

	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/new", &admin.AddController{})
	beego.RegisterController("/dashboard", &admin.DashBoardController{})
	beego.RegisterController("/unlock", &admin.UnlockController{})

	beego.Run()
}
