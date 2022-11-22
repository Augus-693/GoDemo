package main

import "fmt"

/**
 * @Project GoDemo
 * @File 02_arrayUseDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 15:46
 * @Version 1.0
 */

func main() {

	//数组下标的使用
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素的值:", i+1)
		fmt.Scanln(&score[i])
	}

	//变量数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d] = %v\n", i, score[i])
	}

	//四种初始化数组的方式
	var numArray1 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArray1=", numArray1)

	var numArray2 = [3]int{1, 2, 3}
	fmt.Println("numArray2=", numArray2)

	var numArray3 = [...]int{4, 5, 6}
	fmt.Println("numArray3=", numArray3)

	var numArray4 = [...]int{1: 300, 0: 200, 2: 100}
	fmt.Println("numArray4=", numArray4)

	//类型推导
	strArr5 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr5=", strArr5)
}
