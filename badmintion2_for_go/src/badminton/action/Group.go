package action

import (
	"fmt"
	"tweb"
	"tweb/db"

	"database/sql"
)

type GroupAct struct{}

func init() {
	groupAct := &GroupAct{}
	tweb.RegisteHandler("/getAllGroup", groupAct.GetAllGroup)
	tweb.RegisteHandler("/sys/reloadtpl",groupAct.ReloadTpl)
}

//重新加载所有模板
func (t *GroupAct) ReloadTpl(ctx *tweb.Ctx) {
	fmt.Println("刷新模板....")
	ctx.W.Header().Add("Pragma", "No-cache")
	ctx.W.Header().Add("Cache-Control", "no-cache")
	ctx.W.Header().Add("Expires", "0")
	tweb.LoadAllTemplates("./static/templates/")
	//http.Redirect(ctx.W, ctx.R, "/", 301)
}


func (i *GroupAct) GetAllGroup(ctx *tweb.Ctx) {
	//rows,_ :=util.DB.Query("select * from badminton_group order by id desc")
	//defer rows.Close()
	//var id int
	//var name string
	//var group_num string
	//var master_name,master_phone string
	//for rows.Next(){
	//	rows.Scan(&id,&name,&master_name,&master_phone,&group_num)
	//	fmt.Println(id,name,master_name,master_phone,group_num)
	//}
	db.Query(func(sr *sql.Rows) error {
		var id int
		var name string
		var group_num string
		var master_name,master_phone string
		for sr.Next(){
			sr.Scan(&id,&name,&master_name,&master_phone,&group_num)
			fmt.Println(id,name,master_name,master_phone,group_num)
		}
		return nil
	},"select * from badminton_group")

}
