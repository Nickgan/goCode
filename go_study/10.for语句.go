package main

import (
	"fmt"
	"time"
)

func main() {

	// 每隔一秒打印时间
	for {
		fmt.Println(time.Now().Format("2006_01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
