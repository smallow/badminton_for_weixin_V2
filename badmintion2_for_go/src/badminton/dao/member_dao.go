package dao

import (
	"tweb/db"
	"log"
	"database/sql"

)

func SaveMember(name, phone, openid, group_num, group_name, invitation_code string) bool {
	_sql := "insert into badminton_member (name,phone,openid,group_num,group_name,invitation_code)" +
		"values (?,?,?,?,?,?)"
	_, _, err := db.Exec(_sql, name, phone, openid, group_num, group_name, invitation_code)
	if err != nil {
		log.Println("会员注册失败", err.Error())
		return false
	}
	return true
}

func QueryMemberAtyRecordByAtyId(atyId int, openid string) map[string]interface{} {
	_sql := "select id from view_aty_record where aty_id=? and openid=? and is_cancel=0 "
	data := map[string]interface{}{}
	err := db.Query(func(sr *sql.Rows) error {
		var recordId int
		for sr.Next() {
			sr.Scan(&recordId)
			data["recordId"] = recordId
		}
		return nil
	}, _sql, atyId, openid)
	if err != nil {
		log.Println("查询会员:", openid, "活动记录出错", err.Error())
		return nil
	}
	return data

}

func QueryMemberAtyRecordList(memberId,groupId int) []map[string]interface{}{
	//fmt.Println("memberId:",memberId,"groupId:",groupId," 查询...")
	data:=make([]map[string]interface{},0)
	_sql:="select id,date,group_name,address from view_aty_record where member_id=? and group_id=? order by date desc"
	db.Query(func(sr *sql.Rows) error {
		var id int
		var date,groupName,address string
		for sr.Next(){
			sr.Scan(&id,&date,&groupName,&address)
			_map:=map[string]interface{}{
				"recordId":id,
				"date":date,
				"groupName":groupName,
				"address":address,
			}
			//fmt.Println("_map:",_map)
			data=append(data,_map)
		}
		return nil
	},_sql,memberId,groupId)
	return data
}
func QueryMyAtyDetail(recordId int) map[string]interface {}{
	data := map[string]interface{}{}
	_sql:="select name,money,current_day_left_money,friend_num,friend_names,sign_up_time,is_cancel,cancel_time,address,start_time,end_time,date,total_cost,badminton_num,site_num,time_num,avg_cost,total_person,group_name from view_aty_record where id=? "
	err:=db.Query(func(sr *sql.Rows) error {
		var name string
		var money,currentDayLeftMoney float64
		var friendNum int
		var friend_names,signUpTime string
		var isCancel bool
		var cancelTime,address,startTime,endTime,date sql.NullString
		var totalCost,avgCost float64
		var badmintonNum,siteNum,timeNum,totalPerson int
		var groupName string
		 for sr.Next(){
			sr.Scan(&name,&money,&currentDayLeftMoney,&friendNum,&friend_names,&signUpTime,&isCancel,&cancelTime,&address,&startTime,&endTime,&date,&totalCost,&badmintonNum,&siteNum,&timeNum,&avgCost,&totalPerson,&groupName)
			data["name"]=name
			data["money"]=money
			data["currentDayLeftMoney"]=currentDayLeftMoney
			data["friendNum"]=friendNum
			 data["friendNames"]=friend_names
			 data["signUpTime"]=signUpTime
			 data["isCancel"]=isCancel
			 data["cancelTime"]=cancelTime.String
			 data["address"]=address.String
			 data["startTime"]=startTime.String
			 data["endTime"]=endTime.String
			 data["date"]=date.String
			 data["totalCost"]=totalCost
			 data["siteNum"]=siteNum
			 data["badmintonNum"]=badmintonNum
			 data["timeNum"]=timeNum
			 data["avgCost"]=avgCost
			 data["totalPerson"]=totalPerson
			 data["groupName"]=groupName
		}
		return nil
	},_sql,recordId)
	if err!=nil{
		log.Println("QueryMyAtyDetail出错:",err.Error())
	}
	return data

}


