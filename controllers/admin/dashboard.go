package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	//m "github.com/spartacusX/minisam/models"
	//"strconv"
	//"strings"
)

type DashBoardController struct {
	beego.Controller
}

func (this *DashBoardController) Prepare() {

}

func (this *DashBoardController) Get() {
	fmt.Println("Request Method: Get")

	// var counter m.SoftWareCounter

	// this.Ctx.Request.ParseForm()
	// counter.Name = this.Ctx.Request.Form.Get("Name")
	// counter.LicUseRights, _ = strconv.Atoi(this.Ctx.Request.Form.Get("LicUseRights"))
	// counter.EntCount, _ = strconv.Atoi(this.Ctx.Request.Form.Get("EntCount"))
	// counter.SoftInstallCount, _ = strconv.Atoi(this.Ctx.Request.Form.Get("SoftInstallCount"))
	// counter.UnusedInstall, _ = strconv.Atoi(this.Ctx.Request.Form.Get("UnusedInstall"))
	// beego.Info(this.Ctx.Request.Form)

	// fmt.Println(counter)
	//this.Data["Counter"] = counter
	//this.Layout = "admin/dashboard.html"
	//this.Data["ChartScript"] = "var myPie = new Chart(document.getElementById(" + "canvas" + ").getContext(" + "2d" + ")).Pie(pieData);"
	this.TplNames = "dashboard.html"
}

func (this *DashBoardController) Post() {
	fmt.Println("Request Method: Post")
	//数据处理

	this.Ctx.Redirect(302, "/dashboard")
}
