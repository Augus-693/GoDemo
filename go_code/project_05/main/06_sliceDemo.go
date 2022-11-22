package main

import "fmt"

/**
 * @Project GoDemo
 * @File 06_sliceDemo.go
 * @Author Augus Lee
 * @Date 2022/10/11 12:51
 * @Version 1.0
 */

func main() {
	//演示切片的基本使用
	var intArr [5]int = [...]int{0, 1, 2, 3, 4}
	//声明一个切片1
	slice := intArr[1:3] //引用intArr数组的起始下标为1，最后的下标为3（不包括3）
	fmt.Println("intArr: ", intArr)
	fmt.Println("slice: ", slice)
	fmt.Println("slice元素的个数：", len(slice)) //切片的个数是最后的下标减去起始的下标
	fmt.Println("slice的容量：", cap(slice))   //切片的容量是可以动态变化的

	var slice1 []float64 = make([]float64, 5, 10)
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1元素的个数：", len(slice1))
	fmt.Println("slice1的容量：", cap(slice1))

	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice:", strSlice)
	fmt.Println("strSlice size:", len(strSlice))
	fmt.Println("strSlice cap:", cap(strSlice))

}
