package main

import (
	"awesomeProject3/databaseDriver/modelDefinition/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//方法一，嵌套
//var db *gorm.DB
//
//func main() {
//	var err error
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold:             time.Second,
//			LogLevel:                  logger.Info,
//			IgnoreRecordNotFoundError: true,
//			ParameterizedQueries:      true, //设置为 false，在 SQL 日志中显示实际参数
//			Colorful:                  true,
//		},
//	)
//	dsn := "root:123456@tcp(127.0.0.1:3306)/db_local?charset=utf8mb4&parseTime=True&loc=Local"
//
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//	})
//	if err != nil {
//		log.Fatal("数据库连接异常：%v", err)
//	}
//
//	//自动迁移模式，创建或更新表结构钢
//	//err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
//	//if err != nil {
//	//	log.Fatal("创建表失败：%v", err)
//	//}
//	//fmt.Println("表数据创建成功；")
//
//	//createTestData(db)
//
//	//查询某个用户发布的所有文章及其对应的评论信息
//	userID := uint(1) //假设查询用户ID为1的文章和评论
//
//	comments, err := getUserPostWithComments(db, userID)
//	if err != nil {
//		log.Printf("查询异常：%v", err)
//	} else {
//		fmt.Println("\n用户%d 的所有文章以及评论 ：\n", userID)
//		for _, comment := range comments {
//			fmt.Printf("\n文章标题：%v", comment.Post.Title)
//			fmt.Printf("\n文章内容：%v\n", comment.Post.Content)
//			fmt.Printf("\n评论数量：%v", len(comment.Post.Content))
//
//			for _, commentWithUser := range comment.Comments {
//				fmt.Printf(" - 用户 %v 评论：%v\n",
//					commentWithUser.User.Username,
//					commentWithUser.Comment)
//			}
//		}
//	}
//
//	//2.评论数量最多的文章信息
//	post, i, err := getMostCommentedPost(db)
//	if err != nil {
//		log.Printf("查询异常：%v", err)
//	} else {
//		fmt.Printf("\n评论数量最多的文章\n")
//		fmt.Printf("文章标题：%v\n", post.Title)
//		fmt.Printf("文章作者%v\n", post.User.Username)
//		fmt.Printf("评论数量%v\n", i)
//		fmt.Printf("文章内容：%v\n", post.Content)
//	}
//
//}
//
//func getUserPostWithComments(db *gorm.DB, userID uint) ([]model.PostWithComments, error) {
//	var posts []model.Post
//	//预加载用户信息和评论信息（包括评论的用户信息）
//	err := db.Preload("Comments.User"). // 预加载评论的作者信息
//						Where("user_id = ?", userID).
//						Find(&posts).Error
//	if err != nil {
//		return nil, err
//	}
//
//	var result []model.PostWithComments
//	for _, post := range posts {
//		var commentsWithUsers []model.CommentWithUser
//		for _, comment := range post.Comments {
//			commentsWithUsers = append(commentsWithUsers, model.CommentWithUser{
//				Comment: comment,
//				User:    comment.User,
//			})
//		}
//
//		result = append(result, model.PostWithComments{
//			Post:     post,
//			Comments: commentsWithUsers,
//		})
//	}
//	return result, nil
//}
//
//// getMostCommentedPost评论数量最多的文章
//func getMostCommentedPost(db *gorm.DB) (model.Post, int64, error) {
//	var commentCount int64
//	var post model.Post
//
//	//子查询：计算每篇文章的评论数量
//	subQuery := db.Model(&model.Comment{}).
//		Select("post_id,count(*) as comment_count").
//		Group("post_id")
//
//	//主查询：连接文章表和评论统计子查询，按评论数量降序排列，取第一条
//	err := db.Preload("User").
//		Select("posts.*,comment_stats.comment_count").
//		Joins("INNER JOIN (?) AS comment_stats ON posts.id = comment_stats.post_id", subQuery).
//		Order("comment_stats.comment_count desc").
//		Find(&post).Error
//
//	if err != nil {
//		return model.Post{}, 0, err
//	}
//
//	//获取评论数量
//	err = db.Model(&model.Comment{}).
//		Where("post_id = ?", post.ID).
//		Count(&commentCount).Error
//
//	if err != nil {
//		return model.Post{}, 0, err
//	}
//	return post, commentCount, nil
//}
//
//// 可选：添加一个更搞笑的查询评论最多文章的啊方法
//func getMostCommentedPostOptimized(db *gorm.DB) (model.Post, int64, error) {
//	var result struct {
//		model.Post
//		CommentCount int64
//	}
//
//	err := db.Model(&model.Post{}).
//		Select("posts.*,Count(comments.id) as comment_count").
//		Joins("Left join comments on posts.id = comments.post_id").
//		Group("posts.id").
//		Order("comment_count desc").
//		Preload("User").
//		First(&result).Error
//
//	if err != nil {
//		return model.Post{}, 0, err
//	}
//	return result.Post, result.CommentCount, nil
//}

// createTestData 创建测试数据
//func createTestData(db *gorm.DB) {
//	// 初始化随机数生成器
//	rand.Seed(time.Now().UnixNano())
//
//	// 创建用户
//	users := createUsers(db, 10)
//
//	// 创建文章
//	posts := createPosts(db, users, 25)
//
//	// 创建评论
//	createComments(db, users, posts, 100)
//
//	fmt.Println("Test data created successfully!")
//}
//
//// createUsers 创建用户测试数据
//func createUsers(db *gorm.DB, count int) []model.User {
//	users := []model.User{
//		{Username: "john_doe", Email: "john@example.com", Password: "hashed_password_123"},
//		{Username: "jane_smith", Email: "jane@example.com", Password: "hashed_password_456"},
//		{Username: "alice_wonder", Email: "alice@example.com", Password: "hashed_password_789"},
//		{Username: "bob_builder", Email: "bob@example.com", Password: "hashed_password_abc"},
//		{Username: "charlie_brown", Email: "charlie@example.com", Password: "hashed_password_def"},
//		{Username: "diana_prince", Email: "diana@example.com", Password: "hashed_password_ghi"},
//		{Username: "edward_elric", Email: "edward@example.com", Password: "hashed_password_jkl"},
//		{Username: "fiona_gallagher", Email: "fiona@example.com", Password: "hashed_password_mno"},
//		{Username: "george_orwell", Email: "george@example.com", Password: "hashed_password_pqr"},
//		{Username: "hannah_montana", Email: "hannah@example.com", Password: "hashed_password_stu"},
//	}
//
//	if count < len(users) {
//		users = users[:count]
//	}
//
//	for i := range users {
//		db.Create(&users[i])
//	}
//
//	fmt.Printf("Created %d users\n", len(users))
//	return users
//}
//
//// createPosts 创建文章测试数据
//func createPosts(db *gorm.DB, users []model.User, count int) []model.Post {
//	titles := []string{
//		"Getting Started with Go Programming",
//		"Understanding Database Relationships",
//		"Web Development Best Practices",
//		"The Future of Artificial Intelligence",
//		"Building Scalable Web Applications",
//		"Introduction to Machine Learning",
//		"Effective Team Collaboration",
//		"Mastering RESTful APIs",
//		"Cybersecurity Essentials",
//		"Cloud Computing Explained",
//		"Data Structures and Algorithms",
//		"DevOps Culture and Practices",
//		"Mobile App Development Trends",
//		"Blockchain Technology Overview",
//		"User Experience Design Principles",
//		"Agile Methodology Deep Dive",
//		"Microservices Architecture",
//		"Containerization with Docker",
//		"Serverless Computing Benefits",
//		"Progressive Web Apps Guide",
//		"API Security Best Practices",
//		"GraphQL vs REST Comparison",
//		"CI/CD Pipeline Implementation",
//		"Testing Strategies for Web Apps",
//		"Performance Optimization Techniques",
//	}
//
//	contents := []string{
//		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam euismod, nisl eget ultricies ultricies, nunc nisl aliquam nunc, quis aliquam nisl nunc eu nisl.",
//		"Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Nullam euismod, nisl eget ultricies ultricies, nunc nisl aliquam nunc.",
//		"Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Nullam euismod, nisl eget ultricies ultricies, nunc nisl aliquam nunc.",
//		"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis.",
//		"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi.",
//		"Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime.",
//		"Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae.",
//		"Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat.",
//		"On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment.",
//		"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system.",
//	}
//
//	posts := make([]model.Post, 0, count)
//
//	for i := 0; i < count; i++ {
//		user := users[rand.Intn(len(users))]
//		title := titles[rand.Intn(len(titles))]
//		content := contents[rand.Intn(len(contents))]
//
//		post := model.Post{
//			Title:   title,
//			Content: content,
//			UserID:  user.ID,
//		}
//
//		db.Create(&post)
//		posts = append(posts, post)
//	}
//
//	fmt.Printf("Created %d posts\n", count)
//	return posts
//}
//
//// createComments 创建评论测试数据
//func createComments(db *gorm.DB, users []model.User, posts []model.Post, count int) {
//	comments := []string{
//		"Great post! Very informative.",
//		"I completely agree with your points.",
//		"Thanks for sharing this valuable information.",
//		"This has changed my perspective on the topic.",
//		"Could you elaborate more on the second point?",
//		"I've had a similar experience with this.",
//		"Well written and easy to understand.",
//		"Looking forward to reading more from you.",
//		"I shared this with my team, very useful.",
//		"Your examples really helped clarify things.",
//		"Interesting take on this subject.",
//		"I never thought about it this way before.",
//		"This is exactly what I was looking for.",
//		"Clear and concise explanation.",
//		"Would love to see a follow-up on this topic.",
//		"Helpful for my current project, thanks!",
//		"I bookmarked this for future reference.",
//		"Your insights are always valuable.",
//		"This answered all my questions.",
//		"Perfect timing, I needed this information.",
//	}
//
//	for i := 0; i < count; i++ {
//		user := users[rand.Intn(len(users))]
//		post := posts[rand.Intn(len(posts))]
//		comment := comments[rand.Intn(len(comments))]
//
//		commentObj := model.Comment{
//			Content: comment,
//			UserID:  user.ID,
//			PostID:  post.ID,
//		}
//
//		db.Create(&commentObj)
//	}
//
//	fmt.Printf("Created %d comments\n", count)
//}

//方法二，非嵌套，关联查

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_local?charset=utf8&parseTime=True&loc=Local"

	//初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database:%v", err)
	}

	//1.查询某个用户发布的所有文章及其对应的评论信息
	userID := uint(1)
	userPostsWithComments, err := getUserPostsWithComments(db, userID)
	if err != nil {
		log.Fatalf("failed to get userPosts:%v", err)
	} else {
		fmt.Println("\n用户 %s (S)的所有文章及评论：\n",
			userPostsWithComments.Username, userPostsWithComments.Email)

		for _, post := range userPostsWithComments.Posts {
			fmt.Printf("\n文章标题：%s\n", post.Title)
			fmt.Printf("文章内容：%s\n", post.Content)
			fmt.Printf("发布时间：%s\n", post.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Printf("评论数量“%v\n", len(post.Comments))

			for _, comment := range post.Comments {
				fmt.Printf(" - 用户 %v 评论： %v(时间：%v)\n", comment.Username, comment.CreatedAt, comment.CreatedAt.Format("2006-01-02 15:04:05"))
			}
		}
	}

	//2.查询评论数量最多的文章信息
	mostCommentedPost, err := getMostCommentedPost(db)

	if err != nil {
		log.Fatalf("failed to get most commented post:%v", err)
	} else {
		fmt.Printf("\n评论数量最多的文章\n")
		fmt.Printf("文章标题%s\n", mostCommentedPost.Title)
		fmt.Printf("文章作者%s\n", mostCommentedPost.Username)
		fmt.Printf("评论数量：%d\n", mostCommentedPost.CommentCount)
		fmt.Printf("文章内容%v\n", mostCommentedPost.Content)
	}
}

