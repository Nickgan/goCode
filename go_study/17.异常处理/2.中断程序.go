package main

import (
	"fmt"
	"os"
)

func init() {

	_, err := os.ReadFile("./test.txt")

	if err != nil {
		panic(err.Error())
	}

}

func main() {
	fmt.Printf("hello ")
}
