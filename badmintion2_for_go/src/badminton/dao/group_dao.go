package dao

import (
	"badminton/entity"
	"tweb/db"
	"database/sql"

	"log"
)

func FindGroupInfoById(id int) *entity.Group {
	_sql:="select id,name,group_num,group_type,group_master,group_master_phone from badminton_group where id=?"
	var group=&entity.Group{}
	err:=db.Query(func(sr *sql.Rows) error {
		var id int
		var name,group_num,group_type,group_master,group_master_phone string
		for sr.Next(){
			err:=sr.Scan(&id,&name,&group_num,&group_type,&group_master,&group_master_phone)
			if err!=nil{
				log.Println(err.Error())
			}
			parseGroup(group,id,name,group_num,group_type,group_master,group_master_phone)
		}
		return nil
	},_sql,id)

	if err!=nil{
		log.Println(err.Error())
	}

	return group
}

func parseGroup(group *entity.Group,id int ,name,group_num,group_type,group_master,group_master_phone string) {
	//fmt.Println("查询得到:",id,name,group_num)
	group.Id=id
	group.Name=name
	group.GroupNum=group_num
	group.GroupType=group_type
	group.GroupMaster=group_master
	group.GroupMasterPhone=group_master_phone
}
