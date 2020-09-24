package routers

import (
	"hl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.MainController{})
}
