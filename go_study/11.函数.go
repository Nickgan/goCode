package main

import "fmt"

func main() {

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 55)
}

func sum(numList ...int) {

	fmt.Printf("numList==========> %T \n", numList)

	var sum int
	for n, i := range numList {
		sum += i

		fmt.Println("n=", n, "i=", i)
	}
	fmt.Println("sum=", sum)
}
