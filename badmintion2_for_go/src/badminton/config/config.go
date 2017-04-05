package config

import (
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Config struct {
	WxPrefix   string `json:"wxprefix"`
	WxAppid    string `json:"wxappid"`
	WxToken    string `json:"wxtoken"`
	WxSecert   string `json:"wxsecert"`
	WxSsourl   string `json:"wxssourl"`
	IsHttps    bool `json:"ishttps"`
	Pemfile    string `json:"pemfile"`
	Keyfile    string `json:"keyfile"`
	Listenaddr string `json:"listenaddr"`
	FileServer string `json:"file_server"`
	DbInfo 	   string `json:"db_info"`
}


//
func ReadSysConfig()*Config {
	var sysconfig Config
	fi, err := os.Open("./res/config.json")
	if err!=nil{
		log.Println(err.Error())
	}
	defer fi.Close()
	bs, err := ioutil.ReadAll(fi)
	if err!=nil{
		log.Println(err.Error())
	}
	err=json.Unmarshal(bs, &sysconfig)
	if err!=nil{
		log.Println(err.Error())
	}
	bs,_=json.Marshal(sysconfig)
	log.Println("读取系统配置成功,配置信息:",string(bs))
	return &sysconfig
}

