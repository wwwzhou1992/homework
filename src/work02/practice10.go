package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个无缓冲的整数通道
	ch := make(chan int)

	// 使用WaitGroup等待两个协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 等待2个协程完成

	// 生产者协程：生成1到10的整数并发送到通道
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup
		defer close(ch) // 关闭通道，通知接收者没有更多数据了

		fmt.Println("生产者开始生成数据...")
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者发送: %d\n", i)
			ch <- i // 发送数据到通道
		}
		fmt.Println("生产者完成数据生成")
	}()

	// 消费者协程：从通道接收数据并打印
	go func() {
		defer wg.Done()

		fmt.Println("消费者开始接收数据...")
		// 方法1：使用for循环和通道关闭检测
		for {
			num, ok := <-ch
			if !ok {
				// 通道已关闭，没有更多数据
				fmt.Println("消费者检测到通道已关闭")
				break
			}
			fmt.Printf("消费者接收并处理: %d\n", num)
		}
		fmt.Println("消费者完成数据处理")
	}()

	// 等待两个协程都完成
	wg.Wait()
	fmt.Println("程序执行完毕")
}
