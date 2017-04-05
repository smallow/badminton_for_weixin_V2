package dao

import (
	"tweb/db"
	"log"
	"time"
	"database/sql"
	"fmt"
)

func SaveAtyRecord(member_id, aty_id int, money, current_day_left_money float64, friend_num int, friend_names string, group_id int,group_name string)(int,error)  {
	_sql:="select id from badminton_activity_record where aty_id=? and member_id=? "
	var recordId int
	sign:=false
	db.Query(func(sr *sql.Rows) error {
		for sr.Next(){
			sr.Scan(&recordId)
		}
		return nil
	},_sql,aty_id,member_id)
	if recordId!=0{

		_sql="update badminton_activity_record set sign_up_time=?,update_time=?,is_cancel=?,cancel_time=? where id=?"
		now:=time.Now().Format("2006-01-02 15:04:05")
		_,_,err:=db.Exec(_sql,now,now,0,nil,recordId)
		if err!=nil{
			log.Println("UpdateAtyRecord出错: ", err.Error())
			return  0,err
		}
		sign=true

	}else{
		_sql= "insert into badminton_activity_record (member_id,aty_id,money,current_day_left_money,friend_num,friend_names,sign_up_time,is_cancel,memo,group_id,group_name) values " +
			"(?,?,?,?,?,?,?,?,?,?,?)"
		now:=time.Now().Format("2006-01-02 15:04:05")
		__recordId, _, err := db.Exec(_sql, member_id, aty_id, money, current_day_left_money, friend_num, friend_names, now, 0, "", group_id,group_name)
		if err != nil {
			log.Println("SaveAtyRecord出错: ", err.Error())
			return  0,err
		}
		recordId=int(__recordId)
		sign=true
	}

	if sign{
		//报名成功 更新报名人数
		updateSignUpNum(aty_id,group_id)
	}

	return recordId,nil

}

func updateSignUpNum(atyId,groupId int) {
	_sql:="update badminton_activity set total_person=(select count(id) as num from badminton_activity_record where aty_id=? and group_id=? and is_cancel=0)"
	_,_,err:=db.Exec(_sql,atyId,groupId)
	if err!=nil{
		log.Println("更新活动人数出错:",err.Error())
	}
	//fmt.Println("更新成功!")
}

func CancelSignUp(atyId,recordId int) (bool,string) {
	_sql:="select status,group_id from badminton_activity where id=? "
	var status string
	var groupId int
	err:=db.Query(func(sr *sql.Rows) error {
		for sr.Next(){
			err:=sr.Scan(&status,&groupId)
			if err!=nil{
				fmt.Println("err:",err)
				return err
			}
		}
		return nil
	},_sql,atyId)
	if err!=nil{
		fmt.Println("err:",err.Error())
	}

	if status=="" || status=="01"{
		now:=time.Now().Format("2006-01-02 15:04:05")
		_sql = "update badminton_activity_record set is_cancel=?,cancel_time=?,update_time=? where id=?"
		_, _, err := db.Exec(_sql, 1, now,now,recordId)
		if err != nil {
			log.Println("CancelSignUp出错: ", err.Error())
			return false,"error"
		}
		//报名成功 更新报名人数
		if groupId!=0{
			updateSignUpNum(atyId,groupId)
		}
		return true,"success"
	}else if status=="02"{
		return false,"aty_already_started"
	}else if status=="03"{
		return false,"aty_already_finished"
	}

	return false,"fail"

}
