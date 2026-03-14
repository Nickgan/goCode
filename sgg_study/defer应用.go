package main

import "fmt"

func main() {
	fmt.Println("hello world==========main")
	deferDemo()
	fmt.Println("hello golang==========main")
}
func deferDemo() {
	defer fmt.Println("===============defer111111")
	defer fmt.Println("===============defer222222")
	defer m22222()
	fmt.Println("hello golang=====deferDemo")
}

func m22222() {
	defer fmt.Println("===============defer333333")
}
