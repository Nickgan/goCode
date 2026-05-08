package main

import (
	"fmt"
	"time"
)

func sing() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
}

func main() {

	now := time.Now()
	go sing()
	go sing()
	go sing()
	go sing()
	go sing()
	time.Sleep(2 * time.Second)
	fmt.Println("main结束", time.Now().Sub(now))
}
