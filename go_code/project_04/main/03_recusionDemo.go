package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_recusionDemo.go
 * @Author Augus Lee
 * @Date 9/29/2022 7:08 PM
 * @Version 1.0
 */
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n = ", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n = ", n)
	}
}
func main() {
	//test(4)
	test2(4)
}
