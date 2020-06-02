package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	var _d models.Driver
	err := o.QueryTable("driver").One(&_d)
	if err != nil {
		err := orm.RunSyncdb("default", true, true)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect("/install", 302)
		return
	}
	c.Redirect("/driver", 302)
}
