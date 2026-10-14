/*
 * 背景：
 *   - 在微服务中，配置通常在程序启动时从远程服务器加载，这个过程比较慢。我们需要确保无论多少个 Goroutine 并发请求配置，远程加载动作只执行一次，并且加载完成后，其他请求能立刻拿到结果。
 *
 * 需求：
 * 1. 定义一个 ConfigManager 结构体，包含：
 *   - 一个 sync.Once 用于控制初始化。
 *   - 一个 atomic.Bool (或普通 bool + Mutex，推荐尝试 atomic.Bool) 用于标记配置是否已加载完成。
 *   - 一个 map[string]string 用于存储配置项。
 * 2. 实现 GetConfig(key string) string 方法：
 *   - 内部使用 sync.Once 触发 loadConfig() 方法。
 *   - loadConfig() 方法应模拟耗时 1 秒的加载过程（time.Sleep），打印 "Loading global config..."，并填充 map 数据（例如 {"app_name": "MyApp", "version": "1.0"}）。
 *   - 加载完成后，将 atomic.Bool 标记为 true。
 *   - 最后从 map 中返回对应的 key 的值。
 * 3. 在 main 函数中：
 *   - 启动 5 个 Goroutine，并发调用 GetConfig("app_name")。
 *   - 使用 sync.WaitGroup 等待它们全部完成。
 *   - 验证目标：控制台应该只打印一次 "Loading global config..."，但 5 个 Goroutine 都能成功打印出 "MyApp"。
 *
 * 提示：
 * - atomic.Bool 的用法：var flag atomic.Bool，设置用 flag.Store(true)，读取用 flag.Load()。
 * - sync.Once 的 Do 方法非常强大，你只需要把 loadConfig 的逻辑（或对其的调用）放在 Do 的闭包里即可。
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ConfigManager struct {
	once   sync.Once
	loaded atomic.Bool
	config map[string]string
}

// GetConfig
func (cf *ConfigManager) GetConfig(key string) (string, error) {
	cf.once.Do(cf.loadConfig)

	if cf.loaded.Load() {
		return cf.config[key], nil
	} else {
		return "", fmt.Errorf("配置尚未加载")
	}
}

// loadConfig
func (cf *ConfigManager) loadConfig() {
	fmt.Println("Loading global config...")

	cf.loaded.Store(true)
	cf.config["app_name"] = "MyApp"
}

func main() {
	var wg sync.WaitGroup
	cf := &ConfigManager{
		config: make(map[string]string),
	}

	// 启动 5个 Goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()

			if value, err := cf.GetConfig("app_name"); err != nil {
				fmt.Printf("[%d] get config error with app_name: %s", num, err)
			} else {
				fmt.Printf("[%d] app_name -> %s\n", num, value)
			}
		}(i)
	}

	wg.Wait()
}
