package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	m "github.com/spartacusX/minisam/models"
	"html/template"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

const RecordNumPerPage = 20

const DASHBOARD = `    
    var chart1 = new Highcharts.Chart({
        chart: {
            renderTo: 'xxContainer1',
            type: 'bar',
            plotBackgroundColor: null,
            plotBorderWidth: null,
            plotShadow: false
        },
        title: {
            text: 'Browser market shares at a specific website, 2010'
        },
        tooltip: {
          pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: false
                },
                showInLegend: true
            }
        },
        series: [{
            type: 'pie',
            name: 'Browser share',
            data: [
                ['EntitleRate',   {{.EntitleRate}}],
                {
                    name: 'LicenseRate',
                    y: {{.LicenseRate}},
                    sliced: true,
                    selected: true
                },
                ['InstallationRate',   {{.InstallationRate}}],
                ['UnUsedInstallationRate',     {{.UnUsedInstallationRate}}]
            ]
        }]
    });`

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
	resultToShow = counterList[pageNum*RecordNumPerPage : (pageNum+1)*RecordNumPerPage]

	this.Data["CountResult"] = resultToShow

	var sws = m.Statistic()
	total := m.TotalCount(&sws)
	this.Data["TotalRecords"] = len(counterList)
	this.Data["MinPages"] = (int)(len(counterList) / RecordNumPerPage)
	// this.Data["EntitleRate"] = sws.Entitlements * 100 / total
	// this.Data["LicenseRate"] = sws.Licenses * 100 / total
	// this.Data["InstallationRate"] = sws.Installations * 100 / total
	// this.Data["UnUsedInstallationRate"] = sws.UnUsedInstallations * 100 / total

	var strPageHTML string
	for i := 0; i < (len(counterList) / RecordNumPerPage); i++ {
		strPageHTML += `<li><a href="/?PageNum=` + strconv.Itoa(i+1) + `">` + strconv.Itoa(i+1) + `</a></li>`
	}
	this.Data["Test"] = template.HTML(strPageHTML)

	strDash := strings.Replace(DASHBOARD, "{{.EntitleRate}}",
		strconv.FormatFloat((float64)(sws.Entitlements*100/total), 'f', 1, 64), 1)
	strDash = strings.Replace(strDash, "{{.LicenseRate}}",
		strconv.FormatFloat((float64)(sws.Licenses*100/total), 'f', 1, 64), 1)
	strDash = strings.Replace(strDash, "{{.InstallationRate}}",
		strconv.FormatFloat((float64)(sws.Installations*100/total), 'f', 1, 64), 1)
	strDash = strings.Replace(strDash, "{{.UnUsedInstallationRate}}",
		strconv.FormatFloat((float64)(sws.UnUsedInstallations*100/total), 'f', 1, 64), 1)

	this.Data["Dashboard"] = template.JS(strDash)
	//fmt.Println(this.Data)

	this.Layout = "layout.html"
	this.TplNames = "counterlist.tpl"
}
