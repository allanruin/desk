package models

import (
	// "github.com/astaxie/beego"
	// "github.com/allanruin/beego/orm"
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
	State string
}

type User struct {
	Id         int    `orm:"auto"`
	Wid        string `orm:"key"`
	Name       string `valid:"Required"`
	HrCode     string `orm:"column(hrcode)",valid:"Required"`
	Department string
	Position   string
	State      string
	Phone      string
	Tail       string `valid:"Required"`
}

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	// register model
	orm.RegisterModel(new(BaseTasks))
	orm.RegisterModel(new(State))
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql",

		"desk:sdfsjfksjdfks@tcp(localhost:3306)/desk?charset=utf8", 30)
}
