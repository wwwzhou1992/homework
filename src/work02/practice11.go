package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建一个缓冲大小为20的整数通道
	// 这意味着通道可以存储最多20个整数，而不需要立即被消费
	ch := make(chan int, 20)

	// 使用WaitGroup等待生产者和消费者协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 等待2个协程：生产者和消费者

	fmt.Println("程序开始，创建了缓冲大小为20的通道")
	fmt.Printf("初始通道状态: 长度=%d, 容量=%d\n\n", len(ch), cap(ch))

	// 生产者协程：向通道发送100个整数
	go producer(ch, &wg)

	// 消费者协程：从通道接收整数并打印
	go consumer(ch, &wg)

	// 等待两个协程完成
	wg.Wait()
	fmt.Println("\n程序结束")
}

// 生产者函数：向通道发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // 关闭通道，通知消费者没有更多数据了

	fmt.Println("生产者开始工作...")

	for i := 1; i <= 100; i++ {
		// 模拟一些处理时间
		time.Sleep(10 * time.Millisecond)

		// 发送整数到通道
		ch <- i

		// 打印发送状态
		fmt.Printf("生产者: 发送 %d, 通道状态: %d/%d\n",
			i, len(ch), cap(ch))
	}

	fmt.Println("生产者完成，已发送100个整数")
}

// 消费者函数：从通道接收整数并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("消费者开始工作...")
	count := 0

	// 使用range循环从通道接收数据，直到通道关闭
	for num := range ch {
		// 模拟一些处理时间（比生产者慢）
		time.Sleep(20 * time.Millisecond)

		count++
		fmt.Printf("消费者: 接收 %d, 已处理 %d/100, 通道状态: %d/%d\n",
			num, count, len(ch), cap(ch))
	}

	fmt.Printf("消费者完成，总共处理了 %d 个整数\n", count)
}
