package main

import "fmt"

/**
 * @Project GoDemo
 * @File 06_fibonacci.go
 * @Author Augus Lee
 * @Date 10/1/2022 1:02 PM
 * @Version 1.0
 */

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func main() {
	result := fibonacci(3)
	fmt.Println("result = ", result)
	fmt.Println("result = ", fibonacci(4))
	fmt.Println("result = ", fibonacci(5))
	fmt.Println("result = ", fibonacci(6))
}
