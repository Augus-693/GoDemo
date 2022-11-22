package model

import "fmt"

/**
 * @Project GoDemo
 * @File person.go
 * @Author Augus Lee
 * @Date 2022/10/14 18:18
 * @Version 1.0
 */

type person struct {
	Name string
	age  int
	sal  float64
}

// 写一个工厂模式的函数，相当于构造函数
func NewPerson(Name string) *person {
	return &person{
		Name: "name",
	}
}

// 为了访问 age 和 sal 编写一对 Set 方法和 Get 方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
