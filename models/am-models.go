package models

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"fmt"
	conf "github.com/spartacusX/minisam/config"
)

type SWStatistic struct {
	Entitlements        int
	Licenses            int
	Installations       int
	UnUsedInstallations int
}

func TotalCount(sws *SWStatistic) (count int) {
	if sws != nil {
		return sws.Entitlements +
			sws.Licenses +
			sws.Installations +
			sws.UnUsedInstallations
	}
	return 0
}

var statistic SWStatistic

var ColorLevel = [3]string{"info", "warning", "error"}

func Statistic() (sws SWStatistic) {
	return statistic
}

var TABLE_EMPLDEPT = "AMDemo94en.itam.amEmplDept"

const TBL_SOFTWARECOUNTER = "amSoftLicCounter"

const TBL_COUNTERRESULT = "amRightsUsesCount"

/**************************************************************
/ Open database, not really connect, just regist driver and dsn
/**************************************************************/
func sqlConnect() (db *sql.DB, err error) {
	db, err = sql.Open(conf.Driver, conf.CnxString)
	if err != nil {
		fmt.Printf("sql.Open failed. %v\n", err)
		return nil, err
	}
	return db, nil
}

func CounterList() (ctlist []SoftWareCounter, err error) {
	p := []SoftWareCounter{}
	db, err := sqlConnect()
	if err != nil {
		return p, err
	}

	defer db.Close()

	/*strQuery := "SELECT name, code, type, context, bType, bFamily, bAutomated, bCountLic, bCountEnt, bCountInst, bLicUpgrade, bInternal, " +
	" dLicUseRights, dEntCount, dSoftInstallCount, dUnusedInstall " +*/
	strQuery := "SELECT lCountId, name, Code, Context, bType, bFamily, bInternal, bAutomated, " +
		"bCountLic, bCountEnt, bCountInst, bLicUpgrade, dLicUseRights, dEntCount, dSoftInstallCount, dUnusedInstall " +
		" FROM " + conf.Scheme + "." + TBL_SOFTWARECOUNTER +
		" WHERE bType = 0 " +
		" ORDER BY name "

	rows, err := db.Query(strQuery)
	if err != nil {
		fmt.Printf("db.Query failed. %v\n", err)
		return p, err
	}

	// Escape the first record( NULL record)
	rows.Next()

	for rows.Next() {
		var r SoftWareCounter
		err := rows.Scan(&r.LCounterId, &r.Name, &r.Code, &r.Context, &r.BType, &r.BFamily, &r.IsInternal, &r.BAutomated, &r.BCountLic,
			&r.BCountEnt, &r.BCountInst, &r.BLicUpgrade, &r.LLicUseRights, &r.LEntCount, &r.LSoftInstallCount, &r.LUnusedInstall)
		if err != nil {
			fmt.Printf("rows.Scan failed. %v\n", err)
			return p, err
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

			p = append(p, r)
		}
	}
	//fmt.Println(p)
	return p, nil
}

// func AddCounter(c SWCounter) (err error) {
// 	db, err := sqlConnect()
// 	if err != nil {
// 		return err
// 	}

// 	defer db.Close()

// 	//res, err := db.Exec("insert into " + TABLE + " values ( " + c.Id + ", " + c.Model + ", " + c.LicenseNum + ", " c.InstallationNum + ")")
// 	res, err := db.Exec("insert into "+TBL_SOFTWARECOUNTER+" values ( ?, ?, ?, ?)", c.Id, c.Model, c.LicenseNum, c.InstallationNum)
// 	if err != nil {
// 		fmt.Printf("db.Exec failed. %v\n.", err)
// 		return err
// 	}

// 	rowAffected, _ := res.RowsAffected()
// 	fmt.Printf("Affected rows: %v\n", rowAffected)

// 	return nil
// }

// func UnlockAmUser(strUser string) (err error) {
// 	db, err := sqlConnect()
// 	if err != nil {
// 		return err
// 	}

// 	defer db.Close()

// 	fmt.Println(strUser)
// 	_, err = db.Exec("update "+TABLE_EMPLDEPT+" set  LoginFailures=0, seLoginStatus=0 where Name=?", strUser)
// 	if err != nil {
// 		fmt.Printf("db.Exec failed. %v\n.", err)
// 		return err
// 	}

// 	return nil
// }
