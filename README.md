# GoCode Web 项目

基于 Gin 框架的 Go Web 应用程序。

## 项目结构

```
.
├── controller/          # 控制器层
│   ├── admin/          # 后台管理控制器
│   │   ├── articleController.go   # 文章管理
│   │   ├── indexController.go     # 后台首页
│   │   └── userController.go      # 用户管理
│   ├── api/            # API 接口控制器
│   │   └── apiController.go       # API 接口
│   └── nannanwu/       # 前台控制器
│       ├── baseController.go      # 基础控制器
│       └── defaultController.go   # 默认控制器
├── router/             # 路由配置
│   ├── adminRouter.go  # 后台路由
│   ├── apiRouter.go    # API 路由
│   └── defaultRouter.go # 默认路由
├── static/             # 静态资源
│   ├── footer.css      # 底部样式
│   ├── navbar.css      # 导航栏样式
│   └── style.css       # 全局样式
├── templates/          # HTML 模板
│   ├── layout/         # 布局模板
│   │   ├── base.html   # 基础布局
│   │   ├── footer.html # 底部模板
│   │   └── navbar.html # 导航栏模板
│   └── pages/          # 页面模板
│       ├── index.html  # 首页
│       ├── test.html   # 测试页
│       └── user.html   # 用户页
├── main.go             # 程序入口
├── go.mod              # Go 模块配置
└── go.sum              # Go 依赖校验
```

## 技术栈

- **框架**: [Gin](https://github.com/gin-gonic/gin) v1.11.0
- **语言**: Go 1.25+
- **模板**: Go HTML Template

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 运行项目

```bash
go run main.go
```

或编译后运行：

```bash
go build -o app.exe .
./app.exe
```

### 3. 访问应用

- 前台首页: http://localhost:8088/
- 后台管理: http://localhost:8088/admin
- API 接口: http://localhost:8088/api

## 路由说明

### 前台路由 (`/`)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/` | 首页 |
| GET | `/article` | 文章详情 |
| GET | `/success` | 成功页面 |
| GET | `/error` | 错误页面 |

### 后台路由 (`/admin`)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/admin` | 后台首页 |
| GET | `/admin/login` | 登录页面 |
| GET | `/admin/plist` | 列表页面 |
| GET | `/admin/article` | 文章列表 |
| GET | `/admin/user` | 用户列表 |
| GET | `/admin/user/:id` | 用户详情 |

### API 路由 (`/api`)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api` | API 首页 |
| GET | `/api/user` | 用户信息 |
| GET | `/api/plist` | 用户列表 |

## 控制器说明

### BaseController

提供基础的响应方法：
- `Success()` - 返回成功响应
- `Error()` - 返回错误响应

### 继承关系

```
BaseController
    └── DefaultController (继承 BaseController)
```

## 模板系统

使用 Go 的 `html/template`，支持模板继承：

- **基础模板**: `layout/base.html`
- **页面模板**: 继承 base 模板，填充 `content` block

示例：
```html
{{ define "layout/base" }}
    {{ block "content" . }}{{ end }}
{{ end }}
```

## 静态资源

静态文件位于 `/static` 目录，通过 `/static/` 路径访问：

- CSS: `/static/style.css`
- 导航栏样式: `/static/navbar.css`
- 底部样式: `/static/footer.css`

## 配置说明

- **端口**: 8088 (在 `main.go` 中修改)
- **模板路径**: `templates/**/*`
- **静态资源**: `./static` 映射到 `/static`

## 开发规范

1. **控制器**: 按功能模块分组，放在 `controller/` 下
2. **路由**: 每个模块独立路由文件，统一在 `router/` 管理
3. **模板**: 布局模板放 `layout/`，页面模板放 `pages/`
4. **静态资源**: CSS 放 `static/`，按功能分文件

## 常见问题

### 编译错误: missing go.sum entry

```bash
go mod tidy
```

### 端口被占用

修改 `main.go` 中的端口号：
```go
r.Run(":8088")  // 修改为其他端口
```
