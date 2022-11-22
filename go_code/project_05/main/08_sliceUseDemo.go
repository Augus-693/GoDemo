package main

import "fmt"

/**
 * @Project GoDemo
 * @File 08_sliceUseDemo.go
 * @Author Augus Lee
 * @Date 2022/10/12 15:33
 * @Version 1.0
 */

func main() {
	//用append内置函数，可以对切片进行动态增加
	var slice []int = []int{100, 200, 300}
	fmt.Println("slice: ", slice)
	//通过append直接给slice追加具体的元素
	slice = append(slice, 400, 500, 600)
	fmt.Println("slice: ", slice)
	//通过append将slice追加给islice
	slice = append(slice, slice...)
	fmt.Println("slice: ", slice)

	//切片使用 copy 内置函数完成拷贝
	var slice1 []int = []int{1, 2, 3, 4, 5}
	var slice2 = make([]int, 10)
	copy(slice2, slice1)
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
}
