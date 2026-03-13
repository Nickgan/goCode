package main

import (
	"fmt"
)

type myGetSum func(int, int) int

func main() {

	sum := myFun(getSum, 1, 2)
	fmt.Println("sum==========>", sum)

	f := getSum
	fmt.Printf("f==========> %T \n", f)

	sum = myFun2(f, 1, 2)

	fmt.Println("sum==========>", sum)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func myFun(my myGetSum, n1 int, n2 int) int {
	return my(n1, n2)
}

func myFun2(sum myGetSum, n1 int, n2 int) (a int) {
	a = sum(n1, n2)
	return
}
