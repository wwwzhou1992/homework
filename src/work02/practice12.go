package main

/*
*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全
*
*/
import (
	"fmt"
	"sync"
	"time"
)

// 共享计数器结构，包含互斥锁和计数值
type Counter struct {
	value int
	mu    sync.Mutex // 互斥锁，保护value字段
}

// 安全地递增计数器
func (c *Counter) Increment() {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 确保锁被释放
	c.value++
}

// 获取当前计数器的值
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// 创建共享计数器
	counter := &Counter{}

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 记录开始时间
	start := time.Now()

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程执行1000次递增操作
			for j := 0; j < 1000; j++ {
				counter.Increment()

				// 为了演示并发效果，偶尔打印进度
				if j%200 == 0 {
					fmt.Printf("协程 %d: 第 %d 次递增\n", id, j)
				}
			}
			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 计算执行时间
	elapsed := time.Since(start)

	// 输出最终结果
	fmt.Printf("\n最终计数器值: %d\n", counter.Value())
	fmt.Printf("预期值: %d\n", 10*1000)
	fmt.Printf("执行时间: %v\n", elapsed)

	// 验证结果是否正确
	if counter.Value() == 10*1000 {
		fmt.Println("✅ 结果正确！")
	} else {
		fmt.Println("❌ 结果错误！存在数据竞争")
	}
}
