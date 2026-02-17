package controllers

import (
	"personalweb/models"
	"strconv"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type GameController struct {
	web.Controller
}

func (c *GameController) Get() {
	o := orm.NewOrm()
	var games []models.Game
	o.QueryTable("game").Filter("status", 2).OrderBy("-id").All(&games)
	c.Data["Games"] = games
	c.TplName = "games.tpl"
}

func (c *GameController) Play() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	game := models.Game{Id: id}
	if o.Read(&game) == nil && game.Status == 2 {
		c.Data["Game"] = game
		c.TplName = "game/play.tpl"
		return
	}
	c.Redirect("/games", 302)
}
