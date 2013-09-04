package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	m "github.com/spartacusX/minisam/models"
	"html/template"
	"strconv"
	//"strings"
)

type MainController struct {
	beego.Controller
}

const RecordNumPerPage = 15

func (this *MainController) Get() {
	var resultToShow interface{}
	counterList, _ := m.CounterList()

	pageNum, _ := this.GetInt("PageNum")
	if pageNum < 1 {
		pageNum = 1
	}
	pageNum--
	if len(counterList) >= (int)(pageNum+1)*RecordNumPerPage {
		resultToShow = counterList[pageNum*RecordNumPerPage : (pageNum+1)*RecordNumPerPage]
	} else {
		resultToShow = counterList[pageNum*RecordNumPerPage:]
	}

	this.Data["CountResult"] = resultToShow

	//var sws = m.Statistic()
	//total := m.TotalCount(&sws)
	this.Data["TotalRecords"] = len(counterList)

	var strPageHTML string
	for i := 1; i <= (len(counterList) / RecordNumPerPage); i++ {
		strPageHTML += `<li><a href="/?PageNum=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
	}
	this.Data["Test"] = template.HTML(strPageHTML)

	this.Layout = "layout.html"
	this.TplNames = "counterlist.tpl"
}
