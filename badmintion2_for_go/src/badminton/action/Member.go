package action

import (
	"tweb"

	"tweb/db"
	"database/sql"
	"badminton/entity"
	"log"
	"github.com/dchest/captcha"
	"badminton/dao"
	"strconv"


)

type MemberAct struct {
}

func init() {
	memberAct := &MemberAct{}
	tweb.RegisteHandler("/signUp", memberAct.SignUp)
	tweb.RegisteHandler("/cancelSignUp", memberAct.CancelSignUp)
	tweb.RegisteHandler("/memberRegister",memberAct.MemberRegister)
	tweb.RegisteHandler("/captcha",memberAct.Captcha)
	tweb.RegisteHandler("/memberRegisterSubmit",memberAct.MemberRegisterSubmit)
	tweb.RegisteHandler("/toUserInfo",memberAct.ToUserInfo)
	tweb.RegisteHandler("/toMyAty",memberAct.ToMyAty)
	tweb.RegisteHandler("/toMyPayment",memberAct.ToMyPayment)
	tweb.RegisteHandler("/getMyAty",memberAct.GetMyAtyList)
	tweb.RegisteHandler("/toMyAtyDetail",memberAct.ToMyAtyDetail)
	tweb.RegisteHandler("/getMyAtyDetail",memberAct.MyAtyDetail)


	tweb.RegisteHandler("/getMyPayment",memberAct.GetMyPayment)
	tweb.RegisteHandler("/toMyPaymentDetail",memberAct.ToMyPaymentDetail)
	tweb.RegisteHandler("/getMyPaymentDetail",memberAct.MyPaymentDetail)






}

func (i *MemberAct) ToMyAtyDetail(ctx *tweb.Ctx)  {
	recordId:=ctx.R.FormValue("recordId")
	data:=map[string]string{
		"recordId":recordId,
	}
	ctx.Render("myAtyDetail.html",data)
}
func (i *MemberAct) ToMyPaymentDetail(ctx *tweb.Ctx)  {
	recordId:=ctx.R.FormValue("recordId")
	data:=map[string]string{
		"recordId":recordId,
	}
	ctx.Render("myPaymentDetail.html",data)
}

func (i *MemberAct) ToMyAty(ctx *tweb.Ctx)  {
	ctx.Render("myAty.html",nil)
}
func (i *MemberAct) ToMyPayment(ctx *tweb.Ctx)  {
	ctx.Render("myPayment.html",nil)
}

func (i *MemberAct) ToUserInfo(ctx *tweb.Ctx)  {
	data:=map[string]interface{}{}
	_member:=ctx.GetSession_obj("member")
	userInfo:=ctx.GetSession_obj("userInfo")
	var member *entity.Member
	if _member!=nil {
		member=_member.(*entity.Member)
		data["member"]=member
	}
	if userInfo!=nil{
		data["userInfo"]=userInfo
	}
	ctx.Render("memberInfo.html",data)
}

func (i *MemberAct) MemberRegister(ctx *tweb.Ctx) {
	groupId:=ctx.R.FormValue("groupId")
	data:=map[string]string{
		"groupId":groupId,
	}
	ctx.Render("memberRegister.html",data)
}
func (i *MemberAct) MemberRegisterSubmit(ctx *tweb.Ctx) {
	msg:=map[string]string{}
	groupId:=ctx.R.FormValue("groupId")
	openid:=ctx.GetSession("openid")
	name:=ctx.R.FormValue("name")
	phone:=ctx.R.FormValue("phone")
	invitation_code:=ctx.R.FormValue("invitation_code")
	captcha_code:=ctx.R.FormValue("captcha_code")
	CheckCodeId:=ctx.GetSession("CheckCodeId")
	b:=captcha.VerifyString(CheckCodeId,captcha_code)
	if !b{
		log.Println("验证码错误..")
		msg["msg"]="captcha_code_error"
	}else{
		_groupId,_:=strconv.Atoi(groupId)
		group:=dao.FindGroupInfoById(_groupId)
		//fmt.Println("group:",group)
		b:=dao.SaveMember(name,phone,openid,(*group).GroupNum,(*group).Name,invitation_code)
		if b{
			msg["msg"]="success"
		}else{
			msg["msg"]="save_error"
		}
	}
	ctx.Render2JSON(msg)
}
func (i *MemberAct) SignUp(ctx *tweb.Ctx) {
	msg:=map[string]string{}
	_member:=ctx.GetSession_obj("member")
	_group:=ctx.GetSession_obj("group")
	openid:=ctx.GetSession("openid")
	var member *entity.Member
	var group *entity.Group
	if _member!=nil{
		member=_member.(*entity.Member)
	}else{
		if b,_member2:=MemberLogin(openid,ctx);b==true{
			member=_member2
		}else{

			msg["err_msg"]=entity.ERROR_MSG["IS_NOT_MEMBER"]
			ctx.Render2JSON(msg)
			return
		}
	}

	if _group!=nil{
		group=_group.(*entity.Group)
	}
	groupId:=ctx.R.FormValue("groupId")
	atyId:=ctx.R.FormValue("atyId")
	var _atyId,_groupId int
	if atyId!="" && groupId!=""{
		__atyId,_:=strconv.Atoi(atyId)
		__groupId,_:=strconv.Atoi(groupId)
		_atyId=int(__atyId)
		_groupId=int(__groupId)
		recordId,err:=dao.SaveAtyRecord(member.Id,_atyId,0.00,0.00,0,"",_groupId,group.Name)
		if err==nil{
			msg["msg"]="success"
			msg["recordId"]=strconv.Itoa(recordId)
		}else{
			msg["msg"]="fail"
		}
	}
	ctx.Render2JSON(msg)
}
func (i *MemberAct) CancelSignUp(ctx *tweb.Ctx)  {
	msg:=map[string]string{}
	recordId:=ctx.R.FormValue("recordId")
	atyId:=ctx.R.FormValue("atyId")

	if recordId!=""{
		_recordId,_:=strconv.Atoi(recordId)
		_atyId,_:=strconv.Atoi(atyId)

		_,message:=dao.CancelSignUp(_atyId,_recordId)
		msg["msg"]=message
		//fmt.Println(msg)
		ctx.Render2JSON(msg)
	}
}

