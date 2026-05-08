package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wait = sync.WaitGroup{}
)

func sign(i int) {
	fmt.Println("唱歌=============>", i)
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束=======>", i)
	wait.Done()
}

func main() {
	var count = 3
	var startTime = time.Now()
	wait.Add(count)
	//go sign()
	//go sign()
	//go sign()
	//go sign()

	for i := 0; i < count; i++ {
		go sign(i)
	}
	wait.Wait()
	fmt.Println("总耗时：", time.Now().Sub(startTime))
}
