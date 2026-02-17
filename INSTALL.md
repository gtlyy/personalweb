# personalweb 个人主页部署指南

## 环境要求

- Go 1.21+
- SQLite3

## 安装步骤

### 1. 克隆项目

```bash
git clone <your-repo-url>
cd personalweb
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置数据库

首次运行时会自动创建数据库 `data/blog.db` 并创建默认管理员账号。

如需手动初始化，可修改 `conf/app.conf` 中的数据库配置：

```ini
dbtype = sqlite3
dbpath = ./data/blog.db
```

### 4. 运行项目

```bash
go run main.go
```

服务启动后访问 http://localhost:9090

### 5. 管理员后台

- 地址：http://localhost:9090/admin/login
- 默认账号：`admin`
- 默认密码：`admin123`

登录后可管理文章、游戏、工具等内容。

## 项目结构

```
personalweb/
├── conf/          # 配置文件
├── controllers/   # 控制器
├── models/       # 数据模型
├── routers/      # 路由配置
├── static/       # 静态资源 (CSS, JS, 图片)
├── views/        # 模板文件
├── utils/        # 工具函数
├── data/         # 数据库目录 (需手动创建)
└── main.go       # 入口文件
```

## 注意事项

1. 首次运行前确保 `data/` 目录存在（程序会自动创建）
2. 数据库文件会在首次运行时自动创建
3. 静态资源已下载到本地，无需网络连接
4. Markdown 编辑器使用 EasyMDE，列表与段落之间需空一行
