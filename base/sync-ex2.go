package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// 确保任务结束时，通知 WaitGroup
	defer wg.Done()

	fmt.Printf("Worker %d 开始干活\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d 干完了\n", id)
}

func main() {
	var wg sync.WaitGroup

	// 启动 3 个 worker
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // 启动前 Add
		go worker(i, &wg) // 传递指针
	}

	fmt.Println("主函数：等待所有工人完成...")
	wg.Wait() // 优雅等待，无需 Sleep
	fmt.Println("主函数：所有人完工，下班！")
}
