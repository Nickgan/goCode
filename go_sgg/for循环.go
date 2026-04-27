package main

import "fmt"

func main() {

	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//var j = 1
	//for j <= 10 {
	//	fmt.Println("j============>", j)
	//	j++
	//}

	//for {
	//	fmt.Println("infinite loop")
	//}

	//var x interface{}
	//
	//var y = 10.0
	//
	//x = y
	//
	//switch n := x.(type) {
	//case types.Nil:
	//	fmt.Println("n is nil")
	//case float64:
	//	fmt.Println("n is float64", n)
	//
	//}

	var str = "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

}
