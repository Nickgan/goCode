package main

import (
	"fmt"
	"math"
)

func main() {

	var i int = 10
	var i1 int8 = 10
	var i2 int16 = 10
	var i3 int32 = 10
	var i4 int64 = 10
	var f1 float32 = 10.5
	var f2 float64 = 10.5

	fmt.Println(i, i1, i2, i3, i4, f1, f2)

	//字符类型
	var c1 byte = 'a'
	var c2 rune = '中' //可以支持中文字符（占4个字节)

	fmt.Printf("%c,%c \n", c1, c2)

	fmt.Println(c1, c2)

	// 最大值
	fmt.Println(math.MaxFloat32)

	//多行字符串(原样输出)
	var str = `asdfasdf
					adfadf
				adfasdfa
				adfasdf`

	fmt.Println(str)

	// 布尔类型
	var b1 bool = true
	fmt.Printf("%v", b1)

	//最大值
	fmt.Println(math.MaxInt32)

	fmt.Println(math.MaxFloat64)
}
