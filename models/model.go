package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseTasks struct {
	Bid  int    `orm:"auto"`
	Name string `orm:"size(50)"`
}

type State struct {
	Wid   string `orm:"pk"`
	State string
}

func (b *BaseTasks) TableName() string {
	return "basetasks"
}

type Task struct {
	Tid   int `orm:"auto"`
	Taker string
	Asker string
	Taken int
}

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	// register model
	orm.RegisterModel(new(BaseTasks))
	orm.RegisterModel(new(State))
	orm.RegisterModel(new(Task))

	// set default database
	orm.RegisterDataBase("default", "mysql",
		// "hruser:hrisasimpleuser@/56.1.89.172?charset=utf8", 30)
		// "hruser:hrisasimpleuser@tcp(56.1.89.172:3306)/hrresume?charset=utf8", 30)
		"desk:deskisasimpleuser@tcp(localhost:3306)/desk?charset=utf8", 30)
}
