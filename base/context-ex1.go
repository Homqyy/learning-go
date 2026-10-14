package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟一个慢速的数据库查询或外部 API 调用
// 注意：ctx 必须是第一个参数！
func slowDatabaseQuery(ctx context.Context) (string, error) {
	// 模拟一个需要 5 秒才能完成的操作
	queryDone := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		queryDone <- "查询成功：用户数据"
	}()

	// 核心：使用 select 监听两个 channel
	select {
	case result := <-queryDone:
		// 正常完成
		return result, nil
	case <-ctx.Done():
		// 被取消或超时！
		// ctx.Err() 会返回 context.DeadlineExceeded 或 context.Canceled
		return "", fmt.Errorf("查询被中止，原因: %v", ctx.Err())
	}
}

func main() {
	// 1. 创建一个根 context
	ctx := context.Background()

	// 2. 派生出一个带有 2 秒超时的子 context
	// cancel 函数用于手动提前取消（虽然这里超时会自动取消，但最佳实践是声明并 defer 调用，防止内存泄漏）
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // 确保资源被释放

	fmt.Println("主函数：开始查询数据库...")

	// 3. 将 ctx 传递给子任务
	result, err := slowDatabaseQuery(ctx)

	if err != nil {
		fmt.Printf("主函数：查询失败 -> %v\n", err)
	} else {
		fmt.Printf("主函数：查询成功 -> %s\n", result)
	}
}
