package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"os"
	"personalweb/models"
	"personalweb/utils"
	"strconv"
	"time"
)

type AdminController struct {
	web.Controller
}

func (c *AdminController) Prepare() {
	path := c.Ctx.Request.URL.Path
	exclude := []string{"/admin/login", "/admin/logout"}
	for _, p := range exclude {
		if path == p {
			return
		}
	}
	// 修复：正确断言为 models.Admin
	admin, ok := c.GetSession("admin").(models.Admin)
	if !ok {
		c.Redirect("/admin/login", 302)
		c.StopRun()
		return
	}
	c.Data["Admin"] = admin
}

// 登录
func (c *AdminController) Login() {
	c.TplName = "admin/login.tpl"
}

func (c *AdminController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")

	o := orm.NewOrm()
	var admin models.Admin // 确保 models.Admin 可访问
	err := o.QueryTable("admin").Filter("username", username).One(&admin)
	if err != nil {
		c.Data["Msg"] = "账号不存在"
		c.TplName = "admin/login.tpl"
		return
	}
	if !admin.CheckPassword(password) {
		c.Data["Msg"] = "密码错误"
		c.TplName = "admin/login.tpl"
		return
	}

	c.SetSession("admin", admin)
	c.Redirect("/admin/index", 302)
}

// 退出
func (c *AdminController) Logout() {
	c.DestroySession()
	c.Redirect("/admin/login", 302)
}

// 文章列表
func (c *AdminController) Index() {
	o := orm.NewOrm()
	var articles []models.Article
	o.QueryTable("article").OrderBy("-id").All(&articles)
	c.Data["Articles"] = articles
	c.TplName = "admin/index.tpl"
}

// 新增文章
func (c *AdminController) Add() {
	c.TplName = "admin/add.tpl"
}

func (c *AdminController) DoAdd() {
	title := c.GetString("title")
	category := c.GetString("category")
	content := c.GetString("content")
	status, _ := c.GetInt("status")

	o := orm.NewOrm()
	o.Insert(&models.Article{
		Title:     title,
		Category:  category,
		ContentMd: content,
		Status:    status,
	})
	c.Redirect("/admin/index", 302)
}

// 编辑文章（修复 ORM Read 传参）
func (c *AdminController) Edit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	var art models.Article
	art.Id = id
	if err := o.Read(&art); err != nil {
		c.Redirect("/admin/index", 302)
		return
	}
	c.Data["Article"] = art
	c.TplName = "admin/edit.tpl"
}

func (c *AdminController) DoEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title := c.GetString("title")
	category := c.GetString("category")
	content := c.GetString("content")
	status, _ := c.GetInt("status")

	o := orm.NewOrm()
	art := models.Article{Id: id}
	if o.Read(&art) == nil {
		art.Title = title
		art.Category = category
		art.ContentMd = content
		art.Status = status
		o.Update(&art)
	}
	c.Redirect("/admin/index", 302)
}

// 删除文章
func (c *AdminController) Del() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Delete(&models.Article{Id: id})
	c.Redirect("/admin/index", 302)
}

// 修改密码
func (c *AdminController) PasswordPage() {
	c.TplName = "admin/password.tpl"
}

func (c *AdminController) ChangePassword() {
	oldPwd := c.GetString("old_password")
	newPwd := c.GetString("new_password")
	confirmPwd := c.GetString("confirm_password")

	if newPwd != confirmPwd {
		c.Data["Msg"] = "两次密码不一致"
		c.TplName = "admin/password.tpl"
		return
	}
	// 修复：正确断言 models.Admin
	adminObj, ok := c.GetSession("admin").(models.Admin)
	if !ok {
		c.DestroySession()
		c.Redirect("/admin/login", 302)
		return
	}
	o := orm.NewOrm()
	var admin models.Admin // 确保 models.Admin 可访问
	err := o.QueryTable("admin").Filter("id", adminObj.Id).One(&admin)
	if err != nil {
		c.Data["Msg"] = "用户异常"
		c.TplName = "admin/password.tpl"
		return
	}
	if !admin.CheckPassword(oldPwd) {
		c.Data["Msg"] = "旧密码错误"
		c.TplName = "admin/password.tpl"
		return
	}
	admin.Password = newPwd
	admin.EncryptPassword()
	o.Update(&admin, "Password")
	c.DestroySession()
	c.Data["Msg"] = "修改成功，请重新登录"
	c.TplName = "admin/login.tpl"
}

