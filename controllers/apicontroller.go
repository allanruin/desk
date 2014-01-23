package controllers

import (
	"github.com/astaxie/beego"
	"sort"
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
	} else {
		this.Ctx.WriteString("ERROR_SIGNATURE")
	}

	// // 直接不校验返回echostr
	// this.Ctx.WriteString(echostr)
}

func (this *APIController) CheckSignature() bool {
	signature := this.GetString("signature")
	timestamp := this.GetString("timestamp")
	nonce := this.GetString("nonce")
	tmpArr := []string{Token, timestamp, nonce}
	sort.Strings(tmpArr)
	joinstr := strings.Join(tmpArr, "")
	sha1str := GetSha1FromStr(joinstr)

	if sha1str == signature {
		return true
	}

	return false
}
