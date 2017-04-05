package main

import (
	"tweb"
	"badminton/config"
	"badminton/wx"
	"log"
	_ "badminton/action"
	_"github.com/go-sql-driver/mysql"
	"tweb/db"
)

var sc *config.Config

func init() {
	//读取系统配置
	sc = config.ReadSysConfig()
	//静态资源
	tweb.InitStatic("/res", "./static")
	//加载指定目录下的所有模板，要引用的子模板，需要放在inc目录下，其他关于模板的使用，与go原生一致
	tweb.LoadAllTemplates("./static/templates")
	//扩展模板方法，其中入参，出参，都可以任意写，此处是可选
	tweb.AddTemplateFunc("MyExtFn", func(name string) interface{} {
		return "hello " + name
	})
	//微信初始化
	if sc.WxAppid != "" && sc.WxToken != "" && sc.WxSecert != "" {
		wx.InitWeixinSdk(sc.WxPrefix, sc.WxAppid, sc.WxSecert, sc.WxToken, sc.WxSsourl)
		log.Println("微信 init success...")

	}
	if sc.DbInfo!=""{
		//util.InitMySql(sc.User, sc.Password, sc.Ip, sc.DbName)
		db.InitDbPool("mysql",sc.DbInfo,200)
		log.Println("数据库 init success...")
	}
}

func main() {
	// https服务
	if sc.IsHttps {
		tweb.StartSSLWeb(sc.Listenaddr, sc.Pemfile, sc.Keyfile)
	} else {
		tweb.StartWeb(sc.Listenaddr)
	}
}
