# PersonalWeb

一个简单的个人网站系统，支持博客、游戏、工具管理。

## 功能特性

- 📝 文章管理 - Markdown 编辑器
- 🎮 游戏上传与在线游玩
- 🔧 工具上传与在线使用
- 🔐 管理员后台
- 📱 响应式设计

## 技术栈

- Go + Beego v2
- SQLite
- EasyMDE (Markdown 编辑器)
- Tailwind CSS

## 安装

```bash
git clone https://github.com/gtlyy/personalweb.git
cd personalweb
go mod download
go run main.go
```

首次启动会自动创建数据库和默认管理员账号。

## 配置

编辑 `conf/app.conf`：

```ini
appname = personalweb
httpport = 9090
runmode = dev

# 上传功能开关（默认关闭）
enable_upload = true
```

## 默认管理员

- 用户名: `admin`
- 密码: `admin123`

## 许可证

MIT License - 查看 [LICENSE](LICENSE) 文件
