package main

import "fmt"

/**
 * @Project GoDemo
 * @File 14_pointer.go
 * @Author Augus Lee
 * @Date 2022/9/23 15:29
 * @Version 1.0
 */
func main() {
	var a int = 10
	fmt.Println("a的地址是", &a) //获取变量的地址，用&

	var ptr *int = &a
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("ptr的地址是%v\n", &ptr)
	fmt.Printf("ptr指向的值为%v\n", *ptr) //获取指针类型所指向的值，使用：*
}
