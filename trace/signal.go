package main

import (
	"os"
	"os/signal"
)

func main() {

	notfiy := make(chan os.Signal, 1)

	signal.Notify(notfiy, os.Interrupt, os.Kill)

	<-notfiy //阻塞，直到收到信号
}
