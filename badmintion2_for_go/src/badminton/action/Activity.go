package action

import (
	"tweb"
	"log"
	"time"
	"badminton/dao"
	"strconv"


)

type ActivityAct struct{}

func init() {
	activityAct := &ActivityAct{}
	tweb.RegisteHandler("/groupAty", activityAct.GroupAty)
}
func (t *ActivityAct) GroupAty(ctx *tweb.Ctx) {
	openid:=ctx.GetSession("openid")
	openid="o5ON-wBOzYaLu94o1u_12ciWj2EY"
	ctx.PutSession("openid",openid)
	groupId:=ctx.R.FormValue("groupId")
	startDate:=ctx.R.FormValue("startDate")
	endDate:=ctx.R.FormValue("endDate")
	if groupId==""{
		groupId="1"
	}
	if startDate=="" || endDate==""{
		//获取当天活动
		now:=time.Now()
		//dd, _ := time.ParseDuration("24h")
		//dd1 := now.Add(dd)
		startDate=now.Format("2006-01-02")
		//endDate=dd1.Format("2006-01-02")
		endDate=startDate

	}

	_groupId,_:=strconv.Atoi(groupId)
	atyList:=dao.FindAtyByDate(startDate,endDate,_groupId)
	data:=map[string]interface{}{}
	if len(atyList)==1{
		(*atyList[0]).StartTime=atyList[0].StartTime[11:16]
		(*atyList[0]).EndTime=atyList[0].EndTime[11:16]
		data["aty"]=atyList[0]
	}

	b,member:=MemberLogin(openid,ctx)
	if !b{
		data["isLogin"]=false
	}else{
		data["isLogin"]=true
		data["member"]=member
	}
	userInfo,_err:=tweb.WxMux.GetUserInfo(openid)
	if _err!=nil{
		log.Println("获取用户信息失败..")
	}else{
		data["headImgUrl"]=(*userInfo).HeadImageUrl
		ctx.PutSession_obj("userInfo",userInfo)
	}
	data["openid"]=openid
	data["groupid"]=_groupId
	if groupId!=""{
		group:=dao.FindGroupInfoById(_groupId)
		if group!=nil{
			data["group"]=group
			ctx.PutSession_obj("group",group)
		}


		//fmt.Println(data["group"].(entity.Group).Name)
	}
	if len(atyList)==1{
		signUpList:=dao.QuerySignUpMemberByGroupId(_groupId,atyList[0].Id)
		data["signUpList"]=signUpList
	}else{
		data["signUpList"]=""
	}


	if b && len(atyList)==1{
		//如果已登录查询是否报名
		tmp:=dao.QueryMemberAtyRecordByAtyId(atyList[0].Id,openid)
		//fmt.Println("tmp:",tmp)
		if tmp["recordId"]!=nil {
			data["recordId"]=tmp["recordId"].(int)
			//fmt.Println("recordId:",tmp["recordId"])
		}

	}
	ctx.Render("groupAty.html",data)
}
