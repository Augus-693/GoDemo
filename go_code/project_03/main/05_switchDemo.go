package main

import "fmt"

/**
 * @Project GoDemo
 * @File 05_switchDemo.go
 * @Author Augus Lee
 * @Date 9/28/2022 4:56 PM
 * @Version 1.0
 */

func main() {
	var key byte
	fmt.Println("请输入一个字符 a, b, c, d, e, f, g")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	default:
		fmt.Println("输入错误")
	}

	//Type Switch：switch 语句还可以被用于type-switch 来判断某个interface 变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型是 %T\n", i)
	case int:
		fmt.Println("x是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int):
		fmt.Println("x 是 func(int) 型")
	case string, bool:
		fmt.Println("x 是 string 或 bool 型")
	default:
		fmt.Println("x 是 未知型")
	}
}
