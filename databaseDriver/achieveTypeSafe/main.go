package main

import (
	"awesomeProject3/databaseDriver/achieveTypeSafe/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
//并将结果映射到 Book 结构体切片中，确保类型安全。
//进阶gorm
//
//func init() {
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold:             time.Second,
//			LogLevel:                  logger.Info,
//			Colorful:                  true,
//			ParameterizedQueries:      true,
//			IgnoreRecordNotFoundError: true,
//		},
//	)
//	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//	})
//
//	if err != nil {
//		panic(err)
//	}
//
//	//建表
//	db.AutoMigrate(model.Books{})
//	//创建初始化数据
//	books := []model.Books{
//		{Title: "西游记", Author: "吴承恩", Price: 128.0},
//		{Title: "放风筝的人", Author: "史密斯", Price: 50.0},
//		{Title: "潜水艇", Author: "达到", Price: 10.0},
//		{Title: "收到撒打算", Author: "vv", Price: 200.0},
//		{Title: "奋斗奋斗", Author: "广告费", Price: 20.0},
//	}
//	db.Create(&books)
//}

func main() {
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败：%v,服务断开；", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("数据库连接成功；")
	var books []model.Books
	cont := "select id , title , author , price from books where price > ?"
	err = db.Select(&books, cont, 50)
	if err != nil {
		fmt.Errorf("数据查询异常；")
	}
	// 输出结果
	fmt.Printf("找到 %d 本价格大于50元的书籍:\n", len(books))
	for _, v := range books {
		fmt.Printf("ID：%v,书名：%v,作者：%v,价格：%v\n", v.ID, v.Title, v.Author, v.Price)

	}

}
