package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "myblogtest/routers"
	"myblogtest/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

