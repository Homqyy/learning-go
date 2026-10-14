/*
 * 背景：
 *   我们要实现一个数据处理流水线。有一个“数据源”会不断产生数据，我们需要启动多个“处理工人”来消费这些数据。但是，有些工人可能会“卡死”（处理超时）。
 *
 * 需求：
 *   1. 创建一个有缓冲的 Channel jobs (容量为 5)，用于传递整数任务。
 *   2. 创建一个无缓冲的 Channel results，用于接收处理结果。
 *   3. 启动 3个 Goroutine 作为工人（Worker）。每个工人从 jobs 中读取任务，计算其平方后，发送到 results。
 *     - 模拟异常：在工人内部，使用 time.Sleep 模拟处理时间。如果任务是偶数，sleep 1秒；如果是奇数，sleep 3秒（模拟卡死）。
 *   4. 在主 Goroutine 中，循环发送 1 到 5 的任务到 jobs。
 *   5. 核心挑战：在主 Goroutine 收集结果时，使用 select 和 time.After 设置一个 2秒的超时时间。
 *     - 如果 2 秒内收到了结果，打印结果。
 *     - 如果 2 秒超时，打印 "处理超时，跳过"，并继续等待下一个结果（直到收齐 5 个任务的结果或超时）。
 *
 * 提示：
 * - 发送任务后，记得 close(jobs)，这样工人处理完当前任务后就会退出。
 * - 收集结果时，可以用一个 for 循环，配合 select 来监听 results 和 timeout。
 */

package main

import (
	"fmt"
	"time"
)

// 处理任务
func job(jobCh chan int, resCh chan int) {
	for v := range jobCh {
		// 偶数
		if v%2 == 0 {
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(3 * time.Second)
		}
		resCh <- (v * v)
	}
}

func main() {
	// 创建 channel
	jobCh := make(chan int, 5)
	resCh := make(chan int)
	peopleNum := 3
	jobNum := 5

	// 启动工人
	for i := 0; i < peopleNum; i++ {
		go job(jobCh, resCh)
	}

	// 设置超时 2s
	timeoutCh := time.After(2 * time.Second)

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
		case <-timeoutCh:
			fmt.Println("timeout")
			return
		}
	}

	fmt.Println("所有任务处理完毕！")
}
