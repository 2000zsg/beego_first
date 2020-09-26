package main

import (
	"github.com/astaxie/beego"
	_ "hl/routers"
	_"hl/db_mysql"
)

func main() {
	beego.Run()
}

