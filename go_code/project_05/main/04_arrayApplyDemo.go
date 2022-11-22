package main

import "fmt"

/**
 * @Project GoDemo
 * @File 04_arrayApplyDemo.go
 * @Author Augus Lee
 * @Date 2022/10/10 19:05
 * @Version 1.0
 */

func main() {
	//创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用 for 循环访问所有元素并打印出来
	var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < len(myChars); i++ {
		fmt.Printf("%c ", myChars[i])
	}
}
