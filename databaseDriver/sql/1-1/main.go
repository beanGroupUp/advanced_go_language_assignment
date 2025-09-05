package main

import (
	"awesomeProject3/databaseDriver/sql/1-1/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false, //设置为 false，在 SQL 日志中显示实际参数
			Colorful:                  true,
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println("连接成功")
	//err = db.AutoMigrate(&model.Students{})
	//if err != nil {
	//	return
	//}
}

func main() {
	/*	db.Model(&model.Students{}).Create(&model.Students{
			Name:  "张三",
			Age:   20,
			Grade: "三年级",
		})
	*/

	/*	stu := []model.Students{
			{Name: "猪八戒", Age: 20, Grade: "四年级"},
			{Name: "孙悟空", Age: 10, Grade: "一年级"},
			{Name: "唐僧", Age: 18, Grade: "二年级"},
			{Name: "白龙马", Age: 15, Grade: "二年级"},
			{Name: "癞蛤蟆", Age: 16, Grade: "二年级"},
			{Name: "小猪妖", Age: 22, Grade: "五年级"},
		}
		db.Create(&stu)*/

	//stu := []model.Students{
	//	{Name: "小猪妖1", Age: 5, Grade: "一年级"},
	//	{Name: "小猪妖2", Age: 6, Grade: "一年级"},
	//	{Name: "小猪妖3", Age: 7, Grade: "一年级"},
	//	{Name: "小猪妖4", Age: 8, Grade: "一年级"},
	//	{Name: "小猪妖5", Age: 9, Grade: "一年级"},
	//	{Name: "小猪妖6", Age: 10, Grade: "一年级"},
	//}
	//db.Create(&stu)

	//明一个名为 students 的变量，其类型是 model.Students 结构体的切片（slice）。
	var students []model.Students
	// 声明一个长度为10的数组
	//var students [10]model.Students
	//查询所有数据
	tx := db.Find(&students)
	if tx.Error != nil {
		panic(tx.Error)
	}
	for _, itm := range students {
		fmt.Println(itm)
	}

	//查询大于18岁的所有学生信息
	//tx := db.Where("Age > ?", 18).Find(&students)
	//if tx.Error != nil {
	//	panic(tx.Error)
	//}
	//for _, student := range students {
	//	fmt.Println(student)
	//}

	//marshal, err := json.Marshal(students)
	//if err != nil {
	//	log.Fatalln("JSON 编码错误", err)
	//}
	//
	////打印 JSON 字符串
	//fmt.Println("\n多条记录 JSON：")
	//fmt.Println(string(marshal))

	//students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//db.Where("Name = ?", "张三").Select("Grade").Updates(&model.Students{Grade: "四年级"})
	//
	//find := db.Find(&students)
	//
	//if find.Error != nil {
	//	panic(find.Error)
	//}
	//for _, item := range students {
	//	fmt.Println(item)
	//}

	//ctx := context.Background()

	//students 表中年龄小于 15 岁的学生记录。
	//db.WithContext(ctx).Where("Age < ?", 15).Delete(&model.Students{})

	db.Delete(&model.Students{}, "Age < ?", 15)

	find := db.Find(&students)

	if find.Error != nil {
		panic(find.Error)
	}
	for _, item := range students {
		fmt.Println(item)
	}

}
