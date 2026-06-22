package main

import (
	"fmt"
	"sync"
	"time"
)

var wait111 sync.WaitGroup

// 线程安全的map
var mp1 = sync.Map{}

func reader1() {

	for {
		v, ok := mp1.Load("time")
		fmt.Println(v, "==============", ok)
	}

	wait111.Done()
}
func writer1() {

	for {
		mp1.Store("time", time.Now().Format("15:04:05"))
	}

	wait111.Done()
}

func main() {
	wait111.Add(2)
	go reader1()
	go writer1()
	wait111.Wait()
}
