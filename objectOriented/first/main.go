package main

import (
	"fmt"
	"math"
)

/**
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

// Shape接口定义
// 接口包含两个方法：Area()和Perimter() 都返回float64类型
type Shape interface {
	Area() float64      //面积
	Perimeter() float64 //周长
}

// Rectangle（矩形）结构体定义
// 表示矩形，包含宽度和高度两个字段
type Rectangle struct {
	Width  float64 //宽
	Height float64 //高
}

// Area方法实现
// 接收者为Rectangle类型，计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // 面积 = 宽度 * 高度
}

// Perimeter方法实现
// 接收者Rectangle类型，计算矩形的周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height) //周长 = 2 * （宽度 + 高度）
}

// Circle（圆）结构定义
// 表示圆形，包含半径字段
type Circle struct {
	Radius float64 //圆的半径
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area方法实现
// 接收者为Circle类型，计算圆形的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	//创建Rectangle实例，初始化宽度为5，高度为3
	rect := Rectangle{Width: 5, Height: 3}

	//创建Circle实例，初始化半径为4
	circle := Circle{Radius: 4}

	//打印矩形的信息
	fmt.Println("矩形 - 宽度 ： %.2f,高度：%.2f\n", rect.Width, rect.Height)
	fmt.Printf("面积：%.2f\n", circle.Area())      //调用矩形的Area方法
	fmt.Printf("周长：%.2f\n\n", rect.Perimeter()) //调用矩形的Perimeter方法

	//打印圆形的信息
	fmt.Printf("圆形 - 半径：%.2f\n", circle.Radius) //圆形的半径
	fmt.Printf("面积%v\n", circle.Area())           //调用圆形的Area方法
	fmt.Printf("周长：%.2f\n", rect.Perimeter())     //调用圆形的Perimeter方法

	//演示使用Shape接口处理不同形状
	fmt.Printf("\n使用Shape接口：")
	//创建一个Shape类型的切片，包含矩形和椭圆
	shapes := []Shape{rect, circle}
	for _, shape := range shapes {
		fmt.Printf("圆形面积：%v,周长：%v\n", shape.Area(), shape.Perimeter())
	}

}
