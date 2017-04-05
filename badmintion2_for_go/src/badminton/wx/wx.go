package wx

import (
	"io/ioutil"
	"tweb"
	"fmt"
	"net/http"
	"tweb/wx"
	"log"
	"encoding/json"
)

//微信
type WeiXinAct struct {
	Appid  string
	Secert string
	Token  string
	SsoUrl string //单点登录地址
}

//
func InitWeixinSdk(urlprefix, appid, secert, token, ssourl string) {
	wxact := WeiXinAct{
		Appid:  appid,
		Secert: secert,
		Token:  token,
		SsoUrl: ssourl,
	}
	//微信初始化

	tweb.InitWeixin(urlprefix, token, appid, secert,nil)
	//绑定action
	tweb.RegisteHandler("/wx/createmenu", wxact.CreateWxMenu)
	tweb.RegisteHandler("/wx/sso", wxact.Sso)
	//微信其他消息处理，如：文本回复，图像上传事件，扫码事件处理等
	tweb.WxMux.HandleFunc(wx.MsgTypeText, wxact.Echo)
	//微信菜单URL跳转短地址
	//tweb.RegisteHandler("/wx/s/jbts",wxact.Short)
	//tweb.RegisteHandler("/wx/s/tscx",wxact.Short)

}

//创建微信菜单
func (w *WeiXinAct) CreateWxMenu(ctx *tweb.Ctx) {
	menu := &wx.Menu{make([]wx.MenuButton, 3)}
	menu.Buttons[0].Name = "羽球集市"
	menu.Buttons[0].Type = wx.MenuButtonTypeUrl
	menu.Buttons[0].Url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + w.Appid + "&redirect_uri=" + "http://webwhd.qmx.top/wx/sso" + "&response_type=code&scope=snsapi_base&state=yqjs#wechat_redirect"
	//
	menu.Buttons[1].Name = "赛事新闻"
	menu.Buttons[1].Type = wx.MenuButtonTypeUrl
	menu.Buttons[1].Url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + w.Appid + "&redirect_uri=" + "http://webwhd.qmx.top/wx/sso" + "&response_type=code&scope=snsapi_base&state=ssxw#wechat_redirect"

	menu.Buttons[2].Name = "发现"
	menu.Buttons[2].SubButtons = make([]wx.MenuButton, 3)
	menu.Buttons[2].SubButtons[0].Name = "群活动"
	menu.Buttons[2].SubButtons[0].Type = wx.MenuButtonTypeUrl
	menu.Buttons[2].SubButtons[0].Url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + w.Appid + "&redirect_uri=" + "http://smallow.top/wx/sso" + "&response_type=code&scope=snsapi_base&state=group_activity#wechat_redirect"

	menu.Buttons[2].SubButtons[1].Name = "我的活动"
	menu.Buttons[2].SubButtons[1].Type = wx.MenuButtonTypeUrl
	menu.Buttons[2].SubButtons[1].Url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + w.Appid + "&redirect_uri=" + w.SsoUrl + "&response_type=code&scope=snsapi_base&state=my_activity#wechat_redirect"

	menu.Buttons[2].SubButtons[2].Name = "会员信息"
	menu.Buttons[2].SubButtons[2].Type = wx.MenuButtonTypeUrl
	menu.Buttons[2].SubButtons[2].Url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + w.Appid + "&redirect_uri=" + w.SsoUrl + "&response_type=code&scope=snsapi_base&state=member_info#wechat_redirect"

	err := tweb.WxMux.CreateMenu(menu)
	if err != nil {
		fmt.Fprint(ctx.W, err.Error())
	} else {
		fmt.Fprint(ctx.W, "ok")
	}
}

//单点登录处理，微信服务器会回调此地址
func (w *WeiXinAct) Sso(ctx *tweb.Ctx) {
	code, state := ctx.R.FormValue("code"), ctx.R.FormValue("state")
	if code == "" {
		return
	}
	//取得openid
	urlstr := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + w.Appid + "&secret=" + w.Secert + "&code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(urlstr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	//
	data := make(map[string]interface{})
	json.Unmarshal(bs, &data)
	//
	if v, ok := data["openid"]; ok {
		openid := v.(string)
		//存储到session
		ctx.PutSession("openid", openid)
		ctx.Save2Cookie() //更新客户端cookie
		jumpByState(ctx,state)

	} else {
		log.Println("读取openid失败")
	}
}
//短地址跳转
func (w *WeiXinAct)Short(ctx *tweb.Ctx){
	log.Println("aaaa")
	uri:=ctx.R.RequestURI
	var state string //根据不同场景，进行跳转
	fmt.Sscanf(uri,"/wx/s/%s",&state)
	log.Println(uri,state)
	//
	openid:=ctx.GetSession("openid")
	if openid=="" { //未登陆的情况
		log.Println("调用微信oauth2")
		sso_url := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + w.Appid + "&redirect_uri=" + w.SsoUrl + "&response_type=code&scope=snsapi_base&state="+state+"#wechat_redirect"
		http.Redirect(ctx.W,ctx.R,sso_url,301)
	}else{//已经登陆，则直接跳转
		log.Println("已经知道openid",openid,state)
		jumpByState(ctx,state)
	}
}
//取到openid后，根据state跳转
func jumpByState(ctx *tweb.Ctx,state string){
	log.Println("单点登录跳转,state:",state)
	switch state{
	case "group_activity":
		ctx.PutSession("act","群活动")
		ctx.Save2Cookie()
		http.Redirect(ctx.W,ctx.R,"/groupAty",301)
	case "my_activity":
		ctx.PutSession("act","我的活动")
		ctx.Save2Cookie()
		http.Redirect(ctx.W,ctx.R,"/myAty",301)
	case "member_info":
		ctx.PutSession("act","会员信息")
		ctx.Save2Cookie()
		http.Redirect(ctx.W,ctx.R,"/memberInfo",301)
	default:
		http.Redirect(ctx.W,ctx.R,"/",301)
	}
}
