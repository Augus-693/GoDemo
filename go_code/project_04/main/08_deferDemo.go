package main

import "fmt"

/**
 * @Project GoDemo
 * @File 08_deferDemo.go
 * @Author Augus Lee
 * @Date 2022/10/7 17:24
 * @Version 1.0
 */

func sum(n1 int, n2 int) int {
	//当执行 defer 时，暂时不执行，会将 defer 后面的语句压入独立的栈（defer 栈）
	//当函数执行完毕时，再从 defer 栈中按照先入后出的方式出栈
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)

	res := n1 + n2
	fmt.Println("ok3 res = ", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res = ", res)
}
