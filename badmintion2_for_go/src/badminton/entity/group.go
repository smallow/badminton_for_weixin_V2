package entity



type Group struct {
	Id  int `json:"id"` //ID
	Name string `json:"name"`
	GroupNum string `json:"group_num"`  //群号
	GroupType string `json:"group_type"` //群类型  weixin 或者 qq
	GroupMaster string `json:"group_master"`  //群主
	GroupMasterPhone string `json:"group_master_phone"`//群主手机号
	Admins [] Member `json:"admins"`  //管理员

}



type Activity struct {
	Id int `json:"id"`
	Address string `json:"address"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Date      string `json:"date"`
	ChargeMemberId int `json:"charge_member_id"`
	ChargeMemberName string `json:"charge_member_name"`
	ChargeMemberPhone string `json:"charge_member_phone"`
	TotalCost  float64 `json:"total_cost"`
	BadmintonNum int  `json:"badminton_num"`
	SiteNum      int  `json:"site_num"`
	TimeNum      int  `json:"time_num"`
	AvgCost       float64 `json:"avg_cost"`
	TotalPerson   int  `json:"total_person"`
	GroupId      int `json:"group_id"`
	Status        string `json:"status"`
}

type Member struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	OpenId string `json:"open_id"`
	Password string `json:"password"`
	Money float64 `json:"money"`
	GroupNum string `json:"group_num"`
	GroupName string `json:"group_name"`
	ComeInTime string `json:"come_in_time"`  //入会时间
	InvitationCode string `json:"invitation_code"` //邀请码
}

type ActivityRecord struct {
	Id  int  `json:"id"`
	Member Member  `json:"member"`   //活动参加人
	Activity  Activity  `json:"activity"`       //活动
	Money    float64    `json:"money"`          //消费
	CurrentDayLeftMoney  float64 `json:"current_day_left_money"` //当天活动之后的余额
	FriendNum  int  `json:"friend_num"`         //携带人数
	FriendNames string  `json:"friend_names"`   //携带人名称
	SignUpTime  string  `json:"sign_up_time"`   //报名时间
	UpdateTime  string  `json:"update_time"`    //更新时间
	IsCancel    bool    `json:"is_cancel"`      //是否取消
	CancelTime  string  `json:"cancel_time"`    //取消时间
	Memo        string  `json:"memo"`           //备注
}
type PaymentRecord struct {
	Id  int `json:"id"`
	Member Member `json:"member"`  //会员
	Group Group `json:"group"`  //所属群
	PaymentMoney float64 `json:"payment_money"` //缴费金额
	PaymentType string `json:"payment_type"`    //缴费渠道
	BillNumber   string `json:"bill_number"`      //转账单号
	PaymentTime string `json:"payment_time"`    //转账时间
	Memo  string `json:"memo"` //备注
	PayeeName string `json:"payee_name"` //收款人姓名
	PayeeAccount string `json:"payee_account"` //收款人账号
	CreateTime string `json:"create_time"`  //入库时间
}

