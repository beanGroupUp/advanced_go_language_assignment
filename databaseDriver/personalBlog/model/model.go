package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

//数据库模型定义

// User用户模型
type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null" json:"username"` //用户名，唯一且不为空
	Password string    `gorm:"not null" json:"password"`        //密码不为空
	Email    string    `gorm:"unique;not null" json:"email"`    //邮箱唯一且不为空
	Posts    []Post    `gorm:"foreignkey:UserID"`               //用户关联的文章
	Comments []Comment `gorm:"foreignkey:UserID"`               //用户关联的评论
}

// Post文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title"`   //文章标题，不为空
	Content  string    `gorm:"not null" json:"content"` //文章内容，不为空
	UserID   uint      `json:"user_id"`                 //用户ID，外检
	User     User      `gorm:"foreignkey:UserID"`       //关联的用户
	Comments []Comment `gorm:"foreignkey:UserID"`       //文章关联的评论
}

// Comment评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"` //评论内容，不为空
	UserID  uint   `json:"user_id"`                 //用户ID，外检
	User    User   `gorm:"foreignkey:UserID"`       //关联的用户
	PostID  uint   `json:"post_id"`                 //文章ID，外检
	Post    Post   `gorm:"foreignkey:PostID"`       //关联的文章
}

// JWT声明结构
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