func getUserPostsWithComments(db *gorm.DB, userID uint) (model.UserPostsWithComments, error) {
	var result model.UserPostsWithComments

	//获取用户信息
	var user model.User
	if err := db.First(&user, userID).Error; err != nil {
		return result, err
	}

	result.UserID = user.ID
	result.Username = user.Username
	result.Email = user.Email

	//获取用户的所有文章
	var posts []model.Post
	if err := db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return result, err
	}

	//获取每篇文章的评论
	for _, post := range posts {
		postWithComments := model.PostWithComments{
			PostID:    post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		}
		//获取这篇文章的所有评论以及评论者信息
		var comments []struct {
			model.Comment
			Username string
			Email    string
		}

		if err := db.Table("comments").
			Select("comments.*,users.username,users.email").
			Joins("left join users on comments.user_id = users.id").
			Where("comments.post_id = ?", post.ID).
			Scan(&comments).Error; err != nil {
			return result, err
		}

		//转换评论数据
		for _, c := range comments {
			postWithComments.Comments = append(postWithComments.Comments, model.CommentWithUser{
				CommentID: c.ID,
				Content:   c.Content,
				CreatedAt: c.CreatedAt,
				Username:  c.Username,
				Email:     c.Email,
			})
		}
		result.Posts = append(result.Posts, postWithComments)
	}
	return result, nil
}

