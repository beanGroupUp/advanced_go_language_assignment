# 个人博客系统后端

基于 Go + Gin + GORM 开发的个人博客系统后端，支持用户认证、文章管理和评论功能。

## 功能特性

- 用户注册和登录（JWT认证）
- 文章的创建、读取、更新和删除（CRUD）
- 文章的评论功能
- 完整的错误处理和日志记录

## 运行环境

- Go 1.16+
- SQLite (也可配置为MySQL、PostgreSQL等)

## 安装和运行

1. 克隆项目到本地
```bash
git clone <repository-url>
cd go-blog-backend



安装依赖

bash
go mod download
运行项目

bash
go run main.go
服务器将在 http://localhost:8080 启动

API接口文档
认证相关
POST /api/register - 用户注册

POST /api/login - 用户登录

文章相关
GET /api/posts - 获取所有文章

GET /api/posts/:id - 获取指定文章

POST /api/posts - 创建文章（需要认证）

PUT /api/posts/:id - 更新文章（需要认证）

DELETE /api/posts/:id - 删除文章（需要认证）

评论相关
GET /api/posts/:id/comments - 获取文章评论

POST /api/posts/:id/comments - 创建评论（需要认证）