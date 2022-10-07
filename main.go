package main

import (
	"github.com/astaxie/beego"
	_ "myblogdemo/routers"
	"myblogdemo/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
