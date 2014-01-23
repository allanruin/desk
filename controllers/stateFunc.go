package controllers

import (
	"desk/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "strings"
)

type StateFunc func(rc *RecvTextMsg) *RetTextMsg

var StateFuncs map[string]StateFunc = map[string]StateFunc{
	"timeout": ProcessTimeout,
	"init":    ProcessInit,
	// "list":      ProcessList,
	"working":   ProcessWorking,
	"othertask": ProcessOtherTask,
	"asked":     ProcessAsked,
	"cartridge": ProcessCartridge,
}

var backtpl string = "【返回】请回复0"

func (rc *RecvTextMsg) SaveState(state string) {
	o := orm.NewOrm()
	o.Using("default")
	nstate := models.State{Wid: rc.FromUserName, State: state}
	_, err := o.Update(&nstate)
	if err != nil {
		beego.Error("save new state error:" + err.Error())
	}
}

func ProcessInit(rc *RecvTextMsg) *RetTextMsg {
	s := "a"
	content := rc.Content
	switch content {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
	case "7":
	case "8":
	case "9":
	default:
		s = "unknown"
	}

	return rc.Return(s)
}

func ProcessTimeout(rc *RecvTextMsg) *RetTextMsg {
	s := "您好，欢迎试用广东人保桌面设备助手，请输入以下功能编号选择桌面任务:\n"

	o := orm.NewOrm()
	o.Using("default")
	var bts []*models.BaseTasks
	_, err := o.QueryTable("basetasks").All(&bts)
	// fmt.Printf("Returned Rows Num: %s, %s", num, err)
	if err != nil {
		s = s + "数据库出错，无法获取列表"
		return rc.Return(s)
	}
	for _, bt := range bts {
		s = s + fmt.Sprintf("[%d] %s\n", bt.Bid, bt.Name)
	}

	rc.SaveState("init")

	return rc.Return(s)
}

// func ProcessList(rc *RecvTextMsg) *RetTextMsg {
// 	s := "默认回复"
// 	return rc.Return(s)
// }

func ProcessWorking(rc *RecvTextMsg) *RetTextMsg {
	s := "默认回复"
	return rc.Return(s)
}

func ProcessOtherTask(rc *RecvTextMsg) *RetTextMsg {
	s := "默认回复"
	return rc.Return(s)
}

func ProcessAsked(rc *RecvTextMsg) *RetTextMsg {
	content := Filter(rc.Content)
	// 有工单在处理时的默认回复
	s := "您的桌面服务请求工单(%d)正在处理中，请耐心等候"

	switch content {
	case "ok":
	case "urge":
	case "list":
	case "help":
	default:
		s = "无法识别的回复。\n您的桌面服务请求工单(%d)正在处理中,如果工单已完成请回复ok结束桌面工单任务。"
	}

	return rc.Return(s)
}

func ProcessCartridge(rc *RecvTextMsg) *RetTextMsg {
	tpl := "系统已记录下要更换的硒鼓型号%s,请耐心等候桌面服务人员上门更换\n%s"
	content := Filter(rc.Content)
	s := fmt.Sprintf(tpl, content, backtpl)
	return rc.Return(s)
}
