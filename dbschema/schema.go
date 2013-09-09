package dbschema

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	DBOBJECTTYPE_TABLE = iota
	DBOBJECTTYPE_FIELD
	DBOBJECTTYPE_LINK
	DBOBJECTTYPE_INDEX
)

type DbItem struct {
	SqlName  string `xml:"sqlname,attr"`
	Name     string `xml:"name,attr"`
	Desc     string `xml:"desc,attr"`
	LongDesc string `xml:"longdesc,attr"`
	Type     string `xml:"type,attr"`
}

type DbField struct {
	DbItem
	// Size     int    `xml:"size,attr"`
	// Advanced string `xml:"advanced,attr"`
}

type DbLink struct {
	DbItem
	// Size     int    `xml:"size,attr"`
	// Advanced string `xml:"advanced,attr"`
}

type DbIndex struct {
	DbItem
	// Size     int    `xml:"size,attr"`
	// Advanced string `xml:"advanced,attr"`
}

type DbTable struct {
	DbItem
	// MainFieldId        string    `xml:"mainfield,attr"`
	// MainIndex          string    `xml:"mainindex,attr"`
	// InternationalIndex string    `xml:"internationalindex,attr"`
	// Archivable         string    `xml:"archivable,attr"`
	Fields  []DbField `xml:"field"`
	Links   []DbLink  `xml:"link"`
	Indexes []DbIndex `xml:"index"`
}

type DBBSchema struct {
	Indentifier string `xml:"identifier,attr"`
	Version     string `xml:"version,attr"`
	Format      string `xml:"format,attr"`
	// Archivable  string    `xml:"archivable,attr"`
	Tables []DbTable `xml:"table"`
}

var schema DBBSchema

func init() {
	f, err := os.Open("./gbbase_en.xml")
	if err != nil {
		fmt.Printf("Open gbbase.xml failed for reason: %v\n", err)
		return
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Read gbbase.xml failed for reason: %v\n", err)
		return
	}

	err = xml.Unmarshal(content, &schema)
	if err != nil {
		fmt.Println("Unmarshal gbbase.xml failed for reason: %v\n", err)
		return
	}
	//fmt.Println(schema)
}

func GetTableSchemaBySqlName(sqlname string) (t DbTable, err error) {
	for _, v := range schema.Tables {
		if v.SqlName == sqlname {
			return v, nil
		}
	}
	return t, fmt.Errorf("Can not find table: %v ", sqlname)
}

func (t DbTable) GetFieldByName(name string) (f DbField, err error) {
	for _, v := range t.Fields {
		if v.Name == name {
			return v, nil
		}
	}
	return f, fmt.Errorf("Can not find field: %v", name)
}

func (t DbTable) GetLinkByName(name string) (l DbLink, err error) {
	for _, v := range t.Links {
		if v.Name == name {
			return v, nil
		}
	}
	return l, fmt.Errorf("Can not find link: %v", name)
}
