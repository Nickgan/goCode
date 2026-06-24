package main

import "fmt"

//1：空接口可以接受任何类型
//2：换一种，任何类型都实现了空接口的定义

type NullInterface interface {
}

// MyPrint 空接口可以接受任何一种类型（第一种写法，直接写一个空接口）
//func MyPrint(a NullInterface) {
//	fmt.Println(a)
//}

// MyPrint 空接口可以接受任何一种类型（第二种写法，写一个interface{}）
func MyPrint(val interface{}) {
	fmt.Println(val)
}

// MyPrint 空接口可以接受任何一种类型（第三种写法，写一个any）
func MyPrint1(val any) {
	fmt.Println(val)
}

type MyCat struct {
	Name string
	Age  int
}

func emptyInterfaceDemo(value any) {

	if str, ok := value.(string); ok {
		fmt.Println("是字符串:", str)
	}

	switch v := value.(type) {
	case int:
		fmt.Println("是整数:", v)
	case string:
		fmt.Println("是字符串:", v)
	case []int:
		fmt.Println("是切片:", v)
	case map[string]interface{}:
		fmt.Println("是map[string]interface{}字典:", v)
	case map[string]int:
		fmt.Println("是map[string]int字典:", v)
	default:
		fmt.Println("类型未知:", v)

	}
}

func main() {
	MyPrint(MyCat{Name: "小猫", Age: 1})
	MyPrint(1)

	// 空接口
	fmt.Println("\n=== 空接口示例 ===")
	emptyInterfaceDemo(42)
	emptyInterfaceDemo("hello")
	emptyInterfaceDemo([]int{1, 2, 3})
	emptyInterfaceDemo(map[string]int{"a": 1})
	emptyInterfaceDemo(map[string]string{"a": "hello"})

}
