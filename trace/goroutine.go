package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包
	"time"
)

func leakGoroutine() {
	for {
		// 模拟一些工作
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("pprof 监听在 6060 端口")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Println("无法启动 pprof:", err)
		}
	}()

	for i := 0; i < 1_000_00; i++ {
		go leakGoroutine() // 启动新 goroutine，但没有控制它的生命周期
		time.Sleep(100 * time.Millisecond)
	}

	// 主程序睡眠，确保 goroutine 有时间运行

	select {}

	fmt.Println("主程序结束")
}
