package controllers

import (
	"personalweb/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	web.Controller
}

func (c *HomeController) Get() {
	o := orm.NewOrm()
	var articles []models.Article
	o.QueryTable("article").Filter("status", 2).OrderBy("-id").Limit(6).All(&articles)
	c.Data["Articles"] = articles
	c.TplName = "index.tpl"
}
