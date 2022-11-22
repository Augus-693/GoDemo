package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Project GoDemo
 * @File 05_reversePrintDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 19:38
 * @Version 1.0
 */

func main() {
	var intArr [5]int
	len := len(intArr)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("交换前：", intArr)

	//反转打印
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr[len-1-i]
		intArr[len-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后：", intArr)
}
