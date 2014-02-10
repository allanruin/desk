package controllers

import (
	"desk/models"
	// "errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	// "github.com/allanruin/beego/orm"
	"log"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["Wid"] = this.GetString("wid")
	this.TplNames = "register.tpl"
}

type registForm struct {
	Code     string
	UserName string
	Tail     string
}

func (this *RegisterController) Post() {
	err := this.Validate()
	if err != nil {
		beego.Error("验证不通过" + err.Error())
		this.Data["verror"] = "验证不通过," + err.Error()
		this.TplNames = "register.tpl"
		return
	}

	this.TplNames = "registersuccess.tpl"
}

func (this *RegisterController) Validate() error {
	// r := registForm{}
	user := &models.User{}
	if err := this.ParseForm(user); err != nil {
		//handle error
		return err
	}

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	valid := validation.Validation{}
	b, err := valid.Valid(user)

	if err != nil {
		return err
	}
	if !b {
		// validation does not pass
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	err = o.Read(user, "HrCode", "Name", "Tail")
	beego.Info(user)

	if err != nil {
		return err
	}

	user.Wid = this.GetString("wid")
	user.State = "subscribed"
	_, err = o.Update(user)
	if err != nil {
		beego.Trace("更新失败", user)
		return err
	}

	return nil
}
