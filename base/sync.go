/*
 * 背景：
 *   在实际开发中，我们经常需要一个全局的 Map 来缓存数据。但由于多个 Goroutine 可能会同时读写这个 Map，直接使用原生 map 会导致 Fatal Error (concurrent map read and map write)。
 *
 * 需求：
 * 1. 定义一个名为 Cache 的结构体，内部包含：
 *   - 一个原生的 map[string]int 用于存数据。
 *   - 一个 sync.RWMutex 用于保护这个 map。
 * 2. 为 Cache 提供两个方法：
 *   - Set(key string, value int)：向缓存写入数据（使用写锁）。
 *   - Get(key string) (int, bool)：读取缓存数据并返回是否找到（使用读锁）。
 * 3. 在 main 函数中：
 *   - 初始化 Cache。
 *   - 使用 sync.WaitGroup 启动 10 个 Goroutine。
 *   - 其中 5 个 Goroutine 负责并发调用 Set 写入不同的数据（例如 Set("user_1", 10) 到 user_5）。
 *   - 另外 5 个 Goroutine 负责并发调用 Get 读取刚才写入的数据（可以适当加一点 time.Sleep 错开读写顺序，或者直接并发读）。
 *   - 主函数使用 WaitGroup 等待这 10 个 Goroutine 全部执行完毕后，再打印出缓存中当前的元素个数。
 *
 * 提示：
 * - 记住 sync.RWMutex 和 map 都是不能随意拷贝的，方法接收者请使用指针 *Cache。
 * - 在 Get 和 Set 中，不要忘记使用 defer 释放锁。
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]int
	mu   sync.RWMutex
}

// Set
func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

// Get
func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]
	return value, ok
}

func main() {
	c := &Cache{
		data: make(map[string]int),
	}
	var wg sync.WaitGroup

	// 启动5对 set / get
	for i := 0; i < 5; i++ {
		wg.Add(2)
		fmt.Printf("running %d\n", i)

		// invoke set
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("user_%d", num)
			c.Set(key, num*num)
		}(i)

		// invoke get
		go func(num int) {
			defer wg.Done()

			time.Sleep(1 * time.Second)

			key := fmt.Sprintf("user_%d", num)
			if result, ok := c.Get(key); ok {
				fmt.Printf("Get result with %s: %d\n", key, result)
			} else {
				fmt.Printf("Get error with key %s\n", key)
			}
		}(i)
	}

	// 等待结束
	wg.Wait()

	// 打印 Cache 内容
	for key, value := range c.data {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
