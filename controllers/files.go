package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/helper"
	"goPlayer/models"
	"goPlayer/models/drivers"
)

type FilesController struct {
	beego.Controller
}

func (c *FilesController) Get() {
	splat := c.Ctx.Input.Param(":splat")
	driverId, path, err := helper.SplatUrl(splat)
	if err != nil {
		c.Abort(err.Error())
		return
	}
	o := orm.NewOrm()
	driver := models.Driver{Id: *driverId}
	err = o.Read(&driver)
	if err != nil {
		c.Abort(err.Error())
		return
	}
	d := driver.GetDriver()
	fileItem, err := d.GetFileItem(path)
	if err != nil {
		c.Abort(err.Error())
		return
	}
	mode := c.GetString("mode")
	//println(&driverId, driver.Name, driver.Path, path)
	if fileItem.IsDirectory {
		files := d.DirList(path)
		c.Data["Files"] = files
		c.Data["Driver"] = driver
		c.Data["Length"] = len(files)
		parent := helper.Dirname(path)
		if parent == "" || parent == "." {
			c.Data["Parent"] = nil
		} else {
			c.Data["Parent"], _ = d.GetFileItem(parent)
		}

		if mode != "" {
			result := struct {
				Files    []drivers.FileItem
				Length   int
				DriverId int
			}{files, len(files), driver.Id}
			if mode == "json" {
				c.Data["json"] = &result
				c.ServeJSON()
			} else if mode == "jsonp" {
				c.Data["jsonp"] = &result
				c.ServeJSONP()
			}
			return
		}
		c.TplName = "files/index.tpl"
	} else {
		fileItem.SendFile(c.Ctx)
	}
}
