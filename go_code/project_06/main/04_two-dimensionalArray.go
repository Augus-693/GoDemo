package main

import "fmt"

/**
 * @Project GoDemo
 * @File 04_two-dimensionalArray.go
 * @Author Augus Lee
 * @Date 2022/10/13 14:51
 * @Version 1.0
 */

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	//for 循环遍历
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr2[i][j], " ")
		}
		fmt.Println()
	}
	//for-range 遍历
	for i, v := range arr2 {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v] = %v\t", i, j, v2)
		}
		fmt.Println()
	}
}
