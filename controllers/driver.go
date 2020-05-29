package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/models"
)

type DriverController struct {
	beego.Controller
}

func (c *DriverController) Get() {
	o := orm.NewOrm()
	qs := o.QueryTable("driver")
	qs.OrderBy("-created_time")
	var drivers []*models.Driver
	num, err := qs.All(&drivers)
	if err != nil {
		c.Abort(err.Error())
	}
	c.Data["Drivers"] = drivers
	c.Data["Number"] = num
	c.TplName = "driver/index.tpl"
}
