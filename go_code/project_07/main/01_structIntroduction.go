package main

import "fmt"

/**
 * @Project GoDemo
 * @File 01_structIntroduction.go
 * @Author Augus Lee
 * @Date 2022/10/13 21:23
 * @Version 1.0
 */

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	//创建一个 Cat 的变量
	var cat1 Cat
	cat1.Name = "白白"
	cat1.Age = 5
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"
	fmt.Println("cat1:", cat1)

	fmt.Println("猫猫的信息如下：")
	fmt.Println("name:", cat1.Name)
	fmt.Println("age:", cat1.Age)
	fmt.Println("color:", cat1.Color)
	fmt.Println("Hobby:", cat1.Hobby)
}
