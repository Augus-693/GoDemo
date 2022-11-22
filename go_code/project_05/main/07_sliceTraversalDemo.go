package main

import "fmt"

/**
 * @Project GoDemo
 * @File 07_sliceTraversalDemo.go
 * @Author Augus Lee
 * @Date 2022/10/11 14:16
 * @Version 1.0
 */

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	//使用常规遍历
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}
	//使用for-range遍历
	for i, v := range slice {
		fmt.Printf("i = %v v = %v\n", i, v)
	}
}
