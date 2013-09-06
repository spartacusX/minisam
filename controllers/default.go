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
	var counterList []m.SoftWareCounter

	if this.GetSession("counterList") == nil {
		counterList, _ = m.CounterList()
		this.SetSession("counterList", counterList)
	} else {
		counterList = this.GetSession("counterList").([]m.SoftWareCounter)
	}

	//fmt.Println(counterList)

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

	counterId, _ := this.GetInt("CounterId")
	if counterId != 0 {
		for _, v := range counterList {
			if (int)(counterId) == v.LCounterId {
				this.Data["SelectedCounter"] = v
				break
			}

		}
	}

	this.Data["TotalRecords"] = len(counterList)

	var strPageHTML string
	for i := 1; i <= (len(counterList) / RecordNumPerPage); i++ {
		strPageHTML += `<li><a href="/index?PageNum=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
	}
	this.Data["Test"] = template.HTML(strPageHTML)

	this.Layout = "layout.html"
	this.TplNames = "counterlist.tpl"
}
