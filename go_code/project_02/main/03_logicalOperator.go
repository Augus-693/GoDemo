package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_logicalOperator.go
 * @Author Augus Lee
 * @Date 2022/9/26 14:56
 * @Version 1.0
 */

/*
逻辑运算符
*/
func main() {
	var age int = 40
	//演示逻辑与&&
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}
	//演示逻辑或 ||
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//演示逻辑非 !
	if !(age > 30 && age < 50) {
		fmt.Println("ok5")
	}
	if !(age > 30 || age < 40) {
		fmt.Println("ok6")
	}
}
