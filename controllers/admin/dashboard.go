package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	conf "github.com/spartacusX/minisam/config"
	m "github.com/spartacusX/minisam/models"
	"html/template"
	"strconv"
	"strings"
)

const DASHBOARD = `    
    var chart1 = new Highcharts.Chart({
        chart: {
            renderTo: 'xxContainer1',
            type: 'bar',
            plotBackgroundColor: null,
            plotBorderWidth: null,
            plotShadow: false,
            backgroundColor:'transparent'
        },
        title: {
            text: 'Software license/installations statistic'
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

type DashBoardController struct {
	beego.Controller
}

func (this *DashBoardController) Prepare() {

}

func (this *DashBoardController) Get() {
	m.ListCounter()

	var sws = m.Statistic()
	total := m.TotalCount(&sws)

	strDash := strings.Replace(DASHBOARD, "{{.EntitleRate}}",
		strconv.FormatFloat((float64)(sws.Entitlements*100/total), 'f', 1, 64), 1)
	strDash = strings.Replace(strDash, "{{.LicenseRate}}",
		strconv.FormatFloat((float64)(sws.Licenses*100/total), 'f', 1, 64), 1)
	strDash = strings.Replace(strDash, "{{.InstallationRate}}",
		strconv.FormatFloat((float64)(sws.Installations*100/total), 'f', 1, 64), 1)
	strDash = strings.Replace(strDash, "{{.UnUsedInstallationRate}}",
		strconv.FormatFloat((float64)(sws.UnUsedInstallations*100/total), 'f', 1, 64), 1)

	this.Data["Dashboard"] = template.JS(strDash)
	this.Data["DashboardHeight"] = conf.DashboardHeight
	this.Data["DashboardWidth"] = conf.DashboardWidth

	this.Layout = "layout.html"
	this.TplNames = "dashboard.tpl"
}

func (this *DashBoardController) Post() {
	fmt.Println("Request Method: Post")
	//数据处理

	this.Ctx.Redirect(302, "/dashboard")
}
