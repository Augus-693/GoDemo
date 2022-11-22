package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_binaryFind.go
 * @Author Augus Lee
 * @Date 2022/10/13 14:19
 * @Version 1.0
 */

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		//说明要查找的数在 leftIndex 和 middle-1 中间
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明要查找的数在 middle+1 和 rightIndex 中间
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Println("找到了，下标为 %v\n", middle)
	}
}

func main() {
	arr := [6]int{12, 26, 36, 48, 68, 87}
	BinaryFind(&arr, 0, len(arr)-1, 48)
}
