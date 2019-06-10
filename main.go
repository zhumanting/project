package main

import (
	_ "api/routers"
	"api/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// init mysql
	utils.RegisterDB()
	orm.Debug = true
	// redis use
	//redis := utils.RedisConnection()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
