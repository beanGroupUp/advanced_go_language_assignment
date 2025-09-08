package main

import (
	"awesomeProject3/databaseDriver/hook/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//数据库连接配置
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_local?charset=utf8&parseTime=True&loc=Local"

	//初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql 数据库连接失败:%v", err)
	}

	//自动迁移模式，创建或更新表结构
	err = db.AutoMigrate(&model.User2{}, &model.Post2{}, &model.Comment2{})

	if err != nil {
		log.Fatal("表结构创建失败：%v", err)
	}

	fmt.Println("表结构创建成功~！")
	//测试创建用户
	user := model.User2{
		Username: "test_user",
		Email:    "test@example.com",
		Password: "123456",
	}
	db.Create(&user)

	//创建测试文章
	post := model.Post2{
		Title:   "测试文章",
		Content: "这是一篇测试文章的内容",
		UserID:  user.ID,
	}
	//将id回填入post里面了
	db.Create(&post)

	//创建测试评论
	comment := model.Comment2{
		Content: "这是一条测试评论",
		UserID:  user.ID,
		PostID:  post.ID,
	}
	db.Create(&comment)

	//查询用户信息，查看文章数量是否更新
	var updateUser model.User2
	db.First(&updateUser, user.ID)
	fmt.Printf("用户文章数量：%d\n", updateUser.PostsCount)

	//查询文章信息，查看评论数量和状态是否更新
	var updatePost model.Post2
	db.First(&updatePost, post.ID)
	fmt.Printf("文章评论数量%d，评论状态：%s\n", updatePost.CommentsCount, updatePost.CommentStatus)

	//删除评论，查看评论状态是否更新
	db.Delete(&comment)

	//再次查询文章信息
	db.First(&updateUser, post.ID)
	fmt.Printf("删除评论后 - 文章评论数量：%d,评论状态:%s\n", updatePost.CommentsCount, updatePost.CommentStatus)

}
