package main

import "fmt"

/**
 * @Project GoDemo
 * @File    03_inherit.go
 * @Author  Augus Lee
 * @Date    2022/10/15 16:01
 * @Version 1.0
 */

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	Name  string
	score float64
	A
}

func main() {
	var b B
	b.A.Name = "Tom" //明确指定访问A结构体的字段 Name
	b.A.age = 21
	b.A.SayOk() //明确指定访问A结构体的方法 SayOk
	b.A.hello()

	//简化写法
	b.Name = "Jack" //就近原则，会访问B结构体内的 Name 字段
	b.age = 24
	b.SayOk()
	b.hello()
}
