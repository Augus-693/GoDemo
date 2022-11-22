package main

import "fmt"

/**
 * @Project GoDemo
 * @File 04_methodIntroduction.go
 * @Author Augus Lee
 * @Date 2022/10/14 12:38
 * @Version 1.0
 */

type Person2 struct {
	name string
}

// 给Person2绑定一个方法输出名字
func (p Person2) test() {
	fmt.Println("test() name: ", p.name)
}

// 给 Person2 结构体添加speak方法，输出一句话
func (p Person2) speak() {
	fmt.Println(p.name, "最喜欢学Golang")
}

// 给 Person2 结构体添加jisuan方法，可以计算从1加到1000的结果
// 方法体内可以和函数一样，进行各种运算
func (p Person2) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(res)
}

// 给 Person2 结构体 jisuan2 方法,该方法可以接收一个数 n，计算从 1+..+n 的结果
func (p Person2) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(res)
}

// 给 Person2 结构体添加 getSum方法,可以计算两个数的和，并返回结果
func (p Person2) getSum(a float64, b float64) float64 {
	return a + b
}

func main() {
	var p Person2
	p.name = "John"
	p.test() //调用方法
	p.speak()
	p.jisuan()
	p.jisuan2(30)
	res := p.getSum(30, 40)
	fmt.Println(res)
}
