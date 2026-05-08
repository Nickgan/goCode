package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int, 100)
var nameChane1 = make(chan string, 100)
var doneChan = make(chan struct{})

func pay1(name string, money int, wait1 *sync.WaitGroup) {

	fmt.Printf("%s 开始购物 \n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束 \n", name)

	moneyChan1 <- money
	nameChane1 <- name

	wait1.Done()

}

func main() {

	var wait1 sync.WaitGroup
	wait1.Add(3)

	go pay1("张三", 2, &wait1)
	go pay1("王五", 3, &wait1)
	go pay1("李四", 3, &wait1)

	go func() {
		defer close(nameChane1)
		defer close(moneyChan1)
		defer close(doneChan)
		wait1.Wait()
	}()

	var moneyList []int
	var nameList []string
	for {
		select {
		case money := <-moneyChan1:
			moneyList = append(moneyList, money)

		case name := <-nameChane1:
			nameList = append(nameList, name)
		}

	}

	fmt.Println(moneyList)
	fmt.Println(nameList)
	fmt.Println("--------------都购物结束-------------")

}
