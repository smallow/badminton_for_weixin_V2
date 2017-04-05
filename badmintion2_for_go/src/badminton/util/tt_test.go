package util

import (
	"testing"


	_"badminton/entity"
	_"tweb/db"
	_"database/sql"
	"fmt"
)


type AtyTest struct {
	Id int `json:"id"`
	Name string `json:"name"`
	AgeNum int `json:"age_num"`

}
func TestData(t *testing.T) {
	//cols:=[]string{"id","name","ageNum"}
	//dest := make([]interface{}, len(cols))
	//object:=&AtyTest{}
	//el := reflect.ValueOf(object).Elem()
	//for i, c := range cols {
	//	dest[i] = el.FieldByName(strings.Title(c)).Addr().Interface()
	//}
	//
	//fmt.Println("dest:",dest)

	//colum := []string{"id","address","start_time" ,"end_time","date","charge_member_id","charge_member_name","charge_member_phone","total_cost","badminton_nums","site_nums","time_nums","avg_cost","total_person","group_id","status"}
	//dest := make([]interface{}, len(colum))
	//aty:=&entity.Activity{}
	//el := reflect.ValueOf(aty).Elem()
	//for i, c := range colum {
	//	dest[i] = el.FieldByName(strings.Title(c)).Addr().Interface()
	//}
	//fmt.Println("aty:",aty)
	//object:=&AtyTest{}
	//el := reflect.ValueOf(object).Elem()
	//typeOfType := el.Type()
	//for i:=0; i<el.NumField(); i++{
	//	field := el.Field(i)
	//	fmt.Printf("%d. %s %s %s = %v \n", i, typeOfType.Field(i).Name,typeOfType.Field(i).Tag.Get("json"), field.Type(),field.Interface())
	//}

	//o := DBObject{
	//	TableName: "badminton_activity",
	//	Fields:    []string{"id", "address","start_time"},
	//	MapKV:     QuickNewSortMap(KV{"id", 1}),
	//	Order:     "id desc",
	//}
	//dataList:=FindInfo(o,true)
	//fmt.Println(dataList)

	//db.Query(func(rows *sql.Rows) error {
	//	columns, _ := rows.Columns()
	//	values := make([]sql.RawBytes, len(columns))
	//	scanArgs := make([]interface{}, len(values))
	//	for i := range values {
	//		scanArgs[i] = &values[i]
	//		fmt.Println("values[i]:",values[i]," scanArgs[i]:",scanArgs[i])
	//	}
	//
	//	for rows.Next() {
	//		// get RawBytes from data
	//		rows.Scan(scanArgs...)
	//		// Now do something with the data.
	//		// Here we just print each column as a string.
	//		var value string
	//		for i, col := range values {
	//			// Here we can check if the value is nil (NULL value)
	//			if col == nil {
	//				value = "NULL"
	//			} else {
	//				value = string(col)
	//			}
	//			fmt.Println(columns[i], ": ", value)
	//		}
	//		fmt.Println("-----------------------------------")
	//	}
	//	return nil
	//
	//},"select id,address,start_time,end_time from badminton_activity")
	str:=192.568
	fmt.Println(fmt.Sprintf("%.2f",str))

}
