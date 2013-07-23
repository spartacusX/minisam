package main

import (
	"github.com/astaxie/beego"
	"github.com/spartacusX/minisam/controllers"
	"github.com/spartacusX/minisam/controllers/admin"
)

func main() {
	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/new", &admin.AddController{})

	beego.Run()
}
