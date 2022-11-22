package main

import "fmt"

/**
 * @Project GoDemo
 * @File 07_character.go
 * @Author Augus Lee
 * @Date 2022/9/20 18:23
 * @Version 1.0
 */

func main() {
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//直接输出byte值，就是输出其对应的码值
	fmt.Println("c1 = ", c1, "c2 = ", c2)
	//输出对应字符
	fmt.Println("c1=", string(c1), "c2=", string(c2))
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
}
