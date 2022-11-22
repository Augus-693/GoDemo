package main

import "fmt"

/**
 * @Project GoDemo
 * @File 05_methodExercise.go
 * @Author Augus Lee
 * @Date 2022/10/14 13:26
 * @Version 1.0
 */

//编写结构体(MethodUtils)

type MethodUtils struct {
}

// 编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形
func (m MethodUtils) print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 编写一个方法，提供m和 n 两个参数，方法中打印一个m*n 的矩形
func (m1 MethodUtils) print1(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// 编写一个方法算该矩形的面积(可以接收长 len，和宽 width)， 将其作为方法返回值
func (m2 MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

// 编写方法：判断一个数是奇数还是偶数
func (m3 MethodUtils) judgeNum(i int) {
	if i%2 == 0 {
		fmt.Println(i, "是偶数")
	} else {
		fmt.Println(i, "是奇数")
	}
}

func main() {
	var m MethodUtils
	m.print()
	var m1 MethodUtils
	m1.print1(6, 12)
	var m2 MethodUtils
	res := m2.area(10.2, 2.4)
	fmt.Printf("%.2f\n", res)
	var m3 MethodUtils
	m3.judgeNum(5)
}
