package controllers

import (
	"personalweb/models"
	"strconv"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type ToolController struct {
	web.Controller
}

func (c *ToolController) Get() {
	o := orm.NewOrm()
	var tools []models.Tool
	o.QueryTable("tool").Filter("status", 2).OrderBy("-id").All(&tools)
	c.Data["Tools"] = tools
	c.TplName = "tools.tpl"
}

func (c *ToolController) Use() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	tool := models.Tool{Id: id}
	if o.Read(&tool) == nil && tool.Status == 2 {
		c.Data["Tool"] = tool
		c.TplName = "tool/use.tpl"
		return
	}
	c.Redirect("/tools", 302)
}
