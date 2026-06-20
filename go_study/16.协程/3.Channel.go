package main

import "fmt"

//var ch = make(chan int, 10)

func main() {

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//}()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-ch)
	//}
	//
	//close(ch)

	var ch chan int
	ch = make(chan int, 10)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	for i := range ch {
		fmt.Println(i)
	}

}
