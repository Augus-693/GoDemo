package main

import "fmt"

/**
 * @Project GoDemo
 * @File 02_escapeChar.go
 * @Author Augus Lee
 * @Date 2022/9/20 15:34
 * @Version 1.0
 */
func main() {
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("E:\\SC\\GoDemo")
	fmt.Println("tom说\"I love you\"")
	// \r从当前行的最前面开始输出，覆盖掉以前内容（第一个Golang不打印）
	fmt.Println("GolangGolang\rGoogle")
}
