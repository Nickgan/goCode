package main

import (
	"fmt"
	"time"
)

func main() {

	// 获取当前时间
	now := time.Now()
	fmt.Println("当前时间：", now)

	// 格式化当前时间
	fmt.Println("当前时间：", now.Format("2006-01-02 15:04:05"))
	fmt.Println("当前时间：", now.Format("2006-01-02 15:04"))
	fmt.Println("当前时间：", now.Format("2006-01-02"))
	fmt.Println(now.Format("2006_01-02 15:04:05"))

	// 解析时间(layout：根据什么样的格式解析，后面一个参数是要解析的时间字符串)
	t, _ := time.Parse("2006-01-02 15:04:05", "2020-01-01 12:01:01")
	fmt.Println("解析时间：", t)
	fmt.Println("解析时间：", t.Format("2006-01-02 15:04:05"))

	// 获取时间间隔
	duration := time.Since(t)
	fmt.Println("时间间隔：", duration)

	tickerDemo()
}

func tickerDemo() {
	myTicker := time.NewTicker(3 * time.Second)
	for i := 0; i < 3; i++ {
		// <-ticker.C 从 ticker 的 channel 中取出下一次触发的时间，若尚未触发会在此处阻塞等待。
		<-myTicker.C
		fmt.Printf("%v====>ticker执行\n", i)
	}

	defer myTicker.Stop()
	fmt.Println("Main方法执行完毕...")
}
