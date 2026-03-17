package main

import "fmt"

func main() {

	var str string = "hello中文"
	fmt.Println(len(str))

	for i := range str {
		fmt.Printf("%c\n", str[i])
	}

}
