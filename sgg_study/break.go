package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s1()

	// 生成1-100内的随机数
	//for i := 0; i < 100; i++ {
	//	rand.Seed(time.Now().Unix())
	//	fmt.Println("n", rand.Intn(100)+1)
	//}
}

func s1() {
	var count int

lable1:
	for {
		rand.Seed(time.Now().Unix())

		var randomN = rand.Intn(100) + 1
		fmt.Println(randomN)

		count++
		if randomN == 99 {
			fmt.Println("count================>", count)
			break lable1
		}

	}
}