func (i *MemberAct) GetMyAtyList(ctx *tweb.Ctx) {
	data:=map[string]interface{}{}
	openid:=ctx.GetSession("openid")
	_member:=ctx.GetSession_obj("member")
	_group:=ctx.GetSession_obj("group")
	//fmt.Println("_member:",_member," _group:",_group)
	if openid=="" || _member==""{
		data["msg"]="not_login" //用户未登陆
	}else{
		member:=_member.(*entity.Member)
		group:=_group.(*entity.Group)
		list:=dao.QueryMemberAtyRecordList(member.Id,group.Id)
		data["msg"]="success"
		data["list"]=list
	}
	//fmt.Println("list:",data["list"])
	ctx.Render2JSON(data)
}

func (i *MemberAct) GetMyPayment(ctx *tweb.Ctx) {
	data:=map[string]interface{}{}
	openid:=ctx.GetSession("openid")
	_member:=ctx.GetSession_obj("member")
	_group:=ctx.GetSession_obj("group")
	//fmt.Println("_member:",_member," _group:",_group)
	if openid=="" || _member==""{
		data["msg"]="not_login" //用户未登陆
	}else{
		member:=_member.(*entity.Member)
		group:=_group.(*entity.Group)
		list:=dao.QueryMyPaymentRecordList(member.Id,group.Id)
		data["msg"]="success"
		data["list"]=list
	}
	//fmt.Println("list:",data["list"])
	ctx.Render2JSON(data)
}



func (i *MemberAct) MyAtyDetail(ctx *tweb.Ctx) {
	_recordId:=ctx.R.FormValue("recordId")
	if _recordId!=""{
		recordId,_:=strconv.Atoi(_recordId)
		//fmt.Println("recordId:",recordId)
		data:=dao.QueryMyAtyDetail(recordId)
		//fmt.Println("data:",data)
		ctx.Render2JSON(data)
	}
}

func (i *MemberAct) MyPaymentDetail(ctx *tweb.Ctx) {
	_recordId:=ctx.R.FormValue("recordId")
	if _recordId!=""{
		recordId,_:=strconv.Atoi(_recordId)
		//fmt.Println("recordId:",recordId)
		data:=dao.QueryMyPaymentDetail(recordId)
		//fmt.Println("data:",data)
		ctx.Render2JSON(data)
	}
}






func MemberLogin(openid string,ctx *tweb.Ctx) (bool,*entity.Member) {
	log.Println("会员登录验证....openid:",openid)
	if openid==""{
		return false,nil
	}
	var id int
	var name,phone,_openid,password,group_num,group_name string
	var money float64
	var come_in_time string
	_sql:="select id,name,phone,openid,password,money, group_num,group_name,comein_time  from badminton_member m where m.openid=?"
	db.Query(func(sr *sql.Rows) error {
		for sr.Next(){
			sr.Scan(&id,&name,&phone,&_openid,&password,&money,&group_num,&group_name,&come_in_time)
		}
		return nil
	},_sql,openid)
	if name=="" && phone=="" && _openid=="" {
		log.Println("登录失败,openid:",openid,entity.ERROR_MSG["IS_NOT_MEMBER"])
		return false,nil

	}
	member:=&entity.Member{
		Id:id,
		Name:name,
		Phone:phone,
		OpenId:_openid,
		Password:password,
		Money:money,
		GroupNum:group_num,
		GroupName:group_name,
		ComeInTime:come_in_time,
	}
	ctx.PutSession_obj("member",member)
	log.Println("登录成功,member:",member)
	return true,member
}

func (c *MemberAct) Captcha(ctx *tweb.Ctx) {

	id := captcha.NewLen(4) //此id为生成验证码的ID，并不是实际显示的数字，在提交校验时，需要根据此ID进行查询。
	ctx.PutSession("CheckCodeId", id)
	w := ctx.W
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	captcha.WriteImage(w, id, 70, 30)
}
