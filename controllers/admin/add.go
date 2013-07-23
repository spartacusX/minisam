package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	m "github.com/spartacusX/minisam/models"
	"strconv"
	//"strings"
)

type AddController struct {
	beego.Controller
}

func (this *AddController) Prepare() {

}

func (this *AddController) Get() {
	fmt.Println("Request Method: Get")
	this.Layout = "layout.html"
	this.TplNames = "admin/newcounter.tpl"
}

func (this *AddController) Post() {
	fmt.Println("Request Method: Post")
	//数据处理
	this.Ctx.Request.ParseForm()
	Id := this.Ctx.Request.Form.Get("Id")
	Model := this.Ctx.Request.Form.Get("Model")
	LicenseNum := this.Ctx.Request.Form.Get("LicenseNum")
	InstallationNum := this.Ctx.Request.Form.Get("InstallationNum")
	beego.Info(this.Ctx.Request.Form)

	var counter m.SWCounter
	counter.Id, _ = strconv.Atoi(Id)
	counter.Model = Model
	counter.LicenseNum, _ = strconv.Atoi(LicenseNum)
	counter.InstallationNum, _ = strconv.Atoi(InstallationNum)
	m.AddCounter(counter)

	this.Ctx.Redirect(302, "/")
}
