package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "personalweb/routers"
)

func main() {
	web.Run()
}