func getMostCommentedPost(db *gorm.DB) (model.MostCommentedPost, error) {
	var result model.MostCommentedPost

	//使用子查询获取评论数量最多的文章ID
	var postID uint
	err := db.Model(&model.Comment{}).
		Select("post_id").
		Group("post_id").
		Order("Count(*) desc").
		Limit(1).
		//pluck查询单个字段值
		Pluck("post_id", &postID).Error

	if err != nil {
		return result, err
	}

	//获取文章详细信息
	var post model.Post
	if err := db.First(&post, postID).Error; err != nil {
		return result, err
	}

	//获取文章作者信息
	var user model.User
	if err := db.First(&user, post.UserID).Error; err != nil {
		return result, err
	}

	//获取评论数量
	var commentCount int64
	if err := db.Model(&model.Comment{}).
		Where("post_id = ?", postID).
		Count(&commentCount).Error; err != nil {
		return result, err
	}

	result.PostID = post.ID
	result.Title = post.Title
	result.Content = post.Content
	result.Username = user.Username
	result.CommentCount = commentCount

	return result, nil

}

// 可选：更搞笑的查询评论最多文章的方法
func getMostCommentedComment(db *gorm.DB) (model.MostCommentedPost, error) {
	var result model.MostCommentedPost
	err := db.Table("posts").
		Select("posts.id as post_id, posts.title, posts.content, users.username,Count(comments.id) as comment_count)").
		Joins("left join comments on posts.id = comments.post_id").
		Joins("left join users on users.id = comments.user_id").
		Group("post_id,posts.title,posts.content,users.username").
		Order("comments.id desc").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		return model.MostCommentedPost{}, err
	}

	return result, nil
}
