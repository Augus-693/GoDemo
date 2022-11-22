package main

import "fmt"

/**
 * @Project GoDemo
 * @File 01_bubbleSort.go
 * @Author Augus Lee
 * @Date 2022/10/12 21:58
 * @Version 1.0
 */

func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr:", *arr)
	temp := 0

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr:", *arr)
}

func main() {
	arr := [5]int{23, 65, 59, 46, 18}
	BubbleSort(&arr)
	fmt.Println("server arr:", arr)
}
