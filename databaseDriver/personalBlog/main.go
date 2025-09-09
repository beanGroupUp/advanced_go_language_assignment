package main

import (
	"awesomeProject3/databaseDriver/personalBlog/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局变量
var (
	db     *gorm.DB
	jwtKey = []byte("your_secret_key") //JWT秘钥，生产环境应使用环境变量
)

func main() {
	//初始化数据库连接
	initDB()

	//设置Gin路由
	r := setupRouter()

	//启动服务
	r.Run(":8081")
}

// initDB初始化数据库连接
func initDB() {
	var err error
	//创建日志器，显示详细日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_local?charset=utf8mb4&parseTime=True&loc=Local"
	//链接SQLLite数据库（也可替换为MYSQL）
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	//自动迁移数据库结构
	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	log.Println("Database initialized")
}

// setupRouter设置路由
func setupRouter() *gin.Engine {
	r := gin.Default()
	//公共路由-无需认证
	public := r.Group("/api")
	{
		public.POST("/register", register)
		public.POST("/login", login)
		public.GET("/posts", getPosts)
		public.GET("/posts/:id", getPost)
		public.GET("posts/:id/comments", getComments)
	}

	//受保护的路由- 需要JWT认证
	protected := r.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.POST("/posts", createPost)
		protected.PUT("/posts/:id", updatePost)
		protected.DELETE("/posts/:id", deletePost)
		protected.POST("/posts/:id/comments", createComment)
	}
	return r
}

// createComment 创建评论
func createComment(c *gin.Context) {
	postID := c.Param("id")
	var comment model.Comment

	//绑定JSON数据
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	//从上下文中获取用户ID
	userID, exits := c.Get("user_id")
	if !exits {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证的用户"})
		return
	}

	//检查文章是存在
	var post model.Post
	if err := db.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	//设置评论的用户和文章ID
	comment.UserID = userID.(uint)
	comment.PostID = post.ID

	//创建评论
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论创建失败"})
		return
	}

	//返回创建的评论
	db.Preload("User").First(&comment, comment.ID)
	c.JSON(http.StatusCreated, gin.H{"comment": comment})

}

// deletePost删除文章
func deletePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	//查找文章
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	//从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "未认证 的用户"})
		return
	}

	//检查是否为文章作者
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权操作此文章"})
	}

	//删除文章
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "文章删除成功"})

}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	//查找文章
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
	}

	//从上文中获取用户id
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证的用户"})
		return
	}

	//检查是否为文章作者
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权操作此文章"})
		return
	}

	//绑定更新数据
	var updateData model.Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	//更新文章
	if err := db.Model(&post).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章更新失败"})
	}

	c.JSON(http.StatusOK, post)
}

// createPost创建文章
func createPost(c *gin.Context) {
	var post model.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	//从上下文中获取用户ID
	//将变量 userID 的值转换为 uint 类型，并赋值给 post.UserID
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"error": "未认证的用户"})
		return
	}

	//设置文章作者
	post.UserID = userID.(uint)

	//创建文章
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
		return
	}

	//返回创建的文章
	db.Preload("User").First(&post, post.UserID)
	c.JSON(http.StatusCreated, post)
}

// authMiddleware
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头获取token
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
			// 调用 c.Abort() 后，后续中间件和处理函数不会执行(该函数接下来的内容还会继续执行)
			c.Abort()
			return
		}

		//解析和验证token
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的认证令牌"})
			// 调用 c.Abort() 后，后续中间件和处理函数不会执行(该函数接下来的内容还会继续执行)
			c.Abort()
			return
		}

		//将用户ID存入上下文
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// getComments获取文章评论
func getComments(c *gin.Context) {
	postID := c.Param("id")

	var comments []model.Comment
	if err := db.Preload("User").Where("id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func getPosts(c *gin.Context) {
	var posts []model.Post

	//获取所有文章可以关联的用户信息
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// getPost获取单个文章
func getPost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post

	//获取文章以及关联的用户和评论
	if err := db.Preload("User").First("Comments.User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// login用户登录
func login(c *gin.Context) {
	var inputUser model.User
	//绑定JSON数据
	if err := c.ShouldBind(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求数据"})
		return
	}

	//查找用户
	var user model.User
	if err := db.Where("username = ?", inputUser.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	//生成JWT令牌
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := model.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "令牌生成失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
		"token":   tokenString,
		"user_id": user.ID,
	})
}

// register 用户注册
func register(c *gin.Context) {
	var user model.User
	//绑定JSON数据到用户结构
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	//查找用户
	var existingUser model.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	//检查邮箱是否已经存在
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已存在"})
		return
	}

	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}
	//赋值
	user.Password = string(hashedPassword)
	//创建用户
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})

}
