package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入 pprof 包
	"runtime"
	"sync"
)

var (
	mutex sync.Mutex
)

func main() {
	// 启动 pprof HTTP 服务器
	go func() {
		fmt.Println("pprof 监听在 6060 端口")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			fmt.Println("无法启动 pprof:", err)
		}
	}()

	var items = make(map[int]struct{})

	runtime.SetMutexProfileFraction(5)

	for i := 0; i < 1000_000; i++ {
		go func(i int) {
			mutex.Lock()
			items[i] = struct{}{}
			mutex.Unlock()
		}(i)
	}

	// 让主程序等待一段时间以便 goroutines 执行
	select {}
	fmt.Println("主程序结束")

}
