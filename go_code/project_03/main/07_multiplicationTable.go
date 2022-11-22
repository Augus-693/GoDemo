package main

import "fmt"

/**
 * @Project GoDemo
 * @File 07_multiplicationTable.go
 * @Author Augus Lee
 * @Date 9/28/2022 6:22 PM
 * @Version 1.0
 */

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
