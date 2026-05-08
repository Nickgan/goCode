package main

import (
	"fmt"
	"sync"
	"time"
)

var wait11 sync.WaitGroup
var mp = map[string]string{}

var lock sync.Mutex

func reader() {

	for {
		lock.Lock()
		fmt.Println(mp["time"])
		lock.Unlock()
	}

	wait11.Done()
}
func writer() {

	for {
		lock.Lock()
		mp["time"] = time.Now().Format("15:04:05")
		lock.Unlock()
	}

	wait11.Done()
}

func main() {
	wait11.Add(2)
	go reader()
	go writer()
	wait11.Wait()

}
