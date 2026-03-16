package main

import "fmt"

func main() {

	sum1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 55)
	fmt.Println("main sum==========>", sum1)
}

func sum(numList ...int) (sum int) {

	fmt.Printf("numList==========> %T \n", numList)

	for n, i := range numList {
		sum += i

		fmt.Println("n=", n, "i=", i)
	}
	fmt.Println("sum=", sum)
	return
}
