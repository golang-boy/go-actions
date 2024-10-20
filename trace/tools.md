#


##　trace

curl -o .\tracer.out http://localhost:6060/debug/pprof/trace?seconds=30
go tool trace .\tracer.out

协程的创建和销毁，协程的阻塞
网络i/o事件
系统调用事件
垃圾回收事件

## 找一段时间内go程序在干啥

###　分析延迟问题
