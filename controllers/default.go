package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/spartacusX/minisam/dbschema"
	m "github.com/spartacusX/minisam/models"
	"html/template"
	"strconv"
	//"strings"
)

type MainController struct {
	beego.Controller
}

/*****************************************************************************/
/* The request can be divided into two categories, list and detail
/* The two kind of requests distinguished by with parameter counterid or not.
/* 1. For list, serve the request content content as html
/* 2. for detail, serve the request content as json or xml
/****************************************************************************/
func (this *MainController) Get() {

	counterId := this.GetString("CounterId")
	if counterId == "" { // Invalid counterId, means list request
		ServeCounterList(this)
	} else { // Got counterId, means detail request
		ServeCounterDetail(counterId, this)
	}
}

func ServeCounterList(p *MainController) {
	var counters m.SWCounterList
	if p.GetSession("counterList") == nil { // Check session firstly
		counters, _ = m.ListCounter()
		p.SetSession("counterList", counters)
	} else {
		counters = p.GetSession("counterList").(m.SWCounterList)
	}

	//beego.Debug(counters)

	// Pagination
	var currentPage int
	value := p.GetSession("currentPage")
	if value != nil { // Got current page index from session
		currentPage = value.(int)
	} else {
		currentPage = -1
	}
	pageNum, err := p.GetInt32("PageNum")
	if err != nil { // Without parameter "PageNum"
		pageNum = 1 // goto first page
	} else if pageNum == -1 { // goto previous page
		if currentPage != -1 {
			pageNum = currentPage - 1
		}
	} else if pageNum == -2 { // goto next page
		if currentPage != -2 {
			pageNum = currentPage + 1
		}
	}

	totalPages := counters.TotalPages()
	if pageNum < 1 {
		pageNum = 1
	}
	if pageNum > totalPages {
		pageNum = totalPages
	}

	p.SetSession("currentPage", pageNum)

	p.Data["CountResult"] = counters.GetPageByIndex(pageNum)

	RenderListTable(p)
	RenderPagination(p, totalPages)
	RenderDetailInformation(p)
	RenderDetailGeneral(p)

	p.Layout = "layout.html"
	p.TplNames = "counterlist.tpl"
}

func ServeCounterDetail(counterId string, p *MainController) {
	counterDetail, _ := m.QueryCounterDetailById(counterId)
	p.Data["xml"] = counterDetail
	p.ServeXml()
}

type ColumnList struct {
	ColumnName     string
	ColumnClassAtt string
}

func RenderListTable(p *MainController) {
	columnList := []ColumnList{{"", ""}, {"Id", "hide"}, {"Name", ""}, {"Rights", ""},
		{"Entitlements", ""}, {"Installs", ""}, {"UnusedInstalls", ""}}

	p.Data["ColumnList"] = columnList
}

func RenderPagination(p *MainController, totalPages int) {
	var strPageHTML string
	for i := 0; i < totalPages; i++ {
		strPageHTML += `<li><a href="/index?PageNum=` + strconv.Itoa(i+1) + `">` + strconv.Itoa(i+1) + `</a></li>`
	}
	p.Data["Pagination"] = template.HTML(strPageHTML)
}

type DetailInformation struct {
	SqlName string
	Desc    string
	Id      string
}

func RenderDetailInformation(p *MainController) {
	swCounterTable, _ := dbschema.GetTableSchemaBySqlName("amSoftLicCounter")
	//fmt.Println(swCounterTable)

	labels := []string{"Name", "Code", "Type", "Context", "Supervisor", "Definition", "Model"}
	inputInfo := make([]DetailInformation, len(labels))
	for i, v := range labels {
		val, err := swCounterTable.GetFieldByName(v)
		if err == nil {
			inputInfo[i].SqlName = val.SqlName
			inputInfo[i].Desc = val.Desc + ":"
		} else {
			val, _ := swCounterTable.GetLinkByName(v)
			inputInfo[i].SqlName = val.SqlName
			inputInfo[i].Desc = val.Desc + ":"
		}

		inputInfo[i].Id = "DetailInformation_" + v
	}

	p.Data["DetailInformationInput"] = inputInfo

	checkboxes := []string{"bType", "bFamily", "bInternal", "bAutomated", "bCountLic", "bCountEnt", "bCountInst", "bLicUpgrade"}
	checkboxInfo := make([]DetailInformation, len(checkboxes))
	for i, v := range checkboxes {
		val, err := swCounterTable.GetFieldByName(v)
		if err == nil {
			checkboxInfo[i].SqlName = val.SqlName
			checkboxInfo[i].Desc = val.Desc
			checkboxInfo[i].Id = "DetailInformation_" + checkboxInfo[i].SqlName
		}
	}

	p.Data["DetailInformationCheckBox"] = checkboxInfo
}

func RenderDetailGeneral(p *MainController) {
	swCounterTable, _ := dbschema.GetTableSchemaBySqlName("amSoftLicCounter")
	labels := []string{"dLicUseRights", "dSoftInstallCount", "dUnusedInstall"}
	generalInfo := make([]DetailInformation, len(labels))
	for i, v := range labels {
		val, err := swCounterTable.GetFieldByName(v)
		if err == nil {
			generalInfo[i].SqlName = val.SqlName
			generalInfo[i].Desc = val.Desc
			generalInfo[i].Id = "DetailGeneralTab_" + generalInfo[i].SqlName
		}
	}
	//fmt.Println(generalInfo)
	compliance := DetailInformation{"Compliance", "Compliance", "DetailGeneralTab_Compliance"}
	generalInfo = append(generalInfo, compliance)
	//fmt.Println(generalInfo)
	p.Data["DetailGeneralTab"] = generalInfo
}
