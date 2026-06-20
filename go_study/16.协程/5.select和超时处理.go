package main

import (
	"fmt"
	"time"
)

var doneChan = make(chan struct{})

func event() {
	fmt.Println("event 执行开始....")
	time.Sleep(3 * time.Second)
	fmt.Println("event 执行结束....")

	close(doneChan)
}

func main() {

	go event()

	select {
	case <-doneChan:
		fmt.Println("携程执行完毕。。。")
	case <-time.After(1 * time.Second):
		fmt.Println("超时了")
		return
	}
}
