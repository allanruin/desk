package controllers

import (
	"github.com/astaxie/beego/orm"
)

type DeskOrm struct {
	orm.Ormer
}

func NewDeskOrm() DeskOrm {
	o := orm.NewOrm()
	o.Using("default")

	do := DeskOrm{o}
	return do
}

func (o DeskOrm) ReadByExample(v interface{}) {

}
