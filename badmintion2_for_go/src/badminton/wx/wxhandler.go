package wx

import "tweb/wx"

func (wact *WeiXinAct) Echo(w wx.ResponseWriter, r *wx.Request) {
	txt := r.Content          // 获取用户发送的消息
	w.ReplyText(txt)          // 回复一条文本消息
}
