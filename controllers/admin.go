package controllers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"personalweb/models"
	"personalweb/utils"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	MaxUploadSize     = 50 * 1024 * 1024 // 50MB
	AllowedExtensions = ".zip"
)

var (
	ErrFileTooLarge    = errors.New("文件大小超出限制")
	ErrInvalidFileType = errors.New("只允许上传 zip 文件")
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
	admin, ok := c.GetSession("admin").(models.Admin)
	if !ok {
		c.Redirect("/admin/login", 302)
		c.StopRun()
		return
	}
	c.Data["Admin"] = admin
}

func (c *AdminController) validateUpload(filePath string) error {
	// 检查文件大小
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if fileInfo.Size() > MaxUploadSize {
		return ErrFileTooLarge
	}

	// 检查文件扩展名
	ext := strings.ToLower(filepath.Ext(fileInfo.Name()))
	if ext != AllowedExtensions {
		return ErrInvalidFileType
	}

	// 验证文件头 (zip 文件魔数: 50 4B)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	header := make([]byte, 2)
	n, err := f.Read(header)
	if err != nil || n < 2 {
		return errors.New("无法读取文件")
	}

	// ZIP: 50 4B (PK)
	if header[0] != 0x50 || header[1] != 0x4B {
		return ErrInvalidFileType
	}

	return nil
}

func (c *AdminController) handleUpload(folderType string) (string, error) {
	f, header, err := c.GetFile("zipfile")
	if err != nil {
		return "", err
	}
	defer f.Close()

	// 先保存到临时文件
	folder := folderType + "_" + time.Now().Format("20060102150405")
	tempDir := "./static/uploads/temp/"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", err
	}
	tempPath := tempDir + header.Filename

	if err := c.SaveToFile("zipfile", tempPath); err != nil {
		return "", err
	}
	defer os.Remove(tempPath)

	// 验证文件
	if err := c.validateUpload(tempPath); err != nil {
		return "", err
	}

	// 解压
	upload := "./static/uploads/"
	extractPath := upload + folder
	if err := utils.Unzip(tempPath, extractPath); err != nil {
		return "", fmt.Errorf("解压失败: %v", err)
	}

	return folder, nil
}

// 登录
func (c *AdminController) Login() {
	attempts := c.GetSession("login_attempts")
	if attempts == nil {
		c.SetSession("login_attempts", 0)
	}
	c.TplName = "admin/login.tpl"
}

func (c *AdminController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")

	attempts := c.GetSession("login_attempts")
	attemptCount := 0
	if attempts != nil {
		attemptCount = attempts.(int)
	}

	if attemptCount >= 5 {
		lastAttempt := c.GetSession("last_login_attempt")
		if lastAttempt != nil {
			lastTime := lastAttempt.(time.Time)
			if time.Since(lastTime) < 5*time.Minute {
				c.Data["Msg"] = "登录尝试过多，请5分钟后再试"
				c.TplName = "admin/login.tpl"
				return
			}
		}
		c.SetSession("login_attempts", 0)
	}

	o := orm.NewOrm()
	var admin models.Admin
	err := o.QueryTable("admin").Filter("username", username).One(&admin)
	if err != nil {
		c.SetSession("login_attempts", attemptCount+1)
		c.SetSession("last_login_attempt", time.Now())
		c.Data["Msg"] = "账号不存在"
		c.TplName = "admin/login.tpl"
		return
	}
	if !admin.CheckPassword(password) {
		c.SetSession("login_attempts", attemptCount+1)
		c.SetSession("last_login_attempt", time.Now())
		c.Data["Msg"] = "密码错误"
		c.TplName = "admin/login.tpl"
		return
	}

	c.SetSession("login_attempts", 0)
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

// 编辑文章
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
	adminObj, ok := c.GetSession("admin").(models.Admin)
	if !ok {
		c.DestroySession()
		c.Redirect("/admin/login", 302)
		return
	}
	o := orm.NewOrm()
	var admin models.Admin
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

	folder, err := c.handleUpload("game")
	if err != nil {
		logs.Error("处理上传文件失败: %v", err)
		c.Redirect("/admin/game", 302)
		return
	}

	o := orm.NewOrm()
	o.Insert(&models.Game{
		Title:    title,
		Category: category,
		Folder:   folder,
		Status:   status,
	})
	c.Redirect("/admin/game", 302)
}

// 编辑游戏
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
	game := models.Game{Id: id}
	if o.Read(&game) == nil {
		uploadPath := "./static/uploads/" + game.Folder
		os.RemoveAll(uploadPath)
		o.Delete(&game)
	}
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

	folder, err := c.handleUpload("tool")
	if err != nil {
		logs.Error("处理上传文件失败: %v", err)
		c.Redirect("/admin/tool", 302)
		return
	}

	o := orm.NewOrm()
	o.Insert(&models.Tool{
		Title:    title,
		Category: category,
		Folder:   folder,
		Status:   status,
	})
	c.Redirect("/admin/tool", 302)
}

// 编辑工具
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
	tool := models.Tool{Id: id}
	if o.Read(&tool) == nil {
		uploadPath := "./static/uploads/" + tool.Folder
		os.RemoveAll(uploadPath)
		o.Delete(&tool)
	}
	c.Redirect("/admin/tool", 302)
}
