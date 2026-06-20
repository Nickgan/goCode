package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	moneyChan2 = make(chan int)
	nameChan2  = make(chan string)
	doneChan2  = make(chan struct{}) // 用于关闭结束接收
)

func pay2(name string, money int, wait1 *sync.WaitGroup) {
	fmt.Println(name, "开始购物")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束, 花费金额：", money)

	moneyChan2 <- money
	nameChan2 <- name

	wait1.Done()
}

func main() {
	var wait1 sync.WaitGroup
	wait1.Add(3)
	go pay2("张三", 2, &wait1)
	go pay2("王五", 3, &wait1)
	go pay2("李四", 3, &wait1)

	// 关闭
	go func() {
		defer close(moneyChan2)
		defer close(nameChan2)
		defer close(doneChan2)
		wait1.Wait()
	}()

	// 接收
	var nameList []string
	var moneyList []int

	for {
		select { //select 基于事件监听，只要有一个事件就执行
		case money := <-moneyChan2:
			moneyList = append(moneyList, money)
		case name := <-nameChan2:
			nameList = append(nameList, name)
		case <-doneChan2:
			fmt.Println("所有用户购物结束")
			return
		}
	}

	fmt.Println("总金额：", moneyList)
	fmt.Println("用户名：", nameList)
}
