package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/models"
	"strconv"
)

type DriverController struct {
	beego.Controller
}

type DriveForm struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
	Type string `form:"type"`
	Path string `form:"path"`
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

func (c *DriverController) Post() {
	d := DriveForm{}
	if err := c.ParseForm(&d); err != nil {
		c.Abort(err.Error())
		return
	}
	o := orm.NewOrm()
	var driver models.Driver
	if d.Id > 0 {
		driver.Id = d.Id
		if err := o.Read(&driver); err != nil {
			c.Abort(err.Error())
			return
		}
	}
	driver.Name = d.Name
	driver.DriverType = d.Type
	driver.Path = d.Path
	if err := driver.CheckData(); err != nil {
		c.Abort(err.Error())
		return
	}
	if d.Id > 0 {
		if _, err := o.Update(&driver); err == nil {
			c.Redirect("/files/"+strconv.Itoa(driver.Id)+"/~/", 302)
		}
	} else {
		if _, err := o.Insert(&driver); err == nil {
			c.Redirect("/files/"+strconv.Itoa(driver.Id)+"/~/", 302)
		}
	}
}
