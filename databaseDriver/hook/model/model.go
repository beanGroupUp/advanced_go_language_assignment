package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User2 struct {
	ID         uint      `gorm:"primary_key"`
	Username   string    `gorm:"size:50;not null;uniqueIndex"`
	Email      string    `gorm:"size:100;not null;uniqueIndex"`
	Password   string    `gorm:"size:255;not null"`
	PostsCount int       `gorm:"default:0"` //新增：用户文章数量统计字段
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type Post2 struct {
	ID            uint      `gorm:"primary_key"`
	Title         string    `gorm:"size:200;not null;uniqueIndex"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null;index"`
	CommentsCount int       `gorm:"default:0"`
	CommentStatus string    `gorm:"size:20;default:'无评论'"` //新增：文章评论状态字段
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

type Comment2 struct {
	ID        uint      `gorm:"primary_key"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null;index"`
	PostID    uint      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

/**
钩子函数的注意事项：
	1、不能写在main文件里面
	2、避免循环调用：在钩子函数中不要执行会触发相同钩子的操作，例如在 BeforeSave 中调用 Save 方法。
	3、钩子函数与事务：钩子函数内执行的操作与主操作在同一个事务中，所以钩子函数中的错误会导致整个操作回滚。
	4、性能影响：钩子函数会增加数据库操作的额外开销，因为可能会执行额外的SQL语句。确保钩子中的操作是必要的，并尽量优化。
*/

// Post模型的钩子函数
// beforeCreate再创建文章前恒信用户的文章数量统计
/**
触发时机：在创建（插入）Post 记录到数据库之前。
执行逻辑：更新相应用户的文章计数（加1）。
注意：这个钩子只在创建操作时触发，更新操作不会触发。
*/
func (p *Post2) BeforeSave(tx *gorm.DB) (err error) {
	//更新用户的文章数量
	result := tx.Model(&User2{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("posts_count + 1"))

	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("用户 %v 的文章数量已更新\n", p.UserID)
	return nil
}

// beforeDelete再删除文章前更新用户的文章数量统计
/**
触发时机：在删除数据库中的 Post 记录之前。
执行逻辑：更新相应用户的文章计数（减1）。
*/
func (p *Post2) BeforeDelete(tx *gorm.DB) (err error) {
	//先查询文章信息，确保我们有正确的UserId
	if err := tx.First(p).Error; err != nil {
		return err
	}
	//更新用户的文章数量
	result := tx.Model(&User2{}).Where("id = ?", p.UserID).
		Update("posts_count", gorm.Expr("posts_count - 1"))
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("用户 %v 的文章数量以减少\n", p.UserID)
	return nil
}

// Comment 模型的钩子函数
// AfterCreate 再创建评论后更新文章的评论数量和状态
/**
触发时机：在创建（插入）Comment 记录到数据库之后。
执行逻辑：更新对应文章的评论计数（加1）和评论状态（设置为"有评论"）。
*/
func (c *Comment2) AfterSave(tx *gorm.DB) (err error) {
	//更新文章的评论数量
	result := tx.Model(&Post2{}).Where("id = ?", c.PostID).
		Update("comments_count", gorm.Expr("comments_count + 1"))
	if result.Error != nil {
		return result.Error
	}

	//更新文章的评论状态
	result = tx.Model(&User2{}).Where("id = ?", c.PostID).
		Update("comments_count", "有评论")
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("文章 %v 的评论数量已增加，状态更新为 有评论\n", c.PostID)
	return nil
}

// AfterDelete在删除评论后检查文章的评论数据并更新
/**
触发时机：在删除数据库中的 Comment 记录之后。
执行逻辑：检查对应文章的评论数量，如果为0则更新评论状态为"无评论"，否则更新评论数量。
*/
func (c *Comment2) AfterDelete(tx *gorm.DB) (err error) {
	//查询当前文章的评论数量
	var count int64
	if err := tx.Model(&Comment2{}).Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	//更新文章的评论数量
	result := tx.Model(&Post2{}).Where("id = ?", c.PostID).
		Update("comments_count", count)
	if result.Error != nil {
		return result.Error
	}

	//如果评论数量为0 更新评论状态
	if count == 0 {
		result := tx.Model(&Post2{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论")

		if result.Error != nil {
			return result.Error
		}
		fmt.Printf("文章 %v 的评论数量为0，状态更新为 无评论\n", c.PostID)
	} else {
		fmt.Printf("文章 %v 的评论数量已更新为 %d\n", c.PostID, count)
	}
	return nil
}

/**
钩子函数的触发机制
自动调用：当你使用 GORM 的 Create、Save、Delete 等方法时，GORM 会自动检查模型是否实现了对应的钩子方法，并在适当时机调用它们。

事务中执行：钩子函数在执行时，会接收一个 *gorm.DB 实例（即事务对象）。你在这个事务对象上执行的所有操作都会在同一个事务中。这意味着如果钩子函数返回错误，整个操作（包括主操作和钩子中的操作）都会回滚。

错误处理：如果钩子函数返回一个错误，GORM 会停止后续操作并回滚事务。因此，在钩子中遇到错误时，应该返回错误，以确保数据一致性。

避免无限递归：注意在钩子函数中不要调用会再次触发同一个钩子的方法，否则会导致递归调用。例如，在 BeforeSave 中不要调用 Save 方法。

示例流程
创建一篇文章（Post）的流程：
调用 db.Create(&post)

GORM 开始一个事务

调用 post.BeforeCreate(tx)（如果存在）

执行 INSERT 操作，将文章插入数据库

调用 post.AfterCreate(tx)（如果存在）

提交事务（如果所有步骤都成功）

创建一条评论（Comment）的流程：
调用 db.Create(&comment)

GORM 开始一个事务

调用 comment.BeforeCreate(tx)（如果存在，但你的代码中没有定义）

执行 INSERT 操作，将评论插入数据库

调用 comment.AfterCreate(tx)

在 AfterCreate 中更新文章的评论计数和状态

提交事务

删除一条评论（Comment）的流程：
调用 db.Delete(&comment)

GORM 开始一个事务

调用 comment.BeforeDelete(tx)（如果存在，但你的代码中没有定义）

执行 DELETE 操作，从数据库中删除评论

调用 comment.AfterDelete(tx)

在 AfterDelete 中检查并更新文章的评论计数和状态

提交事务

注意事项
钩子函数与事务：钩子函数内执行的操作与主操作在同一个事务中，所以钩子函数中的错误会导致整个操作回滚。

性能影响：钩子函数会增加数据库操作的额外开销，因为可能会执行额外的SQL语句。确保钩子中的操作是必要的，并尽量优化。

避免循环调用：在钩子函数中不要执行会触发相同钩子的操作，例如在 BeforeSave 中调用 Save 方法。

通过合理使用钩子函数，你可以将一些通用的业务逻辑（如更新计数、状态等）与模型的生命周期绑定，保持代码的整洁和可维护性。
*/
