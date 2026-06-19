package main

import (
	"fmt"
	"time"
)

func main() {

	for1()

	// 每隔一秒打印时间
	for {
		fmt.Println(time.Now().Format("2006_01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}

func for1() {

	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
