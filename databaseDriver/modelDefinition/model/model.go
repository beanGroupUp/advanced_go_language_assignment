package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户模型
//type User struct {
//	gorm.Model
//	Username string ` gorm:"size:50;not null;uniqueIndex"`
//	Password string ` gorm:"size:100;not null;uniqueIndex"`
//	Email    string ` gorm:"size:255;not null"`
//	//一个用户拥有多个文章
//	Posts []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"posts,omitempty"` //一对多 的关系
//}

/**
除了 CASCADE，数据库和外键约束还支持其他几种删除行为，你可以根据业务逻辑的需要选择：

约束行为	效果描述	使用场景举例
SET NULL	当父表记录被删除时，子表中对应外键字段的值会被设置为 NULL。要求该外键字段允许为 NULL。	删除文章时，保留评论记录但将其标记为“无主文章”。
SET DEFAULT	当父表记录被删除时，子表中对应外键字段的值会被设置为该字段的默认值。需要为外键字段定义默认值。	用户被删除后，将其发布的文章分配给一个“默认用户”或“匿名用户”。
RESTRICT / NO ACTION	阻止删除父表中还有子记录引用的记录。如果尝试删除，数据库会抛出错误并回滚操作。	确保不会有文章在未被明确处理的情况下被意外删除，要求先手动删除所有评论。
DO_NOTHING	对删除操作不做任何处理。这可能导致数据不一致（产生孤立记录），因此需谨慎使用。	特殊场景，需要自行处理关联数据。
*/
//
//type Post struct {
//	gorm.Model
//	Title    string    ` gorm:"size:200;not null"`
//	Content  string    `gorm:"type:text;not null"`
//	UserID   uint      ` gorm:"size:100;not null"`
//	User     User      `gorm:"foreignKey:UserID"` // 添加用户关联
//	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;"`
//}
//
//type Comment struct {
//	gorm.Model
//	Content string ` gorm:"type:text;not null"`
//	PostID  uint   ` gorm:"size:100;not null"`
//	UserID  uint   ` gorm:"size:100;not null"`
//	User    User   `gorm:"foreignKey:UserID"` // 添加用户关联
//	Post    Post   `gorm:"foreignKey:PostID"` // 添加文章关联
//}
//
//// 查询结果结构体
//type PostWithComments struct {
//	Post     Post
//	Comments []CommentWithUser
//}
//
//type CommentWithUser struct {
//	Comment Comment
//	User    User
//}

// 方法二，不嵌套
type User struct {
	gorm.Model
	Username string `gorm:"size:50;not null;unique"`
	Email    string `gorm:"size:100;not null;uniqueIndex"`
	Password string `gorm:"size:255;not null;uniqueIndex"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"size:200;not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null;index"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"not null;index"`
	UserID  uint   `gorm:"not null;index"`
}

// 查询结果结构体
type UserPostsWithComments struct {
	gorm.Model
	UserID   uint
	Username string
	Email    string
	Posts    []PostWithComments
}

type PostWithComments struct {
	PostID    uint
	Title     string
	Content   string
	Comments  []CommentWithUser
	CreatedAt time.Time
}

type CommentWithUser struct {
	CommentID uint
	Content   string
	CreatedAt time.Time
	Email     string
	Username  string
}

type MostCommentedPost struct {
	PostID       uint
	Title        string
	Content      string
	Username     string
	CommentCount int64
}
