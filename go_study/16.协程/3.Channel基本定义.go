package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个全局的管道变量moneyChan, 该变量只能存储int类型，容量为100
var moneyChan chan int = make(chan int, 100)

func pay(name string, money int, wait1 *sync.WaitGroup) {

	fmt.Printf("%s 开始购物 \n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束 \n", name)

	moneyChan <- money
	wait1.Done()

}

func main() {

	var wait1 = sync.WaitGroup{}
	wait1.Add(3)

	go pay("张三", 2, &wait1)
	go pay("王五", 3, &wait1)
	go pay("李四", 3, &wait1)

	go func() {
		//写法1：
		//defer close(moneyChan)
		//wait1.Wait() //这句代码执行完表示所有的goroutine都执行完毕了，可以关闭管道了（就是所有存放值的操作都完毕了，可以关闭了）

		//写法2：
		wait1.Wait()
		close(moneyChan)
	}()

	// 遍历管道，获取管道中的值，存放到一个切片中
	var moneyList = []int{}
	for i := range moneyChan {
		moneyList = append(moneyList, i)
	}

	// 遍历切片，求和
	var sum int
	for i := range moneyList {
		sum += moneyList[i]
	}

	fmt.Println("总金额==============>", sum)

	fmt.Println("--------------都购物结束-------------")

}
