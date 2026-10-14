package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
	once sync.Once
	conn string // 模拟数据库连接
}

// GetConnection 确保只初始化一次
func (db *Database) GetConnection() string {
	// Do 方法接收一个无参数函数。只有第一次调用时，这个函数才会被执行。
	// 后续所有并发调用都会阻塞，直到第一次执行完毕，然后直接返回。
	db.once.Do(func() {
		fmt.Println("正在初始化数据库连接 (耗时操作)...")
		time.Sleep(2 * time.Second) // 模拟耗时
		db.conn = "DB_CONNECTION_ESTABLISHED"
		fmt.Println("数据库连接初始化完成！")
	})
	return db.conn
}

func main() {
	db := &Database{}
	var wg sync.WaitGroup

	// 模拟 5 个并发请求同时尝试获取数据库连接
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("请求 %d: 获取到连接 -> %s\n", id, db.GetConnection())
		}(i)
	}
	wg.Wait()
}
