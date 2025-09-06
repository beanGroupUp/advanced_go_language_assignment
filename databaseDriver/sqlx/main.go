package main

import (
	"awesomeProject3/databaseDriver/sqlx/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
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
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	//注意赋值短变量声明 :=，这会在 init() 函数的作用域内创建一个新的局部变量 db，而不是赋值给包级别的全局变量 db。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功...")

	//db.AutoMigrate(&model.Employees{})
}

//func main() {
//	// 批量插入数据
//	employees := []model.Employees{
//		{
//			Name:       "小5",
//			Department: "技术部",
//			Salary:     20000,
//		},
//		{
//			Name:       "小6",
//			Department: "技术部",
//			Salary:     30000,
//		},
//		{
//			Name:       "小7",
//			Department: "市场部",
//			Salary:     5555,
//		},
//		{
//			Name:       "小8",
//			Department: "财务部",
//			Salary:     6666,
//		},
//	}
//	//在 GORM 的 Create(&employees) 中使用 &employees 是因为：
//	//允许回填数据：让 GORM 能够将数据库生成的值（如自增 ID）写回到原始变量中
//	//提高性能：传递指针比传递大型数据结构的副本更高效
//	//符合 API 设计：GORM 的方法期望接收指针参数
//	//保持一致性：与其他 GORM 方法的使用方式保持一致
//	//这种设计模式是 GORM 和其他许多 Go ORM 库的标准做法，它确保了数据操作的完整性和效率。
//	db.Model(&model.Employees{}).Create(&employees)
//	for _, employee := range employees {
//		fmt.Println(employee.ID)
//	}
//}

func main() {
	//初始化数据库连接
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败：%v,服务断开；", err)
	}
	defer db.Close()

	//测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal("无法连接到数据库:“%v,服务断开；", err)
	}

	//设置连接池大小
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(20)

	fmt.Println("链接到数据库；")

	//查询技术部所有员工
	department, err := queryEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatal("查询技术部员工失败：%v", err)
	}
	fmt.Printf("技术部员工人数:%v", len(department))
	for _, department := range department {
		fmt.Printf("ID：%v,姓名：%v,部门：%v,薪资：%v\n", department.ID, department.Name, department.Department, department.Salary)
	}

	employee, err := queryHighestPaidEmployee(db)
	if err != nil {
		fmt.Errorf("查询工资最高员工异常：%v", err)
	}
	fmt.Printf("工资最高员工：ID：%v,姓名：%v,部门：%v,薪资：%v\n", employee.ID, employee.Name, employee.Department, employee.Salary)

}

func queryEmployeesByDepartment(db *sqlx.DB, department string) ([]model.Employees, error) {
	var employees []model.Employees
	//使用select方法执行查询，并奖结果自动映射到Emplyees结构体切片中
	//sqlx的select方法极大简化了多行结果的映射过程
	query := "select id, name, department, salary from employees where department = ?"
	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, fmt.Errorf("查询员工信息失败 err:%v", err)
	}
	return employees, nil
}

// 返回值，不过不是指针结果 就返回不出来
func queryHighestPaidEmployee(db *sqlx.DB) (*model.Employees, error) {
	var employees model.Employees

	query := "select id, name, department, salary from employees order by salary desc limit 1"
	//使用 Select() 当你期望多条记录，将结果映射到切片
	//
	//使用 Get() 当你期望单条记录，将结果映射到单个结构体
	err := db.Get(&employees, query)
	if err != nil {
		return nil, fmt.Errorf("查询最高工资员工失败 err:%v", err)
	}
	return &employees, nil
}
