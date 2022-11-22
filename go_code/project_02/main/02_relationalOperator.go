package main

import "fmt"

/**
 * @Project GoDemo
 * @File 02_relationalOperator.go
 * @Author Augus Lee
 * @Date 2022/9/26 14:41
 * @Version 1.0
 */

/*
关系运算符（比较运算符）
*/
func main() {
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) //false
	fmt.Println(n1 != n2) //true
	fmt.Println(n1 > n2)  //true
	fmt.Println(n1 >= n2) //true
	fmt.Println(n1 < n2)  //false
	fmt.Println(n1 <= n2) //false
	flag := n1 > n2
	fmt.Println("flag = ", flag) //true
}
