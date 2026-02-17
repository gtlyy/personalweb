package controllers

import (
	"personalweb/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type ArticleController struct {
	web.Controller
}

func (c *ArticleController) Get() {
	cate := c.GetString("cate")
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("article").Filter("status", 2)
	if cate != "" {
		qs = qs.Filter("category", cate)
	}
	qs.OrderBy("-id").All(&articles)
	c.Data["Articles"] = articles
	c.Data["Cate"] = cate
	c.TplName = "article/list.tpl"
}

func (c *ArticleController) Detail() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	article := models.Article{Id: id}
	if err := o.Read(&article); err != nil || article.Status != 2 {
		c.Redirect("/article", 302)
		return
	}
	extensions := parser.CommonExtensions | parser.FencedCode
	p := parser.NewWithExtensions(extensions)
	opts := html.RendererOptions{Flags: html.CommonFlags}
	renderer := html.NewRenderer(opts)
	htmlBytes := markdown.ToHTML([]byte(article.ContentMd), p, renderer)
	article.ContentMd = string(htmlBytes)
	c.Data["Article"] = article
	c.TplName = "article/detail.tpl"
}
