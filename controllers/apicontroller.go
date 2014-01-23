package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/log"
	"desk/models"
	"strings"
)

type APIController struct {
	beego.Controller
}

func (this *APIController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
	echostr := this.GetString("echostr")

	// 校验后选择是否返回echostr
	if this.CheckSignature() {
		this.Ctx.WriteString(echostr)
		beego.Info("返回微信的echostr: " + echostr)
	} else {
		this.Ctx.WriteString("ERROR_SIGNATURE")
		beego.Info("signature 不正确 ")
	}

	// // 直接不校验返回echostr
	// this.Ctx.WriteString(echostr)
	// beego.Info("返回微信的echostr: " + echostr)
}

func (this *APIController) Post() {
	body := this.Ctx.Input.RequestBody
	bodystr := BtyeToString(body)
	beego.Trace("body: ", bodystr)

	rc := new(RecvTextMsg)
	xml.Unmarshal(body, rc)

	// beego.Trace("received struct", rc)

	// rt := rc.Return("you said : " + rc.Content)

	state := GetState(rc.FromUserName)

	rt := rc.Process(state)

	// rbuf, err := xml.Marshal(rt)
	// if err != nil {
	// 	beego.Error("marshal err:" + err.Error())
	// } else {
	// 	str := BtyeToString(rbuf)
	// 	beego.Info("return xml:" + str)
	// }

	// beego.Trace("return struct", rt)

	this.Data["xml"] = rt
	this.ServeXml()
}

// 获取调用微信的一些特定接口所需要的ACCESS_TOKEN，有效期为2小时
func (this *APIController) GetAccessToken() {

}

func GetState(wid string) string {
	if strings.TrimSpace(wid) == "" {
		beego.Error("查找了空的wid")
		return "init"
	}

	o := orm.NewOrm()
	o.Using("default")
	state := models.State{Wid: wid}
	err := o.Read(&state)

	if err != nil {
		beego.Error("找不到用户的state")
		return "init"
	}
	return state.State
}
