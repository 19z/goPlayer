package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/helper"
	"goPlayer/models"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type VideoController struct {
	beego.Controller
}

func (c *VideoController) Get() {
	splat := c.Ctx.Input.Param(":splat")

	driverId, p, err := helper.SplatUrl(splat)
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
	if driver.DriverType != "local" {
		c.Abort("Only local can use")
		return
	}
	fullPath := helper.PathJoin(driver.Path, p)
	stat, err := os.Stat(fullPath)
	if err != nil {
		c.Abort(err.Error())
		return
	}
	if stat.IsDir() {
		c.Redirect("/files/"+splat, 302)
		return
	}

	mode := c.GetString("mode")
	if mode == "pic" {
		c.GetThumbnail(fullPath)
		return
	}

	c.Data["Path"] = p;
	c.Data["Splat"] = splat;
	//c.Data["Url"] = beego.URLFor("FilesController.Get", ":splat", splat);
	//c.Data["VideoMethod"] = beego.URLFor("VideoController.Get", ":splat", splat);
	c.Data["Stat"] = stat
	c.TplName = "files/player.tpl"
	return
}

func (c *VideoController) _GetThumbnail(fullPath string) {
	savePath := fullPath + ".jpg"
	ffmpeg := helper.FindFFmpeg()
	if ffmpeg == "" {
		c.Abort("not ffmpeg")
		return
	}
	cmd := exec.Command(ffmpeg, "-i ", fullPath, "-y", "-f", "image2", "-ss", "10", savePath)
	stdout, err := cmd.StdoutPipe();
	if err != nil {
		//获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(stdout)
	println(string(b))
	c.Ctx.WriteString(string(b))
	err = cmd.Wait()
	if err != nil {
		log.Printf(err.Error())
	}
}

func (c *VideoController) GetThumbnail(fullPath string) {
	savePath := fullPath + ".jpg"
	ffmpeg := helper.FindFFmpeg()
	if ffmpeg == "" {
		c.Abort("not ffmpeg")
		return
	}
	println(savePath)
	//cmd := exec.Command(ffmpeg, "-i ", fullPath, "-y", "-f", "image2", "-ss", "10", savePath)
	cmd := exec.Command(ffmpeg)
	stdout, err := cmd.StdoutPipe() //指向cmd命令的stdout
	cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))

	c.Ctx.WriteString(string(content))
}
