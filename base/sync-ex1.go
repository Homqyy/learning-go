package main

import (
	"fmt"
	"sync"
)

// 安全的计数器结构体，将 Mutex 和 数据绑定在一起
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// 增加计数 (注意接收者是指针 *SafeCounter，防止 Mutex 被拷贝)
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	// 使用 defer 确保函数退出时一定会 Unlock
	defer c.mu.Unlock()
	c.value++
}

// 获取当前值
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// 启动 1000 个 Goroutine 并发增加计数器
	for i := 0; i < 1000; i++ {
		wg.Add(1) // 计数器 +1
		go func() {
			// defer Done() 确保无论发生什么，wg 都会减 1
			defer wg.Done()
			counter.Inc()
		}()
	}

	// 主线程阻塞等待，直到 wg 计数器归零
	wg.Wait()
	fmt.Printf("最终结果: %d\n", counter.Value()) // 一定是 1000
}
