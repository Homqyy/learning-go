package main

import (
	"fmt"
	"time"
)

// 参数限制为“只写” channel (chan<- int)
func producer(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("生产者：发送了 %d\n", i)
	}
	// 生产完毕，关闭 channel，通知消费者
	close(ch)
}

// 参数限制为“只读” channel (<-chan int)
func consumer(ch <-chan int) {
	// 遍历 channel，直到 channel 被关闭且为空
	for val := range ch {
		fmt.Printf("消费者：收到了 %d\n", val)
	}
	fmt.Println("消费者：channel 已关闭，退出。")
}

func main() {
	// 创建容量为 2 的有缓冲 channel
	ch := make(chan int, 2)

	go producer(ch)
	go consumer(ch)

	// 稍微等待一下，防止主函数提前退出（实际项目中应使用 WaitGroup，这里为了简单演示）
	// 注意：这里只是为了让控制台打印出来，实际开发不推荐用 Sleep 做同步
	time.Sleep(1 * time.Second)
}
