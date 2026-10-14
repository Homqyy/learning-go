package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("工人：开始干活...")
	time.Sleep(2 * time.Second) // 模拟耗时工作
	fmt.Println("工人：干完了！")

	// 向 done channel 发送信号，表示工作完成
	// 如果主函数没有在接收，这里会阻塞！
	done <- true
}

func main() {
	// 1. 创建一个无缓冲的 bool 类型 channel
	done := make(chan bool)

	// 2. 启动 goroutine
	go worker(done)

	// 3. 主 goroutine 阻塞等待，直到从 done 中接收到数据
	// 这起到了类似 C 语言中 pthread_join 的作用
	<-done

	fmt.Println("主函数：收到完成信号，程序退出。")
}
