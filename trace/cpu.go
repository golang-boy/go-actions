package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包
)

func cpuIntensiveTask() {
	for {
		// 模拟 CPU 密集型任务
		_ = 1 + 1 // 简单的数学运算，保持 CPU 占用
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

	// 启动多个 CPU 密集型任务
	for i := 0; i < 4; i++ {
		go cpuIntensiveTask()
	}

	// 主程序睡眠，以便观察 CPU 占用
	select {}
	fmt.Println("主程序结束")
}
