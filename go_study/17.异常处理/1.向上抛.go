package main

import (
	"errors"
	"fmt"
)

func parent() (err error) {
	err = methond1()
	return
}

func methond1() error {
	return errors.New("出错了")
}

func main() {

	//defer func() {
	//	if e := recover(); e != nil {
	//		fmt.Println("发生异常，被recover捕获了")
	//	}
	//}()

	err := parent()
	//panic(err.Error())
	fmt.Println("===============>", err.Error())
}
