package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wait1 = sync.WaitGroup{}
)

func sign1(i int) {
	wait1.Add(1)
	fmt.Println("唱歌=============>", i)
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束=======>", i)
	wait1.Done()
}

func main() {
	var startTime = time.Now()
	go sign1(1)
	go sign1(1)
	go sign1(1)
	go sign1(1)

	wait1.Wait()
	fmt.Println("总耗时：", time.Now().Sub(startTime))
}
