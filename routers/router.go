package routers

import (
	"github.com/astaxie/beego"
	"goPlayer/controllers"
	_ "goPlayer/models"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/driver", &controllers.DriverController{})
	beego.Router("/files/:id([0-9]+)/:splat(.+)", &controllers.FilesController{})
}
