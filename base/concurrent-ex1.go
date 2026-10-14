package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	// 使用 Go 1.19+ 推荐的原子类型
	count atomic.Int64
}

func (c *SafeCounter) Inc() {
	// 原子加 1，无需加锁
	c.count.Add(1)
}

func (c *SafeCounter) Value() int64 {
	// 原子读取
	return c.count.Load()
}

func main() {
	var counter SafeCounter
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()

	fmt.Printf("最终结果: %d\n", counter.Value()) // 必定是 10000
}
