package main

import (
	"fmt"
	"sync"
	"time"
)

var wait111 sync.WaitGroup
var mp1 = sync.Map{}

func reader1() {

	for {
		fmt.Println(mp1.Load("time"))
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
