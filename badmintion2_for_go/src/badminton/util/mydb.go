package util

import (
	"tweb/db"
	"database/sql"




	"fmt"
	"time"
	"reflect"
)

const (
	SEL         = "select %s from %s where %s %s %s"
	INSERT      = "insert into %s (%s) values(%s)"
	INSERTBULK  = "INSERT INTO %s (%s) SELECT %s UNION ALL SELECT %s"
	UPDATE      = "update %s set %s where %s"
	DEL         = "delete from %s where %s"
	CREATE      = "create table if not exists %s (%s)"
	CREATEINDEX = "CREATE INDEX %s ON %s (%s)"
	DROP        = "drop table if exists %s"
	INT ="int"
	STRING="string"
	FLOAT32="float32"
	FLOAT64="float64"
	TIME="time"

)

func Find(_sql string,colType []string ,param ...interface{}) []map[string]interface{} {
	data := make([]map[string]interface{}, 0)
	db.Query(func(sr *sql.Rows) error {
		fs:=getZero(colType)
		scanArgs := make([]interface{}, len(colType))
		for i:=range fs{
			scanArgs[i]=&fs[i]
		}
		for sr.Next(){
			sr.Scan(scanArgs...)
			for _, val:= range fs {
				fmt.Print(reflect.TypeOf(val),",")
			}
		}
		return nil
	},_sql,param...)
	return data
}

func getZero(colType []string) []interface{} {
	res:=make([]interface{},0)
	for _,col:=range colType{
		switch col {
		case INT:
			var i int
			res=append(res,i)
		case TIME:
			var t time.Time
			res=append(res,t)
		default:
			var s string
			res=append(res,s)
		}
	}
	//fmt.Println("res:",res)
	return res

}



