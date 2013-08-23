package models

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"fmt"
)

type SWCounter struct {
	Id              int
	Model           string
	LicenseNum      int
	InstallationNum int
}

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

func Statistic() (sws SWStatistic) {
	return statistic
}

var TABLE = "AMDemo94en.dbo.amSoftwareCounter"

var TABLE_EMPLDEPT = "AMDemo94en.itam.amEmplDept"

const DRIVER = "odbc"

const DB_SCHEME = "AMDemo94en.itam"

const CNXSTRING = "dsn=AMDemo94en;uid=sa;pwd=sasa"

const TBL_SOFTWARECOUNTER = "amSoftLicCounter"

const TBL_COUNTERRESULT = "amRightsUsesCount"

/**************************************************************
/ Open database, not really connect, just regist driver and dsn
/**************************************************************/
func sqlConnect() (db *sql.DB, err error) {
	db, err = sql.Open(DRIVER, CNXSTRING)
	if err != nil {
		fmt.Println("sql.Open failed. %v", err)
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

	strQuery := "SELECT name, bInternal, dLicUseRights, dEntCount, dSoftInstallCount, dUnusedInstall " +
		" FROM " + DB_SCHEME + "." + TBL_SOFTWARECOUNTER +
		" WHERE bType = 0 " +
		" ORDER BY name "

	rows, err := db.Query(strQuery)
	if err != nil {
		fmt.Println("db.Query failed. %v", err)
		return p, err
	}

	// Escape the first record( NULL record)
	rows.Next()

	for rows.Next() {
		var r SoftWareCounter
		err := rows.Scan(&r.Name, &r.IsInternal, &r.LicUseRights, &r.EntCount, &r.SoftInstallCount, &r.UnusedInstall)
		if err != nil {
			fmt.Println("rows.Scan failed. %v", err)
			return p, err
		} else {
			statistic.Entitlements += r.EntCount
			statistic.Licenses += r.LicUseRights
			statistic.Installations += r.SoftInstallCount
			statistic.UnUsedInstallations += r.UnusedInstall
			p = append(p, r)
		}
	}
	//fmt.Println(p)
	return p, nil
}

func AddCounter(c SWCounter) (err error) {
	db, err := sqlConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	//res, err := db.Exec("insert into " + TABLE + " values ( " + c.Id + ", " + c.Model + ", " + c.LicenseNum + ", " c.InstallationNum + ")")
	res, err := db.Exec("insert into "+TABLE+" values ( ?, ?, ?, ?)", c.Id, c.Model, c.LicenseNum, c.InstallationNum)
	if err != nil {
		fmt.Println("db.Exec failed. %v.", err)
		return err
	}

	rowAffected, _ := res.RowsAffected()
	fmt.Println("Affected rows: %v", rowAffected)

	return nil
}

func UnlockAmUser(strUser string) (err error) {
	db, err := sqlConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	fmt.Println(strUser)
	_, err = db.Exec("update "+TABLE_EMPLDEPT+" set  LoginFailures=0, seLoginStatus=0 where Name=?", strUser)
	if err != nil {
		fmt.Println("db.Exec failed. %v.", err)
		return err
	}

	return nil
}
