package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wait = sync.WaitGroup{}
)

func sign() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
	wait.Done()
}

func main() {
	wait.Add(4)
	go sign()
	go sign()
	go sign()
	go sign()
	wait.Wait()
}
