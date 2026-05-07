package main

import "fmt"

func main() {

	h1(4)()
}

// -------------------------高阶函数

func h1(i int) func() {

	var f = func() {
		fmt.Println("h1")
	}

	var m = map[int]func(){
		1: func() {
			fmt.Println("1")
		},
		2: func() {
			fmt.Println("2")
		},
		3: func() {
			fmt.Println("3")
		},
		4: f,
	}

	return m[i]
}

// --------------------------参数-----

func m1() {

}

func m2(n1 int) {

}

func m3(n1, n2 int) {

}

// --------------------------返回值

func r1() {

}

func f2() int {

	return 1
}

func f3() (int, int) {

	return 1, 2
}

func f4() (rt int) {
	rt = 100
	return
}

//   * 声明指针比如： var p *int
