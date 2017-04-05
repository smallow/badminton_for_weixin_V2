package dao

import (
	"tweb/db"
	"database/sql"
	"log"
)

func QueryMyPaymentRecordList(memberId, groupId int) []map[string]interface {}{
	data:=make([]map[string]interface{},0)
	_sql:="select id,name,group_name,money,bill_number,pay_type,payee_name,payee_account,pay_time,memo from view_payment_record where member_id=? and group_id=? "
	err:=db.Query(func(sr *sql.Rows) error {
		var id int
		var name,groupName string
		var money float64
		var billNumber,payType ,payeeName,payeeAccount,payTime,memo string
		for sr.Next(){
			err:=sr.Scan(&id,&name,&groupName,&money,&billNumber,&payType,&payeeName,&payeeAccount,&payTime,&memo)
			if err!=nil{
				return err
			}
			_map:=map[string]interface{}{
				"id":id,
				"name":name,
				"groupName":groupName,
				"money":money,
				"billNumber":billNumber,
				"payType":payType,
				"payeeName":payeeName,
				"payeeAccount":payeeAccount,
				"payTime":payTime,
				"memo":memo,
			}
			data=append(data,_map)
		}

		return nil
	},_sql,memberId,groupId)

	if err!=nil{
		log.Panic(err)
	}

	return data
	
}

func QueryMyPaymentDetail(recordId int) map[string]interface {}{
	data := map[string]interface{}{}

	_sql:="select name,group_name,money,pay_money,bill_number,pay_type,payee_name,payee_account,pay_time,memo from view_payment_record where id=? "
	err:=db.Query(func(sr *sql.Rows) error {

		var name,groupName string
		var money ,payMoney float64
		var billNumber,payType ,payeeName,payeeAccount,payTime,memo string
		for sr.Next(){
			err:=sr.Scan(&name,&groupName,&money,&payMoney,&billNumber,&payType,&payeeName,&payeeAccount,&payTime,&memo)
			if err!=nil{
				return err
			}

				data["name"]=name
				data["groupName"]=groupName
				data["money"]=money
				data["payMoney"]=payMoney
				data["billNumber"]=billNumber
				data["payType"]=payType
				data["payeeName"]=payeeName
				data["payeeAccount"]=payeeAccount
				data["payTime"]=payTime
				data["memo"]=memo


		}

		return nil
	},_sql,recordId)

	if err!=nil{
		log.Panic(err)
	}
	return data
}