// ==================== 游戏管理 ====================
func (c *AdminController) GameList() {
	o := orm.NewOrm()
	var games []models.Game
	o.QueryTable("game").OrderBy("-id").All(&games)
	c.Data["Games"] = games
	c.TplName = "admin/game/list.tpl"
}

func (c *AdminController) GameAdd() {
	c.TplName = "admin/game/add.tpl"
}

func (c *AdminController) GameDoAdd() {
	title := c.GetString("title")
	category := c.GetString("category")
	status, _ := c.GetInt("status")

	f, _, err := c.GetFile("zipfile")
	if err == nil {
		defer f.Close()
		folder := "game_" + time.Now().Format("20060102150405")
		upload := "./static/uploads/"
		os.MkdirAll(upload, 0755)
		zipPath := upload + folder + ".zip"
		c.SaveToFile("zipfile", zipPath)
		utils.Unzip(zipPath, upload+folder)
		os.Remove(zipPath)

		o := orm.NewOrm()
		o.Insert(&models.Game{
			Title:    title,
			Category: category,
			Folder:   folder,
			Status:   status,
		})
	}
	c.Redirect("/admin/game", 302)
}

// 编辑游戏（修复 ORM Read 传参）
func (c *AdminController) GameEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	var g models.Game
	g.Id = id
	if err := o.Read(&g); err != nil {
		c.Redirect("/admin/game", 302)
		return
	}
	c.Data["Game"] = g
	c.TplName = "admin/game/edit.tpl"
}

func (c *AdminController) GameDoEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title := c.GetString("title")
	category := c.GetString("category")
	status, _ := c.GetInt("status")

	o := orm.NewOrm()
	g := models.Game{Id: id}
	if o.Read(&g) == nil {
		g.Title = title
		g.Category = category
		g.Status = status
		o.Update(&g)
	}
	c.Redirect("/admin/game", 302)
}

func (c *AdminController) GameDel() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Delete(&models.Game{Id: id})
	c.Redirect("/admin/game", 302)
}

// ==================== 工具管理 ====================
func (c *AdminController) ToolList() {
	o := orm.NewOrm()
	var tools []models.Tool
	o.QueryTable("tool").OrderBy("-id").All(&tools)
	c.Data["Tools"] = tools
	c.TplName = "admin/tool/list.tpl"
}

func (c *AdminController) ToolAdd() {
	c.TplName = "admin/tool/add.tpl"
}

func (c *AdminController) ToolDoAdd() {
	title := c.GetString("title")
	category := c.GetString("category")
	status, _ := c.GetInt("status")

	f, _, err := c.GetFile("zipfile")
	if err == nil {
		defer f.Close()
		folder := "tool_" + time.Now().Format("20060102150405")
		upload := "./static/uploads/"
		os.MkdirAll(upload, 0755)
		zipPath := upload + folder + ".zip"
		c.SaveToFile("zipfile", zipPath)
		utils.Unzip(zipPath, upload+folder)
		os.Remove(zipPath)

		o := orm.NewOrm()
		o.Insert(&models.Tool{
			Title:    title,
			Category: category,
			Folder:   folder,
			Status:   status,
		})
	}
	c.Redirect("/admin/tool", 302)
}

// 编辑工具（修复 ORM Read 传参）
func (c *AdminController) ToolEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	var t models.Tool
	t.Id = id
	if err := o.Read(&t); err != nil {
		c.Redirect("/admin/tool", 302)
		return
	}
	c.Data["Tool"] = t
	c.TplName = "admin/tool/edit.tpl"
}

func (c *AdminController) ToolDoEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title := c.GetString("title")
	category := c.GetString("category")
	status, _ := c.GetInt("status")

	o := orm.NewOrm()
	t := models.Tool{Id: id}
	if o.Read(&t) == nil {
		t.Title = title
		t.Category = category
		t.Status = status
		o.Update(&t)
	}
	c.Redirect("/admin/tool", 302)
}

func (c *AdminController) ToolDel() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Delete(&models.Tool{Id: id})
	c.Redirect("/admin/tool", 302)
}
