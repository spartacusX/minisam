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

var TABLE = "AMDemo94en.dbo.amSoftwareCounter"
var CNXSTRING = "dsn=AMDemo94en;uid=sa;pwd=sasa"

func CounterList() (ctlist []SWCounter, err error) {
	p := []SWCounter{}
	db, err := sql.Open("odbc", CNXSTRING)
	if err != nil {
		fmt.Println("sql.Open failed. %v", err)
		return p, err
	}

	defer db.Close()

	rows, err := db.Query("select * from " + TABLE)
	if err != nil {
		fmt.Println("db.Query failed. %v", err)
		return p, err
	}

	for rows.Next() {
		var r SWCounter
		err := rows.Scan(&r.Id, &r.Model, &r.LicenseNum, &r.InstallationNum)
		if err != nil {
			fmt.Println("rows.Scan failed. %v", err)
			return p, err
		} else {
			p = append(p, r)
		}
	}
	fmt.Println(p)
	return p, nil
}

func AddCounter(c SWCounter) (err error) {
	db, err := sql.Open("odbc", CNXSTRING)
	if err != nil {
		fmt.Println("sql.Open failed. %v", err)
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
