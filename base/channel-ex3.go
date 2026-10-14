package main

import (
	"fmt"
	"time"
)

func mockAPI() chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second) // 模拟一个很慢的 API
		ch <- "API 返回结果"
	}()
	return ch
}

func main() {
	// time.After 会在指定时间后，向返回的 channel 发送当前时间
	timeout := time.After(2 * time.Second)

	select {
	case result := <-mockAPI():
		fmt.Println("成功:", result)
	case <-timeout:
		// 因为 mockAPI 需要 3 秒，而 timeout 只有 2 秒，所以这里会先被触发
		fmt.Println("失败: 请求超时！")
	}
}
