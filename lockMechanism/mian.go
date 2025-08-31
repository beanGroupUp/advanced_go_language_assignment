package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

/**
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
// getGoroutineID 函数用于获取当前goroutine的ID
func getGoroutineID() int {
	var buf [64]byte                  // 创建一个64字节的缓冲区，用于存储堆栈信息
	n := runtime.Stack(buf[:], false) // 获取当前goroutine的堆栈信息，false表示只获取当前goroutine的信息

	// 处理堆栈字符串：移除"goroutine "前缀，然后按空格分割字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]

	// 将字符串形式的ID转换为整数
	id, err := strconv.Atoi(idField)
	if err != nil {
		// 如果转换失败，抛出异常
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id // 返回goroutine ID
}

// Counter 结构体定义了一个带有互斥锁保护的计数器
type Counter struct {
	mu    sync.Mutex // 互斥锁，用于保护count字段的并发访问
	count int        // 计数器值
}

// Increment 方法用于安全地递增计数器值，并返回goroutine ID和递增前的值
func (c *Counter) Increment() (int, int) {
	c.mu.Lock()         // 获取互斥锁，如果锁已被其他goroutine持有，则阻塞等待
	defer c.mu.Unlock() // 使用defer确保在函数返回时释放互斥锁

	gid := getGoroutineID() // 获取当前goroutine的ID

	oldValue := c.count // 保存递增前的计数器值

	c.count++ // 递增计数器值

	return gid, oldValue // 返回goroutine ID和递增前的值
}

/*
func main() {
	// 创建计数器实例
	var counter Counter

	// 创建WaitGroup实例，用于等待所有goroutine完成
	var wg sync.WaitGroup

	// 创建一个映射，用于记录每个goroutine新增的次数
	goroutineAdditions := make(map[int]int)

	// 创建一个互斥锁，用于保护goroutineAdditions映射的并发访问
	var additionsMutex sync.Mutex

	// 循环创建10个goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1) // 增加WaitGroup的计数器

		// 启动一个匿名函数的goroutine
		go func() {
			defer wg.Done() // 使用defer确保在函数返回时减少WaitGroup的计数器

			// 每个goroutine循环1000次，递增计数器
			for j := 0; j < 1000; j++ {
				// 调用Increment方法递增计数器，并获取goroutine ID和递增前的值
				gid, oldValue := counter.Increment()

				// 获取additionsMutex锁，保护对goroutineAdditions映射的访问
				additionsMutex.Lock()
				// 增加当前goroutine的计数器
				goroutineAdditions[gid] = goroutineAdditions[gid] + 1
				// 释放additionsMutex锁
				additionsMutex.Unlock()

				// 打印每次递增的详细信息
				fmt.Printf("Goroutine %d: 将计数器从 %d 增加到 %d\n", gid, oldValue, oldValue+1)
			}
		}()
	}

	wg.Wait() // 等待所有goroutine完成

	// 打印每个goroutine新增的数据量
	fmt.Println("\n每个 Goroutine 新增的数据量:")
	// 遍历goroutineAdditions映射，打印每个goroutine的新增次数
	for gid, count := range goroutineAdditions {
		fmt.Printf("Goroutine %d: 新增了 %d 次\n", gid, count)
	}

	// 打印最终计数器的值
	fmt.Printf("\n最终计数器值: %d\n", counter.count)
}
*/

func main() {
	//使用int64类型的变量作为计数器，初始化为0
	var counter int64

	//使用waitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	//启动10个goroutine
	for i := 0; i < 10; i++ {
		//增加waitGroup的计数器
		wg.Add(1)

		// 启动一个goroutine，使用匿名函数并立即调用
		// 将循环变量i作为参数传递给匿名函数，避免闭包捕获问题
		go func(id int) {
			//在函数退出时，减少waitgroup的计数器
			defer wg.Done()
			// 每个goroutine内部循环1000次，执行递增操作
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器
				// atomic.AddInt64()函数会原子地将delta(这里是1)添加到counter的地址指向的值
				// 这个操作是线程安全的，不需要额外的锁机制
				atomic.AddInt64(&counter, 1)
			}
			// 打印当前 goroutine 完成的消息
			// 打印当前goroutine完成的消息
			// 使用参数id而不是循环变量i，确保每个goroutine打印正确的ID
			fmt.Printf("Goroutine %d 完成了 1000 次递增操作\n", id)
		}(i)
		//}(i) 的作用详解
		/**
		在您提供的代码中，}(i) 这一行有两个重要作用：

		1. 立即调用匿名函数
		在 Go 中，当我们定义一个匿名函数后，需要在后面加上 () 来立即调用它。所以 }(i) 中的 } 是匿名函数的结束，而 (i) 表示立即调用这个匿名函数并传递参数 i。

		2. 避免闭包捕获问题
		这是更重要的原因。在 Go 的循环中创建 goroutine 时，如果直接在 goroutine 中使用循环变量，会出现一个常见的陷阱：

		错误示例（会产生问题）：
		for i := 0; i < 10; i++ {
		    go func() {
		        fmt.Println(i) // 这里所有 goroutine 可能会打印相同的值（通常是10）
		    }()
		}

		正确做法（使用参数传递）：
		for i := 0; i < 10; i++ {
		    go func(id int) {
		        fmt.Println(id) // 每个 goroutine 会打印不同的值（0-9）
		    }(i) // 将 i 作为参数传递给匿名函数
		}
		为什么需要这样做？
		变量捕获：Go 的闭包会捕获外部变量的引用，而不是值

		执行时机：goroutine 的启动和执行是异步的，循环可能已经完成，此时所有 goroutine 看到的都是循环结束后的 i 值

		数据竞争：多个 goroutine 可能同时读取和修改同一个变量

		解决方案
		通过将循环变量作为参数传递给匿名函数，每个 goroutine 会获得该变量的一个副本，这样就避免了所有 goroutine 共享同一个变量引用的问题。
		go func(id int) {
		    // 这里的 id 是 i 的一个副本，每个 goroutine 都有自己的 id 值
		    fmt.Printf("Goroutine %d 启动了\n", id)
		}(i) // 将当前的 i 值传递给函数
		}(i) 的作用是：

		立即调用定义的匿名函数

		将循环变量 i 的值作为参数传递给匿名函数

		确保每个 goroutine 获得自己独立的变量副本，避免数据竞争和意外的共享状态

		这是 Go 并发编程中的一个重要模式，确保了在循环中创建 goroutine 时的正确行为
		*/
	}

	//等待所有goroutine完成
	// 调用Wait()方法阻塞主goroutine，直到所有其他goroutine完成
	// 当WaitGroup的计数器归零时，Wait()方法会返回
	wg.Wait()
	//输出最终的计数器的值
	// 所有goroutine完成后，打印最终的计数器值
	// 由于使用了原子操作，最终值应该是准确的10000
	fmt.Printf("counter = %d\n", counter)
}
