package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/helper"
	"goPlayer/models"
	"goPlayer/models/drivers"
	"io"
	"os"
	"strconv"
	"strings"
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

func renderFile(c *VideoController, fullPath string, stat os.FileInfo) {
	mode := c.GetString("mode")
	file, err3 := os.OpenFile(fullPath, os.O_RDONLY, 06660)
	fileExt := helper.ArrayEnd(strings.Split(stat.Name(), "."))
	if err3 != nil {
		c.Abort(err3.Error())
		return
	}
	defer file.Close()
	size := strconv.FormatInt(stat.Size(), 10)
	sizeLess1 := strconv.FormatInt(stat.Size()-1, 10)
	// 区域下载
	getRange := c.Ctx.Input.Header("Range")
	if getRange != "" && strings.HasPrefix(getRange, "bytes=") {
		start, err4 := strconv.ParseInt(strings.Split(getRange[len("bytes="):], "-")[0], 10, 64)
		if err4 == nil {
			if start > 0 {
				_, err5 := file.Seek(start, 0)
				if err5 == nil {
					//c.Ctx.Output.SetStatus(206)
					c.Ctx.Output.Header("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+sizeLess1+"/"+size)
				}
			}
		}
	}
	if mode == "download" {
		c.Ctx.Output.Header("Content-Disposition", "form-data; name=\"attachment\"; filename=\""+stat.Name()+"\"")
	}
	c.Ctx.Output.ContentType(fileExt)
	c.Ctx.Output.Header("Content-Length", size)
	c.Ctx.Output.Header("Accept-Ranges", "bytes")
	if c.Ctx.ResponseWriter.Header().Get("Content-Range") != "" {
		//c.Ctx.Output.Header("Content-Type", "application/octet-stream")
		c.Ctx.ResponseWriter.WriteHeader(206)
	}
	tempSize := 1024
	buffer := make([]byte, tempSize)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if n == tempSize {
			_, _ = c.Ctx.ResponseWriter.Write(buffer)
		} else {
			_, _ = c.Ctx.ResponseWriter.Write(buffer[0:n])
		}
	}
}
