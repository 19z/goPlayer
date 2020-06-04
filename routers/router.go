package routers

import (
	"github.com/astaxie/beego"
	"goPlayer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/driver", &controllers.DriverController{})
	beego.Router("/files/*", &controllers.FilesController{})
	beego.Router("/video/*", &controllers.VideoController{})
}
