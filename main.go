package main

import (
	"os"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "personalweb/routers"
	"personalweb/utils"
)

func main() {
	// 初始化日志
	if err := utils.InitLogger("./logs/app.log"); err != nil {
		os.Stderr.WriteString("Failed to initialize logger: " + err.Error() + "\n")
	}

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "X-XSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	web.Run()
}
