package dao

import (
	"badminton/entity"
	"tweb/db"
	"database/sql"
	"log"
)

/**
根据开始结束日期,群ID查询活动
 */
func FindAtyByDate(startDate, endDate string, groupId int) []*entity.Activity {
	//colType:=[]string{util.INT,util.STRING,util.STRING,util.STRING,util.STRING,util.STRING,util.STRING,util.INT,util.INT,util.STRING}
	//util.Find(_sql,colType,startDate,endDate,groupId)
	//fmt.Println("查询日期段群活动,groupId:", groupId, "startDate:", startDate, "endDate:", endDate)
	atyList := make([]*entity.Activity, 0)
	//data:=make([]map[string]interface{},0)
	_sql := "select id,address,start_time,end_time,date,charge_member_name,charge_member_phone,site_num,time_num,total_person,status from badminton_activity where date between ? and ? and group_id=?"
	db.Query(func(sr *sql.Rows) error {
		var id int
		var site_num,time_num,total_person int
		var address,start_time,end_time,date,charge_person_name,charge_person_phone,status string
		for sr.Next(){
			sr.Scan(&id,&address,&start_time,&end_time,&date,&charge_person_name,&charge_person_phone,&site_num,&time_num,&total_person,&status)
			_data:=make(map[string]interface{},0)
			_data["id"]=id
			_data["address"]=address
			_data["start_time"]=start_time
			_data["end_time"]=end_time
			_data["date"]=date
			_data["charge_member_name"]=charge_person_name
			_data["charge_member_phone"]=charge_person_phone
			_data["status"]=status
			_data["site_num"]=site_num
			_data["time_num"]=time_num
			_data["total_person"]=total_person
			//fmt.Println("_data",_data)
			aty:=&entity.Activity{}
			parseAty(aty,_data)
			//fmt.Println("aty:",aty)
			atyList=append(atyList,aty)
		}
		return nil
	},_sql,startDate,endDate,groupId)
	return atyList

}

func QuerySignUpMemberByGroupId(groupId,atyId int)[]map[string]interface{} {
	data:=make([]map[string]interface{},0)
	_sql:="select member_id,name,openid from view_aty_record where group_id=? and aty_id=?  and is_cancel=0 order by id desc"
	err:=db.Query(func(sr *sql.Rows) error {
		var memberId int
		var name ,openid string
		for sr.Next(){
			sr.Scan(&memberId,&name,&openid)
			_map:=map[string]interface{}{
				"memberId":memberId,
				"name":name,
				"openid":openid,
			}
			data=append(data,_map)
		}
		return nil
	},_sql,groupId,atyId)
	if err!=nil{
		log.Println("查询群活动报名成员出错:",err.Error())
	}
	return data
}


func parseAty(aty *entity.Activity,_map map[string]interface{})  {
	if len(_map)>0{
		for k,v:=range _map{
			switch k {
			case "id":
				(*aty).Id=v.(int)
			case "address":
				(*aty).Address=v.(string)
			case "start_time":
				(*aty).StartTime=v.(string)
			case "end_time":
				(*aty).EndTime=v.(string)
			case "date":
				(*aty).Date=v.(string)
			case "charge_member_id":
				(*aty).ChargeMemberId=v.(int)
			case "charge_member_name":
				(*aty).ChargeMemberName=v.(string)
			case "charge_member_phone":
				(*aty).ChargeMemberPhone=v.(string)
			case "total_cost":
				(*aty).TotalCost=v.(float64)
			case "avg_cost":
				(*aty).AvgCost=v.(float64)
			case "badminton_num":
				(*aty).BadmintonNum=v.(int)
			case "site_num":
				(*aty).SiteNum=v.(int)
			case "time_num":
				(*aty).TimeNum=v.(int)
			case "total_person":
				(*aty).TotalPerson=v.(int)
			case "group_id":
				(*aty).GroupId=v.(int)
			case "status":
				(*aty).Status=v.(string)
			default:

			}
		}

	}
}
