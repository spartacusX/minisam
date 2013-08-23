package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	m "github.com/spartacusX/minisam/models"
)

type MainController struct {
	beego.Controller
}

const RecordNumPerPage = 20

func (this *MainController) Get() {
	// crupkg := models.GetCruPkg(this.Ctx.Params[":pkg"])
	// this.Data["CruPkg"] = crupkg
	// pkglist := models.GetAllPkg()
	// pm := make([]map[string]interface{}, len(pkglist))
	// for _, pk := range pkglist {
	// 	m := make(map[string]interface{}, 2)
	// 	m["PKG"] = pk
	// 	if this.Ctx.Params[":pkg"] == pk.Pathname {
	// 		m["Cru"] = true
	// 	} else {
	// 		m["Cru"] = false
	// 	}
	// 	pm = append(pm, m)
	// }
	// this.Data["PkgList"] = pm
	// if crupkg.Id == 0 {
	// 	this.Data["Content"] = welcome //template.HTML(string(blackfriday.MarkdownCommon([]byte(welcome))))
	// } else {
	// 	at := models.GetArticle(crupkg.Id)
	// 	this.Data["Content"] = at.Content //template.HTML(string(blackfriday.MarkdownCommon([]byte(at.Content))))
	// }

	var resultToShow interface{}
	counterList, _ := m.CounterList()

	pageNum, _ := this.GetInt("PageNum")
	if pageNum < 0 {
		pageNum = 0
	}
	resultToShow = counterList[pageNum*RecordNumPerPage:]

	this.Data["CountResult"] = resultToShow

	var sws = m.Statistic()
	total := m.TotalCount(&sws)
	this.Data["TotalRecords"] = len(counterList)
	this.Data["MinPages"] = (int)(len(counterList) / RecordNumPerPage)
	this.Data["EntitleRate"] = sws.Entitlements * 100 / total
	this.Data["LicenseRate"] = sws.Licenses * 100 / total
	this.Data["InstallationRate"] = sws.Installations * 100 / total
	this.Data["UnUsedInstallationRate"] = sws.UnUsedInstallations * 100 / total

	fmt.Println(this.Data)

	this.Layout = "layout.html"
	this.TplNames = "counterlist.tpl"
}
