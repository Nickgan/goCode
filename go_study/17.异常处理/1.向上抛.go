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
	fmt.Println(parent())
}
