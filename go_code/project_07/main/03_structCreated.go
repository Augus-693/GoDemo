package main

import "fmt"

/**
 * @Project GoDemo
 * @File 03_structCreated.go
 * @Author Augus Lee
 * @Date 2022/10/14 11:46
 * @Version 1.0
 */

type Person1 struct {
	name string
	age  int
	sex  string
}

func main() {
	//方式1：直接说明
	var p1 Person1
	p1.name = "John"
	p1.age = 21
	p1.sex = "male"
	fmt.Println(p1)

	//方式2：{}
	p2 := Person1{name: "Jack", age: 19, sex: "male"}
	fmt.Println(p2)

	//方式3：&
	var p3 *Person1 = new(Person1)
	//因为p3是一个指针，因此标准的给字段赋值方式为(*p3).name，也可以写成p3.name
	(*p3).name = "Mary"
	p3.age = 20
	p3.sex = "female"
	fmt.Println(*p3)

	//方式4：{}
	var p4 *Person1 = &Person1{}
	p4.name = "Tom"
	(*p4).age = 22
	p4.sex = "female"
	fmt.Println(*p4)

	//创建结构体变量时指定字段的值
	var p5 = Person1{"小明", 22, "男"}
	p6 := Person1{"小明~~~", 22, "男"}
	var p7 = Person1{
		name: "小红",
		age:  21,
		sex:  "女",
	}
	p8 := Person1{
		name: "小红~~~",
		age:  21,
		sex:  "女",
	}
	fmt.Println(p5, p6, p7, p8)

	var p9 *Person1 = &Person1{"小明", 22, "男"}
	p10 := &Person1{"小明~~~", 22, "男"}
	var p11 = &Person1{
		name: "小红",
		age:  21,
		sex:  "女",
	}
	p12 := &Person1{
		name: "小红~~~",
		age:  21,
		sex:  "女",
	}
	fmt.Println(*p9, *p10, *p11, *p12)
}
