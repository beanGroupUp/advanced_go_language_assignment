package main

import "fmt"

/**
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/

// 定义Person结构体
type Person struct {
	Name string
	Age  int
}

// 定义Employee结构体，组合Person
type Employee struct {
	Person     //匿名嵌入，实现组合
	EmployeeID string
}

// 为Employee实现PrintInfo方法
func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名：%v\n员工年龄：%v\n员工ID：%v\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	//创建Employee实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  20,
		},
		EmployeeID: "123",
	}
	emp.PrintInfo()
}
