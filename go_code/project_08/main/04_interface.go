package main

import "fmt"

/**
 * @Project GoDemo
 * @File    04_interface.go
 * @Author  Augus Lee
 * @Date    2022/10/16 12:42
 * @Version 1.0
 */

type AInterface interface {
	Say()
}

type integer int

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu Say()")
}

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}

func main() {
	var stu Stu            //实现了接口的结构体
	var a AInterface = stu //接口指向一个实现了该接口的自定义类型的变量(实例)
	a.Say()
	var i integer = 10
	var b AInterface = i
	b.Say()
}
