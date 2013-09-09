package models

import (
	_ "code.google.com/p/odbc"
	"fmt"
	conf "github.com/spartacusX/minisam/config"
)

type SWCounterDetail struct {
	BAutomated  bool `xml:"bAutomated"`
	BCountEnt   bool `xml:"bCountEnt"`
	BCountInst  bool `xml:"bCountInst"`
	BCountLic   bool `xml:"bCountLic"`
	BFamily     bool `xml:"bFamily"`
	BInternal   bool `xml:"bInternal"`
	BLicUpgrade bool `xml:"bLicUpgrade"`
	BType       bool `xml:"bType"`
	Code        string
	Context     string
	Name        string
	Type        string
}

func QueryCounterDetailById(counterId string) (counter SWCounterDetail, err error) {
	db, err := sqlConnect()
	if err != nil {
		return counter, err
	}

	defer db.Close()

	strQuery := "SELECT name, Code, Context, bType, bFamily, bInternal, bAutomated, " +
		" bCountLic, bCountEnt, bCountInst, bLicUpgrade" +
		" FROM " + conf.Scheme + "." + TBL_SOFTWARECOUNTER +
		" WHERE lCountId = " + counterId

	rows, err := db.Query(strQuery)
	if err != nil {
		fmt.Printf("db.Query failed. %v\n", err)
		return counter, err
	}

	for rows.Next() {
		err := rows.Scan(&counter.Name, &counter.Code, &counter.Context, &counter.BType, &counter.BFamily,
			&counter.BInternal, &counter.BAutomated, &counter.BCountLic, &counter.BCountEnt, &counter.BCountInst, &counter.BLicUpgrade)
		if err != nil {
			fmt.Printf("rows.Scan failed. %v\n", err)
			return counter, err
		}
	}
	//fmt.Println(counter)
	return counter, nil
}
