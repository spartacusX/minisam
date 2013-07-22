package main

import (
	_ "code.google.com/p/odbc"
	"database/sql"
	"fmt"
)

type People struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("odbc", "dsn=AMDemo94en;user id=sa;password=sasa")
	if err != nil {
		fmt.Println("sql.Open failed. %v", err)
	}

	defer db.Close()

	rows, err := db.Query("select id, name from AMDemo94en.dbo.amTest")
	if err != nil {
		fmt.Println("db.Query failed. %v", err)
	}

	p := []People{}

	for rows.Next() {
		var r People
		err := rows.Scan(&r.Id, &r.Name)
		if err != nil {
			fmt.Println("rows.Scan failed. %v", err)
		} else {
			p = append(p, r)
		}
	}
	fmt.Println(p)
}
