package main

import "fmt"

//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
/*
*
*ptr += 10 表示"将指针指向的值增加10"
ptr += 10 表示"将指针本身的值（内存地址）增加10个单位的偏移量"
*/
func custom01(p *int) *int {
	*p = *p + 10
	return p
	//但是函数内部并没有修改指针指向的值，而是返回了一个新的值，但是调用处没有接收返回值，所以a还是原来的值。
	//return *p +10
}

/**
1. 函数的目的不同
返回 int：表示函数是一个"计算函数"，它接收输入并返回计算结果，但不修改原始数据

返回 *int：表示函数可能分配了新内存或返回指向某个数据的指针

2. 语义区别
int 返回值：返回的是值的副本

*int 返回值：返回的是指向某个内存位置的指针
*/

func custom02(p *int) int {
	*p = *p + 10
	return *p
	//但是函数内部并没有修改指针指向的值，而是返回了一个新的值，但是调用处没有接收返回值，所以a还是原来的值。
	//return *p +10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func doubleSlice01(slice *[]int) {
	for k, v := range *slice {
		fmt.Printf("k:%d, v:%d\n", k, v)
		//这种方式无法实现需求，因为for range 循环中的变量 v 是值的副本，而不是原始元素的引用
		//这里修改的是副本v，不是切片中的元素
		v = v * 2
		fmt.Printf("k:%d, v:%d\n", k, v)
	}
}

func doubleSlice02(slice *[]int) {
	/*	for i:= range *slice {
		(*slice)[i] *= 2
	}*/
	for k, v := range *slice {
		fmt.Printf("k:%d, v:%d\n", k, v)
		(*slice)[k] *= 2
		fmt.Printf("k:%d, v:%d\n", k, (*slice)[k])
	}
}

/**
核心区别
特性	数组(Array)	切片(Slice)
大小	固定长度，声明时确定	动态长度，可随时扩展
类型定义	[n]T (n是长度，T是元素类型)	[]T (T是元素类型)
值传递	值类型(赋值和传参时复制整个数组)	引用类型(底层数组共享)
内存分配	在栈上分配(通常)	在堆上分配(通常)
初始化	var arr [3]int	s := make([]int, 3)


package main

import "fmt"

func main() {
    // 数组示例
    var arr [3]int           // 声明一个长度为3的整型数组
    arr = [3]int{1, 2, 3}    // 数组初始化
    fmt.Println("数组:", arr)

    // 切片示例
    slice := make([]int, 3)  // 创建一个长度为3的切片
    slice[0] = 1
    slice[1] = 2
    slice[2] = 3
    fmt.Println("切片:", slice)

    // 切片可以动态扩展
    slice = append(slice, 4, 5)
    fmt.Println("扩展后的切片:", slice)

    // 从数组创建切片
    arrSlice := arr[1:3]     // 包含arr[1]和arr[2]
    fmt.Println("从数组创建的切片:", arrSlice)

    // 修改切片会影响底层数组
    arrSlice[0] = 99
    fmt.Println("修改切片后的数组:", arr)

    // 长度与容量
    fmt.Printf("切片长度: %d, 容量: %d\n", len(slice), cap(slice))

    // 数组是值类型
    arr2 := arr
    arr2[0] = 100
    fmt.Println("原始数组:", arr)
    fmt.Println("复制的数组:", arr2)
}

关键点说明
数组长度固定，而切片长度可变

数组是值类型，赋值会创建副本；切片是引用类型，赋值共享底层数组

切片有容量(capacity)的概念，表示底层数组可容纳的元素总数

可以使用make()创建切片，或从数组/其他切片创建

使用append()函数可以向切片添加元素，必要时会自动扩容

使用建议
需要固定大小时使用数组

需要动态大小或不确定大小时使用切片

切片更灵活，在实际开发中使用频率更高

理解数组和切片的区别对于编写高效、正确的Go代码至关重要。


*/

func main() {
	/*	a := 10
		custom01(&a)
		fmt.Println("a = ", a)*/

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片：", numbers)

	doubleSlice02(&numbers)
	fmt.Println(numbers)

}
