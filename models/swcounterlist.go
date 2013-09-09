package models

import (
	_ "code.google.com/p/odbc"
	//"database/sql"
	"fmt"
	conf "github.com/spartacusX/minisam/config"
)

type SWCounter struct {
	LCounterId        int
	LEntCount         int
	LLicUseRights     int
	LSoftInstallCount int
	LUnusedInstall    int
	Name              string
	Status            string
}

type SWCounterList []SWCounter

func (p SWCounterList) TotalPages() int {
	num := len(p) / conf.RecordNumPerPage
	if num != 0 {
		num++
	}
	return num
}

func (p SWCounterList) GetPageByIndex(pageNum int) (result SWCounterList) {
	// range check
	if pageNum < 1 || pageNum > p.TotalPages() {
		return result
	}

	pageNum-- //base 0

	if len(p) >= (pageNum+1)*conf.RecordNumPerPage {
		result = p[pageNum*conf.RecordNumPerPage : (pageNum+1)*conf.RecordNumPerPage]
	} else if len(p) <= (pageNum+2)*conf.RecordNumPerPage { // last page
		result = p[pageNum*conf.RecordNumPerPage:]
	} else {
		result = p[pageNum*conf.RecordNumPerPage:]
	}
	return result
}

func ListCounter() (counters []SWCounter, err error) {
	db, err := sqlConnect()
	if err != nil {
		return counters, err
	}

	defer db.Close()

	strQuery := "SELECT lCountId, name, dLicUseRights, dEntCount, dSoftInstallCount, dUnusedInstall " +
		" FROM " + conf.Scheme + "." + TBL_SOFTWARECOUNTER +
		" WHERE bType = 0 " +
		" ORDER BY name "

	rows, err := db.Query(strQuery)
	if err != nil {
		fmt.Printf("db.Query failed. %v\n", err)
		return counters, err
	}

	// Escape the first record( NULL record)
	rows.Next()

	for rows.Next() {
		var r SWCounter
		err := rows.Scan(&r.LCounterId, &r.Name, &r.LLicUseRights, &r.LEntCount, &r.LSoftInstallCount, &r.LUnusedInstall)
		if err != nil {
			fmt.Printf("rows.Scan failed. %v\n", err)
			return counters, err
		} else {
			statistic.Entitlements += r.LEntCount
			statistic.Licenses += r.LLicUseRights
			statistic.Installations += r.LSoftInstallCount
			statistic.UnUsedInstallations += r.LUnusedInstall

			switch compliance := r.LLicUseRights - r.LSoftInstallCount; {
			case compliance > 5:
				r.Status = ColorLevel[NORMAL]
			case compliance >= 0:
				r.Status = ColorLevel[WARNING]
			case compliance < 0:
				r.Status = ColorLevel[ERROR]
			default:
				r.Status = ColorLevel[NORMAL]
			}

			counters = append(counters, r)
		}
	}
	//fmt.Println(counters)
	return counters, nil
}
