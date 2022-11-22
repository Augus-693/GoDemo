package main

import "fmt"

/**
 * @Project GoDemo
 * @File 04_moreVariate.go
 * @Author Augus Lee
 * @Date 2022/9/20 15:55
 * @Version 1.0
 */

//一次性声明多个变量
//var n1, n2, n3 int
//var n4, name, n6 = 100, "augus", 200
//n1, name1, n3 := 100, "augus", 200

// 函数外部声明，全局变量
var n1 = 100
var n2 = 200
var name1 = "tom"

// 改成一次性声明
var (
	n3    = 300
	n4    = 400
	name2 = "jerry"
)

func main() {
	fmt.Println("n1:", n1, "n2:", n2, "name1:", name1)
	fmt.Println("n3:", n3, "n4:", n4, "name2:", name2)
}
