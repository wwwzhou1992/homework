package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 使用原子操作的无锁计数器
type AtomicCounter struct {
	value int64 // 必须使用int64类型，因为atomic包需要知道确切的大小
}

// 使用原子操作递增计数器
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// 获取当前计数器的值
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// 使用CAS（Compare-And-Swap）实现的安全递增
func (c *AtomicCounter) IncrementWithCAS() {
	for {
		current := atomic.LoadInt64(&c.value)
		if atomic.CompareAndSwapInt64(&c.value, current, current+1) {
			break
		}
		// 如果CAS失败，重试
		time.Sleep(time.Nanosecond) // 短暂休眠避免忙等待
	}
}

func main() {
	// 创建原子计数器
	counter := &AtomicCounter{}

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 记录开始时间
	start := time.Now()

	fmt.Println("开始执行原子计数器测试...")

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程执行1000次递增操作
			for j := 0; j < 1000; j++ {
				counter.Increment()

				// 每200次操作打印一次进度
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
	finalValue := counter.Value()
	expectedValue := int64(10 * 1000)

	fmt.Printf("\n=== 测试结果 ===\n")
	fmt.Printf("最终计数器值: %d\n", finalValue)
	fmt.Printf("预期值: %d\n", expectedValue)
	fmt.Printf("执行时间: %v\n", elapsed)

	// 验证结果是否正确
	if finalValue == expectedValue {
		fmt.Println("✅ 结果正确！原子操作保证了数据安全")
	} else {
		fmt.Printf("❌ 结果错误！期望 %d，实际得到 %d\n", expectedValue, finalValue)
	}
}
