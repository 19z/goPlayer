package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goPlayer/helper"
	"goPlayer/models"
	"io"
	"io/ioutil"
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
		c.GetThumbnail(fullPath, splat)
		return
	}

	c.Data["Path"] = p
	c.Data["Splat"] = splat
	//c.Data["Url"] = beego.URLFor("FilesController.Get", ":splat", splat);
	//c.Data["VideoMethod"] = beego.URLFor("VideoController.Get", ":splat", splat);
	c.Data["Stat"] = stat
	c.TplName = "files/player.tpl"
	return
}

func (c *VideoController) GetThumbnail(fullPath string, splat string) {
	basename := helper.Basename(fullPath)
	dirname := helper.PathJoin(helper.TempDir(), splat)
	savePath := helper.PathJoin(dirname, "thumbnail.jpg")
	if helper.FileExist(savePath) {
		c.Ctx.Output.Download(savePath, basename+"_thumbnail.jpg")
		return
	}
	if !helper.FileExist(dirname) {
		_ = helper.MakeDir(dirname)
	}
	ffmpeg := helper.FindFFmpeg()
	if ffmpeg == "" {
		c.Abort("not ffmpeg")
		return
	}
	println(savePath)
	cmd := exec.Command(ffmpeg, "-hide_banner", "-i", fullPath, "-ss", "100", "-y", "-t", "0.001", "-f", "image2", savePath)
	//cmd := exec.Command("C:\\MineProgram\\ffmpeg\\bin\\ffmpeg.exe", "-hide_banner")
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	b, _ := ioutil.ReadAll(stderr)
	println(string(b))

	if helper.FileExist(savePath) {
		c.Ctx.Output.Download(savePath, basename+"_thumbnail.jpg")
		return
	} else {
		c.Ctx.WriteString(string(b))
	}
}

func printCmd(cmd *exec.Cmd) {
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	b, err := ioutil.ReadAll(stderr)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(b))
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}
func readFile(file io.Reader) {
	tempSize := 10
	buffer := make([]byte, tempSize)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if n == tempSize {
			print(string(buffer))
		} else {
			print(string(buffer[0:n]))
		}
	}
}
