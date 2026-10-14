/*
 * 背景：
 *   在第一阶段，如果主函数因为超时 return 了，Worker 还会因为向无缓冲 channel 发送数据而永久阻塞。现在我们要让 Worker 变得“聪明”，能够感知外部的取消信号。
 *
 * 需求：
 * 1.修改 job 函数的签名，让它接收一个 context.Context 作为第一个参数：func job(ctx context.Context, jobCh chan int, resCh chan int)。
 * 2. 在 job 函数内部，模拟奇数任务需要 3 秒处理时间。但是，不能只用 time.Sleep。你需要使用 select 同时监听 time.After (或 time.Sleep 的 channel 版本) 和 ctx.Done()。
 *   - 提示：可以使用 select 配合 time.After(3 * time.Second) 来模拟耗时，如果在这期间 <-ctx.Done() 被触发，则打印 "Worker 收到取消信号，退出" 并 return。
 * 3. 在 main 函数中：
 *   - 使用 context.WithTimeout 创建一个 2 秒超时的 context。
 *   - 将这个 context 传递给所有的 Worker。
 *   - 当主函数因为超时退出时，观察 Worker 是否能优雅地打印退出信息，而不是永久阻塞。
 */
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 处理任务
func job(ctx context.Context, wg *sync.WaitGroup, jobCh chan int, resCh chan int) {
	defer wg.Done()

	for v := range jobCh {
		var sleepTime time.Duration
		// 偶数
		if v%2 == 0 {
			sleepTime = 1
		} else {
			sleepTime = 3
		}

		timer := time.NewTicker(sleepTime * time.Second)
		defer timer.Stop()

		select {
		case <-timer.C:
			resCh <- (v * v)
		case <-ctx.Done():
			fmt.Println("Work被提前中止")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup

	// 设置超时 2s
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// 创建 channel
	jobCh := make(chan int, 5)
	resCh := make(chan int)
	peopleNum := 3
	jobNum := 5

	// 启动工人
	for i := 0; i < peopleNum; i++ {
		wg.Add(1)
		go job(ctx, &wg, jobCh, resCh)
	}

	// 发送任务
	for i := 0; i < jobNum; i++ {
		jobCh <- (i + 1)
	}
	close(jobCh)

	// 接收结果
	resNum := 0
	for resNum < jobNum {
		select {
		case result := <-resCh:
			fmt.Printf("recv a result: %d\n", result)
			resNum++
		case <-ctx.Done():
			fmt.Println("main timeout")
			goto done
		}
	}
done:
	wg.Wait()
	fmt.Println("所有任务处理完毕！")
}
