package main

import "fmt"

/**
 * @Project GoDemo
 * @File 02_structDetail.go
 * @Author Augus Lee
 * @Date 2022/10/13 21:57
 * @Version 1.0
 */

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	//使用slice前，一定要make初始化并分配内存空间
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	//使用map前,一定要make初始化并分配内存空间
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"
	fmt.Println(p1)
}
