package main

import "fmt"

func main() {
	ttest(4)
}

func ttest(n int) {
	if n > 2 {
		n--
		ttest(n)
	} else {
		fmt.Println("n = ", n)
	}

}
