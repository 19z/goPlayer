package routers

import (
	"goPlayer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/driver", &controllers.DriverController{})
	beego.Router("/files/*", &controllers.FilesController{})
    beego.Router("/video/*", &controllers.VideoController{})
}
