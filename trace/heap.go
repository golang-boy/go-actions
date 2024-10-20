package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包
	"runtime"
	"time"
)

func allocateMemory() {
	for i := 0; i < 1_000_000; i++ {
		// 创建大量临时对象
		_ = make([]byte, 1e6) // 每次分配 1 MB
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

	// 启动一个 goroutine 来监控内存使用情况
	go func() {
		for {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			fmt.Printf("Alloc = %v MiB", memStats.Alloc/1024/1024)
			fmt.Printf("\tTotalAlloc = %v MiB", memStats.TotalAlloc/1024/1024)
			fmt.Printf("\tSys = %v MiB", memStats.Sys/1024/1024)
			fmt.Printf("\tNumGC = %v\n", memStats.NumGC)
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("开始内存分配...")
	for {
		allocateMemory()
		time.Sleep(1 * time.Second) // 每秒分配一次
	}
}
