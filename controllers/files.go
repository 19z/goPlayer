package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/models"
	"strconv"
	"strings"
)

type FilesController struct {
	beego.Controller
}

func (c *FilesController) Get() {
	splat := c.Ctx.Input.Param(":splat")
	if splat[0] == '~' {
		splat = strings.Trim(splat[1:], "/\\")
	}
	driverId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Abort(err.Error())
		return
	}
	o := orm.NewOrm()
	driver := models.Driver{Id: driverId}
	println(driver.Path)
	err = o.Read(&driver)
	if err != nil {
		c.Abort(err.Error())
		return
	}
	d := driver.GetDriver()
	files := d.DirList(splat)
	for _, file := range files {
		println(file.Name)
	}
	c.Data["Files"] = files
	c.Data["Driver"] = driver
	c.Data["Length"] = len(files)
	c.TplName = "files/index.tpl"
}
