package main

import "fmt"

/**
 * @Project GoDemo
 * @File 07_addUpperDemo.go
 * @Author Augus Lee
 * @Date 2022/10/7 16:29
 * @Version 1.0
 */

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
