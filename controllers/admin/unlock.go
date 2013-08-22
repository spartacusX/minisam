package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	m "github.com/spartacusX/minisam/models"
)

type UnlockController struct {
	beego.Controller
}

func (this *UnlockController) Get() {
	fmt.Println("Request Method: Get")
	this.Layout = "layout.html"
	this.TplNames = "admin/unlock.tpl"
}

func (this *UnlockController) Post() {
	fmt.Println("Request Method: Post")
	//数据处理
	this.Ctx.Request.ParseForm()
	strUser := this.Ctx.Request.Form.Get("UserName")
	beego.Info(this.Ctx.Request.Form)

	m.UnlockAmUser(strUser)

	this.Ctx.Redirect(302, "/")
}
