package main

import (
	"github.com/astaxie/beego"
	_ "movienotes/gosession"
	_ "movienotes/routers"
	_ "movienotes/settings"
)

func main() {
	beego.Run()
}
