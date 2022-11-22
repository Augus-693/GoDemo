package main

import "fmt"

/**
 * @Project GoDemo
 * @File    cal.go
 * @Author  Augus Lee
 * @Date    2022/10/18 17:17
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int {
	return n1 - n2
}

func main() {
	res1 := addUpper(10)
	res2 := getSub(10, 5)
	fmt.Println("res1 = ", res1)
	fmt.Println("res2 = ", res2)
}
