package main

import (
	schedulerTask "Goroutine/utils"
)

// 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
/*func main() {
	//使用缓冲通道
	ch := make(chan struct{})
	//声明了一个等待组(WaitGroup)变量，它是实现多个goroutine(协程)同步的重要工具。
	var wg sync.WaitGroup
	//设置需要等待2个goroutine完成
	wg.Add(2)

	//打印奇数的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数：%d\n", i)
			//通知偶数协程
			ch <- struct{}{}
			//防止最后一次接收阻塞（因为到9就结束）
			if i < 9 {
				<-ch
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch //奇数协程通知
			fmt.Printf("偶数：%d\n", i)
			//防止最后一次发送阻塞
			if i < 10 {
				//通知奇数协程
				ch <- struct{}{}
			}
		}
	}()

	wg.Wait()
	close(ch)
}*/

/*
func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数：", i)
			ch <- struct{}{}
			if i < 9 {
				<-ch
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			//等待奇数协程通知
			<-ch
			fmt.Println("偶数：", i)
			if i < 10 {
				ch <- struct{}{}
			}
		}
	}()

	wg.Wait()
	//defer在函数运行完毕后执行
	defer close(ch)
}
*/

/*func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数：", i)
			ch <- 1
			if i < 9 {
				<-ch
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch
			fmt.Println("偶数：", i)
			if i < 10 {
				ch <- 1
			}
		}
	}()

	wg.Wait()
	defer close(ch)
}
*/

/**
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

func main() {
	schedulerTask.TestTask()
}
