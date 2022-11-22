package main

import "fmt"

/**
 * @Project GoDemo
 * @File 01_arrayDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 15:20
 * @Version 1.0
 */

func main() {
	//定义一个数组
	var hens [7]float64
	//给数组的每个元素赋值，元素的下标是从0开始的
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 3.5
	hens[6] = 5.6

	//遍历数组求和
	total := 0.0
	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}

	//求平均值
	average := fmt.Sprintf("%.2f", total/float64(len(hens)))
	fmt.Printf("total = %v, average = %v\n", total, average)
}
